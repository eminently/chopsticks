#!/bin/sh

## example of use:
## sh chopsticks-get-tx-byhash.sh <hash> <api_token>

hash=$(echo $1)
token=$(echo $2)
printf "running chopsticks-get-tx-byhash script...\n"
printf "querying chopsticks API : https://api.chopsticks.cash/transactions/${hash} ...\n"

## call API
curl -v -H "Authorization: User ${token}" -H "Content-Type: application/json" -X GET https://api.chopsticks.cash/transactions/${hash}

printf "\n ...finished!\n"