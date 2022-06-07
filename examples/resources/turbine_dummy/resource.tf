resource "turbine_dummy" "foo" {
  metadata {
    name = "foo"
  }
  spec {
    arg = "bar"
  }
}
