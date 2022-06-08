resource "turbine_hudi_sink_job" "foo" {
  metadata {
    name = "foo"
  }
  spec {
    table {
      database = "demo"
      table    = "demo"
    }
    source {
      type             = "kafka"
      format           = "json"
      topic            = "demo"
      starting_offsets = "earliest"
    }
    delta_streamer_task {
      enabled                 = true
      operation               = "INSERT"
      interval                = 300
      batch_size_limit        = 209715200
      batch_records_limit     = 2000000
      backoff                 = 900
      max_backoff             = 3600
      default_lag_records     = 10000
      default_avg_record_size = 1000
      parallelism             = 25
      extra_properties        = {}
    }
    sync_task {
      enabled          = true
      interval         = 900
      backoff          = 900
      max_backoff      = 3600
      extra_properties = {}
    }
  }
}
