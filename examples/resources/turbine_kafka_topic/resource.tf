resource "turbine_kafka_topic" "foo" {
  metadata {
    name = "foo"
  }
  spec {
    name = "foo"
    config = {
      "retention.ms" : "86400000"
    }
  }
}
