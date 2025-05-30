#!/bin/bash

cd "${0%/*}" # ensure cwd is script dir

echo -e "\n starting mock infrastructure...\n"

docker compose -f mock-infrastructure.yml up -d --build 

sleep 10 # wait for localstack

region=us-east-1
table_name=kvp-table
endpoint_url=http://localhost:4566

aws dynamodb create-table \
--endpoint-url $endpoint_url \
--table-name $table_name \
--region $region \
--key-schema AttributeName=key,KeyType=HASH \
--attribute-definitions AttributeName=key,AttributeType=S \
--provisioned-throughput ReadCapacityUnits=10,WriteCapacityUnits=5 \
--no-cli-pager

