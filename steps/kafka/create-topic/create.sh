/usr/bin/kafka-topics --zookeeper ${ZOOKEPER_HOST} --topic ${TOPIC_NAME} --create --partitions ${PARTITIONS:-1} --replication-factor ${REPLICATION_FACTOR:-1}