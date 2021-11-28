variable "gcp_region" {
    description = "GCP region"
    type        = string
    default     = "us-central1"
}

variable "gcp_project" {
    description = "GCP Project ID"
    type = string
    default = "dyolcekaj-misc"
}

variable "discord_bots" {
    description = "List of Discord bots to deploy"
    type = map(any)
    default = {
        spongbob = {
            docker_image = "gcr.io/dyolcekaj-misc/discord-spongebob"
            service_name = "discord-spongebob"
        }
    }
}