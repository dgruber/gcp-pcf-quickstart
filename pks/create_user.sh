#!/bin/sh

# Requires Credentials from:
# OpsMan -> PKS -> Credentials -> Uaa Admin Secret

# REPLACE ADDRESS
uaac target api.pks.YOUR.DOMAIN:8443 --skip-ssl-validation
uaac token client get admin

# REPLACE USERNAME and PASSWORT
uaac user add USERNAME --emails your@email.io -p PASSWORD 
uaac member add pks.clusters.admin daniel
