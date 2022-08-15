#!/bin/bash

repository=$1
version=$2
username=$3
password=$4

sh fluent_bit_config_generator.sh "tail" "/host/var/lib/docker/containers/*/*.log" "*" "$repository" "$version" "log" "logsight.ai" "443" "$username" "$password" > ./fluent-bit.conf
docker-compose up -d