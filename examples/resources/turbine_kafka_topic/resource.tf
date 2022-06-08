resource "turbine_kafka_topic" "foo" {
  metadata {
    name = "foo"
  }
  spec {
    name               = "foo"
    partitions         = 1
    replication_factor = 2
    config = {
      "retention.ms" : "86400000"
    }
  }
}
