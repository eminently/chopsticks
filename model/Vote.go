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
	"context"
	"encoding/json"
	"github.com/eminently/chopsticks/common"
	. "github.com/eminently/chopsticks/errors"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog"
)

const VOTE_REGISTERED = 1
const VOTE_VALIDATED = 2
const VOTE_SUSPENDED = 3
const VOTE_DELETED = 4

const VOTE_DATA_VERSION = 1

type Vote struct {
	DataVersion         int 			`json:"dataVersion,omitempty"`
	Uid                 string 			`json:"uid,omitempty"`
	Created			 	int64			`json:"timestamp,omitempty"`
	UserId     			string			`json:"userId,omitempty"`
	PreferredChains 	map[string]int	`json:"preferredChains,omitempty"`
	Signature			string			`json:"signature,omitempty"`
	Status				int				`json:"status,omitempty"`
}

var (
	loggerVote zerolog.Logger
)

func VoteInit() {
	loggerVote = common.Logger("model.vote", context.Background())
}

func VoteLoadSettings() {}

func NewVote() Vote {
	return Vote{
		VOTE_DATA_VERSION,
		"",
		0,
		"",
		map[string]int{},
		"",
		0,
	}
}

// encoding/decoding methods

func VoteToJSON(obj *Vote) (string, error) {

	result := Vote{}

	copier.Copy(&result, &obj)

	bytes, err := json.Marshal(&result)

	return string(bytes), err
}

func JSONtoVote(obj string) (Vote, error) {

	s := NewVote()

	err := json.Unmarshal([]byte(obj), &s)

	return s, err
}

func VoteArrayToJSONArray(objs []Vote) []string {

	n := len(objs)

	// marshall skills array to json array
	var bin []string
	bin = make([]string, n)

	for i := 0; i < n; i++ {
		json, err := VoteToJSON(&objs[i])
		if err != nil {
			PanicOnAppError(NewAppError(err, err.Error(), -1, nil))
		}
		bin[i] = json
	}

	return bin
}

func JSONArrayToVoteArray(objs []string) []Vote {

	n := len(objs)

	// marshall objs array to json array
	var bin []Vote
	bin = make([]Vote, n)

	for i := 0; i < n; i++ {
		vote, err := JSONtoVote(objs[i])
		if err != nil {
			PanicOnAppError(NewAppError(err, err.Error(), -1, nil))
		}
		bin[i] = vote
	}

	return bin
}
