#!/bin/bash

repository=$1
version=$2
username=$3
password=$4
url=$5
port=$6

sh fluent_bit_config_generator.sh "tail" "/host/var/lib/docker/containers/*/*.log" "*" "$repository" "$version" "log" "$url" "$port" "$username" "$password" > ./fluent-bit.conf
docker-compose up -d
sleep 150s
docker-compose down
