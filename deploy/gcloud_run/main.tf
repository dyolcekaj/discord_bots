resource "google_cloud_run_service" "default" {
  name     = var.service_name
  location = "us-central1"

  template {
    spec {
      containers {
        image = var.docker_image
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}