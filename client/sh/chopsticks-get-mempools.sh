#!/bin/sh

## example of use:
## sh chopsticks-get-mempoolraw.sh <api_token>

token=$(echo $1)
printf "running chopsticks-get-mempoolraw...\n"

## call API
curl -v -H "Authorization: User ${token}" -H "Content-Type: application/json" -X GET https://api.chopsticks.cash/blockchains/mempool

printf "\n ...finished!\n"