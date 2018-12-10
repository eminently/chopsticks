#!/bin/sh

## example of use:
## sh chopsticks-get-tx-byhash.sh <address> <api_token>

address=$(echo $1)
token=$(echo $2)
printf "running chopsticks-get-tx-byhash script...\n"
printf "querying chopsticks API : https://api.chopsticks.cash/transactions/address/${address} ...\n"

## call API
curl -v -H "Authorization: User ${token}" -H "Content-Type: application/json" -X GET http://api.chopsticks.cash/transactions/address/${address}

printf "\n ...finished!\n"