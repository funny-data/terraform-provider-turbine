resource "turbine_xdconsole_sink" "foo" {
  metadata {
    name = "foo"
  }
  spec {
    topic = "foo"
    event = "foo"
    app   = "foo"
  }
}
