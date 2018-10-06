/**
 *  @copyright defined in chopsticks/LICENSE.txt
 *  @author Romain Pellerin - romain@eminent.ly
 *
 *  Donation appreciated :)
 *
 *  Bitcoin Cash $BCH wallet: 1HrhBfFRFovHv8EMxsuB9EcZgamtuH3fMc
 */
package model

import "github.com/chopsticks/errors"

type Transaction struct {
	DataVersion  int              `json:"data_version"`
	Uid          string           `json:"uid"`
	Created      int64            `json:"created"`
	UserId       string           `json:"user_id"`
	Inputs       map[string]int64 `json:"inputs,omitempty"`
	Outputs      map[string]int64 `json:"outputs,omitempty"`
	Amount       int64            `json:"amount,omitempty"`
	SignedTx     string           `json:"tx_signed,omitempty"`
	Hash         string           `json:"tx_hash,omitempty"`
	ChainType    string           `json:"blockchain_type,omitempty"`
	ChainVersion string           `json:"blockchain_version,omitempty"`
	BlockHeight  int              `json:"block_height,omitempty"`
	Status       int              `json:"status,omitempty"`
}

type TransactionRequest struct {
	TxHex       string   `json:"tx_hex,omitempty"`
	Blockchains []string `json:"blockchains,omitempty"`
	Voting      bool     `json:"voting,omitempty"`
}

type TransactionResponse struct {
	TxHex         string            `json:"tx_hex,omitempty""`
	Transactions  []Transaction     `json:"transactions,omitempty"`
	Vote          Vote              `json:"vote,omitempty"`
	VoteSignature string            `json:"vote_signature,omitempty"`
	Errors        []errors.AppError `json:"errors,omitempty"`
}

type Transactions struct {
	Transactions []Transaction
}
