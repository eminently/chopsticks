/**
 *  @copyright defined in chopsticks/LICENSE.txt
 *  @author Romain Pellerin - romain@eminent.ly
 *
 *  Donation appreciated :)
 *
 *  Bitcoin Cash $BCH wallet: 1HrhBfFRFovHv8EMxsuB9EcZgamtuH3fMc
 */
package client

import (
	"encoding/json"
	"fmt"
	"github.com/chopsticks/model"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// PLEASE REVIEW THESE PARAMETERS BEFORE EXECUTING YOUR FIRST TESTS
var (
	testWalletPrivKey = os.Getenv("TEST_WALLET_PRIV_KEY") // SET IT! never put private keys in your code or it could finish on a repo...

	amount int64 = 600  // satoshis
	fee    int64 = 1000 // satoshis

	recipientAddress = "qzuw0upmzk83lupq86lc6l62znpqs3k6svtf292dql" // CHANGE IT !!!! or keep it if you want to DONATE amount to eminent.ly

	utxoSource        = "d9f01d93b215ac701c7f9a56c8d5d1bdf2e0498047da81d3bf6fc23a2ee3bbb5" // CHANGE IT !! this is an example of source UTXO
	utxoIndex  uint32 = 1                                                                 // CHANGE IT !!  this is an example of source UTXO
	utxoValue  int64  = 32114

	addressSource = "1E7nGNVkhCfHjWaPS2Q9YPppsPDifqXr2d" // CHANGE IT !!  this is an example of wallet address

	signedHex = "" // CHANGE IT !!  if you want to run TestBchSendRawTransaction without TestClientCreateBitcoinCashTransaction

	apiToken = os.Getenv("API_TOKEN") // SET IT !! get a token on https://chopsticks.cash

	transaction *model.Transaction
)

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

func TestGetTransaction(t *testing.T) {

	response, err := GetTransaction(utxoSource, apiToken)

	assert.Nil(t, err, "shoult not report any")

	data, _ := json.Marshal(response)

	fmt.Println("transactions retrieved: ", string(data))

}

func TestGetTransactionsByAddress(t *testing.T) {

	response, err := GetTransactionByAddress(addressSource, apiToken)

	assert.Nil(t, err, "shoult not report any")

	data, _ := json.Marshal(response)

	fmt.Println("transactions retrieved: ", string(data))

}

func TestGetRawMempool(t *testing.T) {

	response, err := GetRawMempools(apiToken)

	data, _ := json.Marshal(response)

	fmt.Println("response data: ", string(data))

	assert.Nil(t, err, "shoult not report any")

}

func TestGetMiningInfos(t *testing.T) {

	response, err := GetMiningInfos(apiToken)

	data, _ := json.Marshal(response)

	fmt.Println("response data: ", string(data))

	assert.Nil(t, err, "shoult not report any")
}

func TestGetInfos(t *testing.T) {

	response, err := GetInfos(apiToken)

	data, _ := json.Marshal(response)

	fmt.Println("response data: ", string(data))

	assert.Nil(t, err, "shoult not report any")
}

func TestClientCreateBitcoinCashTransaction(t *testing.T) {

	var err error

	transaction, err = CreateBitcoinCashTransaction(testWalletPrivKey, recipientAddress, amount, fee, utxoSource, utxoIndex, utxoValue)

	assert.Nil(t, err, "shoult not report any")

	data, _ := json.Marshal(transaction)

	fmt.Println("transaction created: ", string(data))
	fmt.Println("transaction signed hex: ", transaction.SignedTx)

}

func TestBchSendRawTransaction(t *testing.T) {

	if transaction == nil {
		transaction = &model.Transaction{}
		transaction.SignedTx = signedHex
	}

	response, err := SendRawTransactionToChopsticks(transaction.SignedTx,
		[]string{
			model.BLOCKCHAIN_TYPE_XBC_MAINNET, // your first choice
			model.BLOCKCHAIN_TYPE_XBU_MAINNET, // your second choice
			model.BLOCKCHAIN_TYPE_XBS_MAINNET, // your third choice
			model.BLOCKCHAIN_TYPE_XBN_MAINNET,
		},
		true,
		apiToken)

	data, _ := json.Marshal(response)

	fmt.Println("response data: ", string(data))

	assert.Nil(t, err, "shoult not report any")

}