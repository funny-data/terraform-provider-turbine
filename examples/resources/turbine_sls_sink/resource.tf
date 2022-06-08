resource "turbine_sls_sink" "foo" {
  metadata {
    name = "foo"
  }
  spec {
    parallelism             = 1
    topic                   = "foo"
    dead_letter_queue_topic = "foo-dlq"
    project                 = "foo"
    logstore                = "foo"
    time_field              = "time"
    flush_size              = 1000
    flush_interval_seconds  = 5
  }
}
