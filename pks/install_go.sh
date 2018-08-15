#!/bin/sh
wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz
echo "export GOPATH=$HOME/go" >> .profile
echo "export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin" >> .profile

