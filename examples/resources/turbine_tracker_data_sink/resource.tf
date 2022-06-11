resource "turbine_tracker_data_sink" "foo" {
  metadata {
    name = "foo"
  }
  spec {
    topic       = "foo"
    environment = "foo"
  }
}
