#!/bin/bash

inputName=$1
fileLocation=$2
matchPattern=$3
repository=$4
version=$5
message=$6
host=$7
port=$8
logsightUsername=$9
logsightPassword=${10}

echo "[INPUT]
    Name $inputName
    Path $fileLocation
    multiline.parser  docker, cri
    DB /tail_docker.db
    Refresh_Interval 1
[SERVICE]
    Flush 1
    Daemon Off
[FILTER]
    Name modify
    Match $matchPattern
    Rename $message message
    Add tags.version $version
    Add tags.repository $repository
[FILTER]
    Name nest
    Match $matchPattern
    Operation nest
    Wildcard tags.*
    Nested_under tags
    Remove_prefix tags.
[OUTPUT]
    Name http
    Match $matchPattern
    Host $host
    Port $port
    http_User $logsightUsername
    http_Passwd $logsightPassword
    tls On
    uri /api/v1/logs/singles
    Format json
    json_date_format iso8601
    json_date_key timestamp
    Retry_Limit False"