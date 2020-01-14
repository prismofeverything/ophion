(defproject ophion "0.0.14"
  :description "graph queries as data"
  :url "http://github.com/bmeg/ophion"
  :license {:name "Eclipse Public License"
            :url "http://www.eclipse.org/legal/epl-v10.html"}
  :dependencies [[org.clojure/clojure "1.10.0"]
                 [aleph "0.4.3"]
                 [polaris "0.0.19"]
                 [cheshire "5.7.1"]
                 [com.taoensso/timbre "4.8.0"]
                 [protograph "0.0.19"]
                 [clojurewerkz/elastisch "2.2.1" :exclusions [commons-codec]]
                 [org.apache.tinkerpop/gremlin-core "3.4.1" :extension "pom"]
                 [com.novemberain/monger "3.1.0"]]
  :repositories [["sonatype snapshots"
                  "https://oss.sonatype.org/content/repositories/snapshots"]
                 ["sonatype releases"
                  "https://oss.sonatype.org/content/repositories/releases"]]
  :jvm-opts ["-Xmx4g" "-Xms4g"]
  :main ophion.server)
