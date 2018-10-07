#!/bin/sh

## example of use:
## sh chopsticks-get-mininginfos.sh <api_token>

token=$(echo $1)
printf "running chopsticks-get-mininginfos...\n"

## call API
curl -v -H "Authorization: User ${token}" -H "Content-Type: application/json" -X GET https://api.chopsticks.cash/blockchains/miningInfo

printf "\n ...finished!\n"