terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
    }
  }
}


provider "google" {
  project     = var.project
  credentials = "./admin-key.json"
}

data "google_project" "project" {}

resource "google_project_service" "service" {
  for_each = toset([
    "bigquery.googleapis.com",
    "iam.googleapis.com",
    "cloudkms.googleapis.com",
    "compute.googleapis.com",
    "storage.googleapis.com",
    "sqladmin.googleapis.com",
    "dataproc.googleapis.com",
    "container.googleapis.com",
    "logging.googleapis.com",
    "pubsub.googleapis.com",
    "cloudresourcemanager.googleapis.com"
  ])

  service            = each.key
  disable_on_destroy = false
}

resource "google_project_iam_custom_role" "confromity_custom_role" {
  role_id     = "ConfromityCustomRole"
  title       = "Custom Role for Confromity"
  description = "Custom Role for Confromity Terraform Provider"
  permissions = [
    "bigquery.datasets.get",
    "resourcemanager.projects.get",
    "resourcemanager.projects.getIamPolicy",
    "cloudkms.keyRings.list",
    "cloudkms.cryptoKeys.list",
    "cloudkms.cryptoKeys.getIamPolicy",
    "cloudkms.locations.list",
    "compute.firewalls.list",
    "compute.networks.list",
    "compute.subnetworks.list",
    "storage.buckets.list",
    "storage.buckets.getIamPolicy",
    "compute.instances.list",
    "compute.images.list",
    "compute.images.getIamPolicy",
    "cloudsql.instances.list",
    "compute.backendServices.list",
    "compute.globalForwardingRules.list",
    "dataproc.clusters.list",
    "container.clusters.list",
    "logging.sinks.list",
    "pubsub.topics.list"
  ]
}

resource "google_service_account" "conformity_service_account" {
  account_id   = "conformity-integration-sa"
  display_name = "Conformity Integration SA"
}

resource "google_project_iam_binding" "conformity_custom_role_binding" {
  project = data.google_project.project.id
  role = "projects/${data.google_project.project.project_id}/roles/${google_project_iam_custom_role.confromity_custom_role.role_id}"

  members = [
    "serviceAccount:${google_service_account.conformity_service_account.email}",
  ]
}

resource "google_service_account_key" "conformity_sa_key" {
  service_account_id = google_service_account.conformity_service_account.name
  public_key_type    = "TYPE_X509_PEM_FILE"
}

resource "local_file" "sa_key_json" {
  content  = base64decode(google_service_account_key.conformity_sa_key.private_key)
  filename = "${path.module}/conformity-sa-key.json"

}