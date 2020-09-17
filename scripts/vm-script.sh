#!/bin/bash

sudo apt install unzip

unzip food_web_app.zip

sudo apt update && install
sudo apt-get update

sudo apt-get remove docker docker-engine docker.io containerd runc
sudo apt-get install docker-ce docker-ce-cli containerd.io

sudo docker run hello-world

docker load -i food_image.docker

sudo docker-compose up -d

exit