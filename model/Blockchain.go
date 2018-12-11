/**
 *  @copyright defined in chopsticks/LICENSE.txt
 *  @author Romain Pellerin - romain@eminent.ly
 *
 *  Donation appreciated :)
 *
 *  Bitcoin Cash $BCH wallet: 1HrhBfFRFovHv8EMxsuB9EcZgamtuH3fMc
 */
package model

import (
	"github.com/eminently/chopsticks/common"
	"github.com/eminently/chopsticks/errors"
	"github.com/gcash/bchd/btcjson"
)

// types of blockchain supported
const BLOCKCHAIN_TYPE_XBC_TESTNET = "TXBC" // Bitcoin ABC testnet

const BLOCKCHAIN_TYPE_XBC_MAINNET = "XBC" // Bitcoin ABC
const BLOCKCHAIN_TYPE_XBS_MAINNET = "XBS" // Bitcoin BSV
const BLOCKCHAIN_TYPE_XBN_MAINNET = "XBN" // Bitcoin NayBC
const BLOCKCHAIN_TYPE_XBU_MAINNET = "XBU" // Bitcoin Unlimited
const BLOCKCHAIN_TYPE_XBX_MAINNET = "XBX" // Bitcoin XT
const BLOCKCHAIN_TYPE_XBD_MAINNET = "XBD" // Gash bchd
const BLOCKCHAIN_TYPE_XBB_MAINNET = "XBB" // Bcoin bcash

// structs

type Mempool struct {
	Transactions	map[string]btcjson.GetRawMempoolVerboseResult	`json:"transactions,omitempty"`
	ChainType   	string           								`json:"blockchain_type,omitempty"`
	ChainVersion 	string           								`json:"blockchain_version,omitempty"`
}

type MempoolsResponse struct {
	Mempools	[]Mempool			`json:"mempools,omitempty"`
	Errors  	[]errors.AppError	`json:"errors,omitempty"`
}

type MiningInfo struct {
	MiningInfo		btcjson.GetMiningInfoResult	`json:"mining_info,omitempty"`
	ChainType   	string           			`json:"blockchain_type,omitempty"`
	ChainVersion 	string           			`json:"blockchain_version,omitempty"`
}

type MiningInfosResponse struct {
	MiningInfos	[]MiningInfo		`json:"mining_infos,omitempty"`
	Errors  	[]errors.AppError	`json:"errors,omitempty"`
}

type Info struct {
	Info			btcjson.GetBlockChainInfoResult	`json:"info,omitempty"`
	ChainType   	string           				`json:"blockchain_type,omitempty"`
	ChainVersion 	string           				`json:"blockchain_version,omitempty"`
}

type InfosResponse struct {
	Infos	[]Info				`json:"infos,omitempty"`
	Errors  []errors.AppError	`json:"errors,omitempty"`
}

type Block struct {
	Height			int						`json:"height,omitempty"`
	Ecmh			string					`json:"ecmh,omitempty"`
	EcmhColor  		common.RGBT				`json:"ecmh_color,omitempty"`
	EcmhSimpleValue	int						`json:"ecmh_simple_value,omitempty"`
	Timestamp 		int64					`json:"timestamp,omitempty"`
	TransactionsNb	int						`json:"transactions_nb,omitempty"`
	Transactions	map[string]Transaction	`json:"transactions,omitempty"`
}

type BlocksResponse struct {
	Blocks	map[string]map[int64]Block	`json:"blocks,omitempty"`
	Errors  []errors.AppError			`json:"errors,omitempty"`
}