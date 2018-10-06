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
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/chopsticks/common"
	"github.com/chopsticks/errors"
	"github.com/chopsticks/model"
	"github.com/chopsticks/network"
	"github.com/cpacia/bchutil"
	"strings"
)

const CHOPSTICKS_API_URL = "https://api.chopsticks.cash"

/**
 * This method creates a Bitcoin Cash transaction and sign it
 * 	- privateKey: private key of transaction origin (sender)
 * 	- recipientAddress: the recipient address in legacy format (not in cashaddr)
 * 	- amount: transaction amount in satoshis
 * 	- fee: transaction fee
 * 	- utxoSourceHash: hash of the transaction you want to use as an input (source UTXO) and will take the amount and fee from
 * 	- utxoSourceIndex: index of the output of the source UTXO you want to use
 * 	- utxoSourceAmount: total unspent amount of satoshis (balance) available in the source UTXO
 */
func CreateBitcoinCashTransaction(privateKey string, recipientAddress string, amount int64, fee int64,
	utxoSourceHash string, utxoSourceIndex uint32, utxoSourceAmount int64) (*model.Transaction, error) {

	transaction := model.Transaction{}

	secretHex, err := hex.DecodeString(privateKey)

	// extract PublicKey and PrivateKey
	var priv *btcec.PrivateKey
	var pub *btcec.PublicKey
	var err2 *errors.AppError

	priv, pub = common.PrivKeyFromBytes(btcec.S256(), secretHex)

	if err2 != nil {
		return &transaction, err
	}

	// convert pub key to cash addr
	serial := pub.SerializeCompressed()
	serialBtc, err := btcutil.NewAddressPubKey(serial, &chaincfg.MainNetParams)

	if err != nil && serialBtc == nil {
		return &transaction, err
	}

	addresspubkey, err := bchutil.NewCashAddressPubKeyHash(serialBtc.AddressPubKeyHash().Hash160()[:], &chaincfg.MainNetParams)

	if err != nil && serialBtc == nil {
		return &transaction, err
	}

	// initialize source tx from utxoSourceHash (unspent UTXO)
	sourceTx := wire.NewMsgTx(wire.TxVersion)

	sourceUtxoHash, err := chainhash.NewHashFromStr(utxoSourceHash)

	if err != nil {
		return &transaction, err
	}

	// For this example, create a fake transaction that represents what
	// would ordinarily be the real transaction that is being spent.  It
	// contains a single output that pays to address in the amount of 1 BTC.
	originTx := wire.NewMsgTx(wire.TxVersion)
	prevOut := wire.NewOutPoint(sourceUtxoHash, utxoSourceIndex)
	txIn := wire.NewTxIn(prevOut, nil, [][]byte{})
	originTx.AddTxIn(txIn)

	sourceAddress, err := bchutil.DecodeAddress(addresspubkey.EncodeAddress(), &chaincfg.MainNetParams)

	if err != nil {
		return &transaction, err
	}

	sourcePkScript, err := bchutil.PayToAddrScript(sourceAddress)

	if err != nil {
		fmt.Println(err)
		return &transaction, err
	}

	txOut := wire.NewTxOut(utxoSourceAmount, sourcePkScript)
	originTx.AddTxOut(txOut)
	//originTxHash := originTx.TxHash()

	// Create the transaction to redeem the input transaction.
	redeemTx := wire.NewMsgTx(wire.TxVersion)

	// Add the input(s) the redeeming transaction will spend.  There is no
	// signature script at this point since it hasn't been created or signed
	// yet, hence nil is provided for it.
	prevOut = wire.NewOutPoint(sourceUtxoHash, utxoSourceIndex)
	txIn = wire.NewTxIn(prevOut, nil, nil)
	redeemTx.AddTxIn(txIn)

	// load destination address
	destinationAddress, err := bchutil.DecodeAddress(recipientAddress, &chaincfg.MainNetParams)

	if err != nil {
		fmt.Println(err)
		return &transaction, err
	}

	// build pay to destination script
	destinationPkScript, err := bchutil.PayToAddrScript(destinationAddress)

	if err != nil {
		fmt.Println(err)
		return &transaction, err
	}

	// destination of the funds
	txOut = wire.NewTxOut(amount, destinationPkScript)
	redeemTx.AddTxOut(txOut)

	txOut2 := wire.NewTxOut(utxoSourceAmount-amount-fee, sourcePkScript)
	redeemTx.AddTxOut(txOut2)

	// Sign the redeeming transaction.
	lookupKey := func(a btcutil.Address) (*btcec.PrivateKey, bool, error) {
		// Ordinarily this function would involve looking up the private
		// key for the provided address, but since we passed it in param
		// this is just returning the priv key we extracted above
		return priv, true, nil
	}

	// sign the source UTXO to prove ownership of funds
	sigScript, err := bchutil.SignTxOutput(&chaincfg.MainNetParams,
		redeemTx, 0, originTx.TxOut[0].PkScript, txscript.SigHashAll,
		txscript.KeyClosure(lookupKey), nil, nil, utxoSourceAmount)

	if err != nil {
		fmt.Println(err)
		return &transaction, err
	}

	redeemTx.TxIn[0].SignatureScript = sigScript

	// serialize the full transaction in hexadecimal form
	var unsignedTx bytes.Buffer
	var signedTx bytes.Buffer

	sourceTx.Serialize(&unsignedTx)
	redeemTx.Serialize(&signedTx)

	fmt.Println("signed hexadecimal representation of transaction: ", signedTx)

	// build and return the transaction object
	transaction.Hash = redeemTx.TxHash().String()
	transaction.Amount = amount
	transaction.SignedTx = hex.EncodeToString(signedTx.Bytes())

	return &transaction, nil
}

func SendRawTransactionToChopsticks(signedTxHex string, chains []string, voting bool, apiToken string) (*model.TransactionResponse, *errors.AppError) {

	request := model.TransactionRequest{}
	request.TxHex = signedTxHex
	request.Blockchains = chains
	request.Voting = voting

	bytes, err := json.Marshal(&request)

	if err != nil {
		return nil, errors.NewAppError(err, "error trying to marshall request", -1, nil)
	}

	fmt.Println("request: ", string(bytes))

	trxData, appErr := network.PostRawData(CHOPSTICKS_API_URL+"/transactions", string(bytes), apiToken)

	if appErr != nil {
		return nil, appErr
	}

	response := model.TransactionResponse{}

	dec := json.NewDecoder(strings.NewReader(string(trxData)))
	errD := dec.Decode(&response)

	if errD != nil {
		return nil, errors.NewAppError(nil, "cannot parse transaction response: "+string(trxData), -1, nil)
	}

	return &response, nil
}

func GetTransaction(hash string, apiToken string) (*model.TransactionResponse, *errors.AppError) {
	trxData, appErr := network.Get(CHOPSTICKS_API_URL+"/transactions/"+hash, nil, apiToken )

	if appErr != nil {
		return nil, appErr
	}

	response := model.TransactionResponse{}

	dec := json.NewDecoder(strings.NewReader(string(trxData)))
	errD := dec.Decode(&response)

	if errD != nil {
		return nil, errors.NewAppError(nil, "cannot parse transaction response: "+string(trxData), -1, nil)
	}

	return &response, nil
}