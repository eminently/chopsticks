#!/bin/sh

## example of use:
## sh chopsticks-send-signedrawtx.sh <JSON> <api_token>
## JSON ex: '{"tx_hex":"0100000001d6a2d0329f340a3ab5d5aa8863e5e021104b169d5b1f68d5d4c5ff4169011cdd010000006a47304402204bea3cc9cd21ee474ae75c06f5dea5b7407c8dba6cbe38b721dc79c7e423888002200f749b0013b840100d355fe5e83ee97dbfaafc3e94bed72abe96eca4969513ad412102e7d2f77f45b171ba9e38f3acac9bb8fdbfd20f548a8fb2706035271c3dab5871ffffffff0258020000000000001976a914b8e7f03b158f1ff0203ebf8d7f4a14c20846da8388acda740000000000001976a9148fe1ea91d75774681790ad08f3ee9760f518d4f588ac00000000","blockchains":["XBC","XBS","XBN"],"voting":true}'


json=$(echo $1)
token=$(echo $2)
printf "running chopsticks-send-signedrawtx script...\n"
printf "\n posting json data : $json ...\n"
printf " to chopsticks API : https://api.chopsticks.cash/transactions ...\n"

## call API
curl -v -d "${json}" -H "Content-Type: application/json" -H "Authorization: User ${token}" -X POST https://api.chopsticks.cash/transactions

printf "\n ...finished!\n"