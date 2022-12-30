## 1. Choosing the partitions count
- Each partition can handle a thorughput of a few MB/s (measure it for your setup!)
- More partitions implies:
  + Better parallelism, better throughput
  + Ability to run more consumers in a group to scale (max as many consumers per group as partitions)
  + Ability to leverage more brokers if you have a large cluster
  + BUT more elections to perform for Zookeeper (if using Zookeeper)
  + BUT more files opened on Kafka
- Guidelines:
  + **Partitions per topic = MILLION DOLLAR QUESTION**
    + (Intuition) Small cluster (< 6 brokers): 3 x # brokers
    + (Intuition) Big cluster (> 12 brokers): 2 x # brokers
    + Adjust for number of consumers you need to run in parallel at peak throughput
    + Adjust for producer throughput (increase if super high throughput or projected increase in the next 2 years)
  + **TEST** Every Kafka cluster will have different performance
  + Don't systematically create topics with 1000 partitions
## 2. Choosing the partitions count
- Should be at least 2, usually 3, maximum 4
- The higher the replication factor (N):
  + Better durability of your system (N-1 brokers can fail)
  + Better availability of your system (N-min.insync.replicas if producer acks=all)
  + BUT more replication (higher latency if acks=all)
  + BUT more disk space on your system (50% more if RF is 3 instead of 2)
- **Guidelines**:
  + Set it to 3 to get started (you must have at least 3 brokers for that)
  + If replication performance is an issue, get a better broker instead of less RF
  + Never set it to 1 in production

## 3. Cluster guidelines
- Total number of partitions in the cluster:
  + Kafka with Zookeeper: max 200,000 partitions (Nov 2018) - Zookeeper Scaling limit
    + Still recommend a maximum of 4000 partitions per broker (soft limit)
  + Kafka with KRaft: potenially millions of partitions
- If you need more partitions in your cluster, add brokers instead
- If you need more 200,000 partitions in your cluster (it will take time to get there), follow the Netflix model and create more Kafka clusters
- Overall, you don't need a topic with 1000 partitions to achieve high throughput. Start at a reasonable number and test the performance

## 4. Topic Naming conventions
- FROM: https://cnr.sh/essays/how-paint-bike-shed-kafka-topic-naming-conventions
- < mesage type >.< dataset name >.< data name >.< data format >
- Message Type:
  + **logging**: For logging data (slf4j, syslog, etc)
  + **queuing**: For classical queuing use cases
  + **tracking**: For tracking events such as user clicks, page views, ad views, etc
  + **etl/db**: For ETL and CDC use cases such as database feeds
  + **streaming**: For intermediate topics created by stream processing pipelines.
  + **push**: For data that's being pushed from offline (batch computation) enviroments into online enviroments
  + **user**: For user-specific data such as scratch and test topics
- The dataset name is analogous to a database name in traditional RDBMS systems. It's used as a category to group topics together
- The data name field is analogous to a table name in traditional RDBMS systems, though it's fine to include further dotted notation if developers wish to impose their own hierarchy within the dataset namespace
- The data format for example .avro, .json, .text, .protobuf, .csv, .log
- Use snake_case