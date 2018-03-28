#!/bin/sh
wget https://storage.googleapis.com/kubernetes-helm/helm-v2.8.2-linux-amd64.tar.gz
tar zxvf helm-v2.8.2-linux-amd64.tar.gz
mv linux-amd64/helm ./
rm helm-v2.8.2-linux-amd64.tar.gz
rm -rf linux-amd64

