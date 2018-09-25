/**
 *  @copyright defined in chopsticks/LICENSE.txt
 *  @author Romain Pellerin - romain@eminent.ly
 *
 *  Donation appreciated :)
 *
 *  Bitcoin Cash $BCH wallet: 1HrhBfFRFovHv8EMxsuB9EcZgamtuH3fMc
 */
 package model

import "github.com/btcsuite/btcd/wire"

type Transaction struct {
	TxId              	string 			`json:"txid,omitempty"`
	SourceAddress      	string 			`json:"source_address,omitempty"`
	DestinationAddress 	string 			`json:"destination_address,omitempty"`
	Amount             	int64  			`json:"amount,omitempty"`
	UnsignedTx         	string 			`json:"unsignedtx,omitempty"`
	SignedTx           	string 			`json:"signedtx,omitempty"`
	MsgTx			   	*wire.MsgTx	 	`json:"-"`
	TxHex 				string			`json:"txhex,omitempty"`
	Hash 				string			`json:"txhash,omitempty"`
	Type				string			`json:"blockchain_type,omitempty"`
	BlockHeight 		int				`json:"block_height,omitempty"`
}

type TransactionRequest struct {
	TxHex 				string			`json:"txhex"`
	Blockchains 		[]string	 	`json:"blockchains,omitempty"`
	Voting				bool			`json:"voting,omitempty"`
}


type TransactionResponse struct {
	TxHex 				string			`json:"txhex"`
	Blockchains		 	[]Transaction	`json:"blockchains,omitempty"`
	Vote				Vote			`json:"vote,omitempty"`
	VoteSignature		string			`json:"vote_signature,omitempty"`
}

type Transactions struct {
	Transactions 	[]Transaction
}
