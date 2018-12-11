#!/bin/sh

## example of use:
## sh chopsticks-get-tx-byhash-fromchain.sh <hash> <api_token> <chain>

hash=$(echo $1)
token=$(echo $2)
chain=$(echo $3)
printf "running chopsticks-get-tx-byhash script...\n"
printf "querying chopsticks API : https://api.chopsticks.cash/transactions/${hash}?chainType=${chain} ...\n"

## call API
curl -v -H "Authorization: User ${token}" -H "Content-Type: application/json" -X GET https://api.chopsticks.cash/transactions/${hash}?chainType=${chain}

printf "\n ...finished!\n"