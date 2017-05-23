(ns ophion.ingest
  (:require
   [taoensso.timbre :as log]
   [cheshire.core :as json]
   [protograph.kafka :as kafka]
   [protograph.core :as protograph]
   [ophion.db :as db]
   [ophion.query :as query]
   [ophion.config :as config]
   [ophion.search :as search])
  (:import
   [protograph Protograph]))

(defn ingest-vertex
  [graph data]
  (let [vertex (query/add-vertex! graph data)
        id (.id vertex)]
    {:id id
     :data data
     :graph "vertex"}))

(defn ingest-edge
  [graph data]
  (let [edge (query/add-edge! graph data)
        id (.id edge)
        out-id (.getOutVertexId id)
        type-id (.getTypeId id)
        edge-id (.getRelationId id)
        in-id (.getInVertexId id)]
    {:id edge-id
     :out-id out-id
     :label type-id
     :in-id in-id
     :data data
     :graph "edge"}))

(defn ingest-message
  [graph producer prefix continuation message]
  (let [label (kafka/topic->label (.topic message))
        data (json/parse-string (.value message) keyword)
        _ (log/info label data)
        result (condp = label
                 "Vertex" (ingest-vertex graph data)
                 "Edge" (ingest-edge graph data))
        ;; id (.id ingested)
        ;; json (Protograph/writeJSON result)
        ;; result {:id id :data data :graph label}
        ]
    ;; (kafka/send-message producer (str prefix "." label) json)
    (continuation result)))

(defn ingest-graph
  [config graph continuation]
  (let [host (get-in config [:kafka :host])
        port (or (get-in config [:kafka :port]) "9092")
        kafka-url (str host ":" port)
        group-id (kafka/uuid)
        prefix (get-in config [:protograph :prefix])
        input-topics [(str prefix ".Vertex") (str prefix ".Edge")]
        output-prefix (get-in config [:protograph :output])
        consumer (kafka/consumer {:host kafka-url :topics input-topics})
        producer (kafka/producer kafka-url)]
    (kafka/consume
     consumer
     (partial ingest-message graph producer output-prefix continuation))))

(defn -main
  []
  (let [config (config/read-config "config/ophion.clj")
        graph (db/connect (:graph config))
        search (search/connect (merge search/default-config (:search config)))]
    (ingest-graph
     protograph/default-config
     graph
     (partial search/index-message search))))
