resource "google_compute_firewall" "pks-api-firewall" {
  name    = "pks-api-firewall"
  network = "${google_compute_network.pcf-network.name}"

  allow {
    protocol = "tcp"
    ports    = ["8443", "9021"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags = ["pks-api"]
}