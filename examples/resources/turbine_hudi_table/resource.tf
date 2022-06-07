resource "turbine_hudi_table" "foo" {
  metadata {
    name = "foo"
  }
  spec {
    database        = "demo"
    name            = "demo"
    type            = "COPY_ON_WRITE"
    schema          = <<-EOF
      {
        "fields": [
          {
            "name": "log_id",
            "type": "string"
          },
          {
            "name": "time",
            "type": "long"
          }
        ]
      }
    EOF
    record_key      = "log_id"
    pre_combine_key = "time"
    partition_spec = [
      { "name" : "dt", "transform" : "hudi_partition_transform_date_from_unixtime(time, 'milliseconds', 'Asia/Shanghai')" }
    ]
  }
}
