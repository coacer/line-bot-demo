provider "google" {
  project = "associate-practice-318201"
  region  = "asia-northeast1"
  zone    = "asia-northeast1-a"
}

resource "google_container_cluster" "primary" {
  name               = "practice"
  location           = "asia-northeast1-a"
  initial_node_count = 2
  node_config {
    machine_type = "n1-standard-1"
  }
}

# TODO Create Spanner Cluster
# TODO Create GCR
