/**
 *  @copyright defined in chopsticks/LICENSE.txt
 *  @author Romain Pellerin - romain@eminent.ly
 *
 *  Donation appreciated :)
 *
 *  Bitcoin Cash $BCH wallet: 1HrhBfFRFovHv8EMxsuB9EcZgamtuH3fMc
 */
package errors

import (
	"encoding/json"
)

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error struct {
		Code int    `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
		Details []struct {
			Message    string `json:"message"`
			File       string `json:"file"`
			LineNumber int    `json:"line_number"`
			Method     string `json:"method"`
		} `json:"details"`
	} `json:"error"`
}

func HTTPErrorTOJSON(error HTTPError) (string, *AppError) {

	json, err := json.Marshal(error)

	if err != nil {
		return "", NewAppError(err, "error trying to marshal HTTPError", -1, nil)
	}

	return string(json), nil
}
