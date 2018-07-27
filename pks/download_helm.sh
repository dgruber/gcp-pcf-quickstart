#!/bin/sh
HELM="helm-v2.10.0-rc.1-linux-amd64.tar.gz"
if [ ! -f helm ]; then
   wget https://storage.googleapis.com/kubernetes-helm/$HELM
   tar zxvf $HELM 
   mv linux-amd64/helm ./
   rm $HELM
   rm -rf linux-amd64
fi

