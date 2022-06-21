resource "turbine_template_table" "foo" {
  metadata {
    name = "foo"
  }
  spec {
    database = "default"
    name     = "foo"
    schema   = <<-EOF
    {
      "fields": [
        {
          "name": "seq",
          "type": "int"
        },
        {
          "name": "sign",
          "type": "string"
        },
        {
          "name": "name",
          "type": "string"
        }
      ]
    }
    EOF
    data     = <<-EOF
    [
      [1, "feature-3", "feature-3"],
      [1, "feature-5", "feature-5"],
      [1, "feature-21", "feature-21"]
    ]
    EOF
  }
}
