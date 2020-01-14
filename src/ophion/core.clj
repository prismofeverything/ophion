(ns ophion.core
  (:require
   [clojure.string :as string]
   [clojure.java.io :as io]
   [manifold.stream :as stream]
   [aleph.http :as http]
   [cheshire.core :as json]
   [taoensso.timbre :as log]
   [ring.middleware.resource :as resource]
   [ring.middleware.params :as params]
   [ring.middleware.keyword-params :as keyword]
   [polaris.core :as polaris]
   [protograph.template :as protograph]
   [ophion.config :as config]
   [ophion.query :as query]
   [ophion.search :as search]
   [ophion.mongo :as mongo]
   [ophion.store :as store]
   [ophion.aggregate :as aggregate])
  (:import
   [java.io InputStreamReader]
   [ch.qos.logback.classic Logger Level]))

(defn read-json
  [body]
  (json/parse-stream (InputStreamReader. body) keyword))

(defn append-newline
  [s]
  (str s "\n"))

(def output
  (comp
   append-newline
   json/generate-string))

(defn default-graph
  []
  (let [config (config/read-config "config/ophion.clj")
        graph (db/connect (:graph config))
        search (search/connect (:search config))]
    {:graph graph
     :search search}))

(defn fetch-schema-handler
  [schema request]
  {:status 200
   :headers {"content-type" "application/json"}
   :body (json/generate-string schema)})

(defn find-vertex-handler
  [graph request]
  (let [gid (-> request :params :gid)
        vertex (query/find-vertex graph gid)
        out (query/vertex-connections vertex 100)]
    {:status 200
     :headers {"content-type" "application/json"}
     :body (json/generate-string out)}))

(defn check-query-cache
  [graph search cache query]
  (if-let [cached (get @cache query)]
    cached
    (let [result (query/evaluate {:graph graph :search search} query)
          out (mapv output result)]
      (swap! cache assoc query out)
      out)))

(defn vertex-query-handler
  [graph search cache request]
  (let [raw-query (read-json (:body request))
        ;; query (query/delabelize raw-query)
        _ (log/info (mapv identity (query/delabelize raw-query)))
        ;; out (check-query-cache graph search cache query)
        ;; transaction (.newTransaction graph)
        ;; db {:graph graph :search search :transaction transaction}
        ;; result (query/evaluate db query)
        {:keys [results commit]} (query/perform {:graph graph :search search} raw-query)
        out (map output results)
        source (stream/->source out)]
    (stream/on-drained
     source
     (fn []
       (log/debug "query complete")
       ;; (.rollback transaction)
       (commit)))
    {:status 200
     :headers {"content-type" "application/json"}
     :body source}))

(defn find-edge-handler
  [graph request]
  (let [gid (query/build-edge-gid (:params request))
        _ (log/info gid)
        edge (query/find-edge graph gid)
        out (query/edge-connections edge)]
    {:status 200
     :headers {"content-type" "application/json"}
     :body (json/generate-string out)}))

(defn edge-query-handler
  [graph request]
  {:status 200
   :headers {"content-type" "application/json"}
   :body "found"})

(defn save-query-handler
  [graph search mongo request]
  (let [query (read-json (:body request))
        relevant (select-keys query [:user :key :focus :path :query :current])]
    (log/info query)
    (store/store-query {:graph graph :search search} mongo relevant)
    {:status 200
     :headers {"content-type" "application/json"}
     :body (json/generate-string (store/all-queries mongo))}))

(defn all-queries-handler
  [mongo request]
  (let [queries (store/all-queries mongo)]
    {:status 200
     :headers {"content-type" "application/json"}
     :body (json/generate-string queries)}))

(defn aggregate-query-handler
  [mongo protograph request]
  (let [query (read-json (:body request))
        label (get-in request [:params :label])
        result (aggregate/evaluate mongo label query)]
    {:status 200
     :headers {"content-type" "application/json"}
     :body (stream/->source result)}))

(defn query-comparison-handler
  [mongo request]
  (let [queries (read-json (:body request))
        comparison (store/query-comparison mongo (:queries queries))]
    (log/info "queries" queries)
    {:status 200
     :headers {"content-type" "application/json"}
     :body (json/generate-string comparison)}))

(defn parse-int
  ([n] (parse-int n 0))
  ([n default]
   (try
     (Integer/parseInt n)
     (catch Exception e default))))

(defn search-counts-handler
  [search request]
  (let [{:keys [from size query terms]} (:params request)
        out (search/aggregate search (string/split terms #",") query (parse-int size 100) (parse-int from))]
    (log/info out)
    (log/info from size query terms)
    {:status 200
     :headers {"content-type" "application/json"}
     :body (json/generate-string out)}))

(defn fetch-schema
  [schema]
  (fn [request]
    (log/info "--> fetch schema")
    (#'fetch-schema-handler schema request)))

(defn find-vertex
  [graph]
  (fn [request]
    (log/info "--> find vertex")
    (#'find-vertex-handler graph request)))

(defn vertex-query
  [graph search]
  (let [cache (atom {})]
    (fn [request]
      (log/info "--> vertex query")
      (#'vertex-query-handler graph search cache request))))

(defn find-edge
  [graph]
  (fn [request]
    (log/info "--> find edge")
    (#'find-edge-handler graph request)))

(defn edge-query
  [graph]
  (fn [request]
    (log/info "--> edge query")
    (#'edge-query-handler graph request)))

(defn search-counts
  [search]
  (fn [request]
    (log/info "--> search counts")
    (#'search-counts-handler search request)))

(defn save-query
  [graph search mongo]
  (fn [request]
    (log/info "--> save query")
    (#'save-query-handler graph search mongo request)))

(defn all-queries
  [mongo]
  (fn [request]
    (log/info "--> all queries")
    (#'all-queries-handler mongo request)))

(defn aggregate-query
  [mongo protograph]
  (fn [request]
    (log/info "--> aggregate query")
    (#'aggregate-query-handler mongo protograph request)))

(defn query-comparison
  [mongo]
  (fn [request]
    (log/info "--> query comparison")
    (#'query-comparison-handler mongo request)))

(defn home
  [request]
  {:status 200
   :headers {"content-type" "text/html"}
   :body (config/read-resource "public/index.html")})

(defn ophion-routes
  [graph search protograph mongo]
  [["/" :home #'home]
   ["/schema/protograph" :schema (fetch-schema protograph)]
   ["/vertex/find/:gid" :vertex-find (find-vertex graph)]
   ["/vertex/query" :vertex-query (vertex-query graph search)]
   ["/edge/find/:from/:label/:to" :edge-find (find-edge graph)]
   ["/edge/query" :edge-query (edge-query graph)]
   ["/search/counts" :search-counts (search-counts search)]
   ["/ophion/query/:label" :aggregate-query (aggregate-query mongo protograph)]
   ["/query/save" :save-query (save-query graph search mongo)]
   ["/query/all" :all-queries (all-queries mongo)]
   ["/query/compare" :query-comparison (query-comparison mongo)]])

(defn start
  []
  (log/info "starting server")
  (let [config (config/read-config "config/ophion.clj")
        graph (db/connect (:graph config))
        search (search/connect (:search config))
        protograph (protograph/load-protograph
                    (or
                     (get-in config [:protograph :path])
                     "resources/config/protograph.yml"))
        schema (protograph/graph-structure protograph)
        mongo (mongo/connect! (get config :mongo))
        routes (polaris/build-routes (ophion-routes graph search schema mongo))
        router (polaris/router routes)
        app (-> router
                (resource/wrap-resource "public")
                (keyword/wrap-keyword-params)
                (params/wrap-params))]
    (log/info schema)
    (http/start-server app {:port (or (:port config) 4443)})))

(defn -main
  []
  (start)
  (while true
    (Thread/sleep 1111)))
