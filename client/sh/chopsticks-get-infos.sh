#!/bin/sh

## example of use:
## sh chopsticks-get-infos.sh <api_token>

token=$(echo $1)
printf "running chopsticks-get-infos...\n"

## call API
curl -v -H "Authorization: User ${token}" -H "Content-Type: application/json" -X GET https://api.chopsticks.cash/blockchains/info

printf "\n ...finished!\n"