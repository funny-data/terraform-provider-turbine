resource "turbine_grafana_datasource" "foo" {
  metadata {
    name = "foo"
  }
  spec {
    name          = "foo"
    org_name      = "foo"
    hudi_database = "foo"
  }
}
