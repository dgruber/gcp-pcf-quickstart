#!/bin/sh
sudo apt-add-repository ppa:brightbox/ruby-ng -y
sudo apt-get update -y
sudo apt-get install ruby2.4 ruby2.4-dev g++ -y
sudo gem install cf-uaac -y
