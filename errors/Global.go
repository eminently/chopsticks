/**
 *  @copyright defined in chopsticks/LICENSE.txt
 *  @author Romain Pellerin - romain@eminent.ly
 *
 *  Donation appreciated :)
 *
 *  Bitcoin Cash $BCH wallet: 1HrhBfFRFovHv8EMxsuB9EcZgamtuH3fMc
 */
package errors

import "fmt"

type AppError struct {
	Error   error
	Message string
	Code    int
	Custom  interface{}
}

func NewAppError(error error, message string, code int64, custom interface{}) *AppError {
	return &AppError{
		error,
		message,
		int(code),
		custom,
	}
}

func MarshallingError(err error) *AppError {
	return NewAppError(err, "", -1, nil)
}

func UnsupportedOperation() *AppError {
	return &AppError{nil, "unsupported operation", 14, nil}
}

func PanicOnAppError(err *AppError) {
	if err.Error != nil {
		fmt.Printf("app error: %d", err)
		panic(err)
	} else {
		fmt.Printf("app error: %s", err.Message)
	}
}