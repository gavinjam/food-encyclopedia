#!/bin/bash

cd /Users/gavinjampani/go/src

# verify that zip and files do not exist in VM; if they do, then rm them

# docker save
docker save -o food_image.docker food_image

# zip docker image, docker-compose, initdb.sql
zip -r food_web_app.zip food_image.docker docker-compose.yml initdb.sql

# scp zip to vm
scp food_web_app.zip gavin@104.42.62.212:/home/gavin/

# ssh vm
#ssh admin@192.168.0.215

# unzip
#unzip food_web_app.zip

# docker load image
#docker load -i food_image.docker

# docker-compose up
#docker-compose up

# exit VM
#exit

rm food_web_app.zip
rm food_image.docker