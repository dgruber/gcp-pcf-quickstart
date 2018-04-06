resource "google_compute_address" "pks-api" {
  name = "pks-api"
}

resource "google_compute_target_pool" "pks-api" {
  name = "pks-api"
  session_affinity = "NONE"
}

resource "google_compute_forwarding_rule" "pks-api-uaa" {
  name                  = "pks-api-uaa"
  target                = "${google_compute_target_pool.pks-api.self_link}"
  ip_address            = "${google_compute_address.pks-api.address}"
  ip_protocol           = "TCP"
  port_range            = "8443"
}

resource "google_compute_forwarding_rule" "pks-api-pks" {
  name                  = "pks-api-pks"
  target                = "${google_compute_target_pool.pks-api.self_link}"
  ip_address            = "${google_compute_address.pks-api.address}"
  ip_protocol           = "TCP"
  port_range            = "9021"
}
