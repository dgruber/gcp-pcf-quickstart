resource "random_id" "pks_account" {
  byte_length = 4
}

resource "google_service_account" "pks_account" {
  display_name = "PKS Account"
  account_id   = "pks-${random_id.pks_account.hex}"
}

resource "google_service_account_key" "pks_account" {
  service_account_id = "${google_service_account.pks_account.id}"
}

resource "google_project_iam_member" "pks_account" {
  project = "${var.project}"
  role    = "roles/owner"
  member  = "serviceAccount:${google_service_account.pks_account.email}"
}
