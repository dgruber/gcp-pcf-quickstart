#!/bin/sh
if [ ! -f kubectl ]; then
   curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.10.3/bin/linux/amd64/kubectl
   chmod +x kubectl
fi

