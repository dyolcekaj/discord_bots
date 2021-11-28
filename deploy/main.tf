provider "google" {
 project     = "dyolcekaj-misc"
 region      = var.gcp_region
}

module "gcloud_run" {
    source = "./gcloud_run"

    for_each = var.discord_bots

    docker_image = each.value.docker_image
    service_name = each.value.service_name
}