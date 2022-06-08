resource "turbine_ingest_metadata" "foo" {
  metadata {
    name = "foo"
  }
  spec = <<-EOF
    {
      "name": "sausage_debug",
      "access_keys": [
        {
          "id": "rQJEk4mz66",
          "secret": "CcYyCc=="
        }
      ],
      "messages": [
        {
          "type": "Event",
          "processors": [
            {
              "type": "DebugProcessor",
              "config": {
                "id": "99"
              }
            }
          ]
        }
      ]
    }
  EOF
}
