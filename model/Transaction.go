/**
 *  @copyright defined in chopsticks/LICENSE.txt
 *  @author Romain Pellerin - romain@eminent.ly
 *
 *  Donation appreciated :)
 *
 *  Bitcoin Cash $BCH wallet: 1HrhBfFRFovHv8EMxsuB9EcZgamtuH3fMc
 */
 package model

type Transaction struct {
	DataVersion int              `json:"dataVersion"`
	Uid         string           `json:"uid"`
	Created     int64            `json:"created"`
	UserId      string           `json:"userId"`
	Inputs      map[string]int64 `json:"source_address,omitempty"`
	Outputs     map[string]int64 `json:"destination_address,omitempty"`
	Amount      int64            `json:"amount,omitempty"`
	SignedTx    string           `json:"signedtx,omitempty"`
	Hash        string           `json:"txhash,omitempty"`
	ChainType   string           `json:"blockchain_type,omitempty"`
	BlockHeight int              `json:"block_height,omitempty"`
	Status      int              `json:"status,omitempty"`
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
