/**
 *  @copyright defined in chopsticks/LICENSE.txt
 *  @author Romain Pellerin - romain@eminent.ly
 *
 *  Donation appreciated :)
 *
 *  Bitcoin Cash $BCH wallet: 1HrhBfFRFovHv8EMxsuB9EcZgamtuH3fMc
 */
 package network

import (
	"encoding/json"
	"net/http"
	"strings"
	"io/ioutil"
	"net/url"
	"github.com/eminently/chopsticks/errors"
)

// simple HTTP GET request taking some URL params
func Get(_url string, _params map[string]string, token string) ([]byte, *errors.AppError) {

	baseUrl, err := url.Parse(_url)

	if err != nil {
		return nil, errors.NewAppError(err, "cannot parse url", -1, nil)
	}

	params := url.Values{}

	for k, v := range _params {

		u := &url.URL{Path: k}
		k = u.String()

		u = &url.URL{Path: v}
		v = u.String()

		params.Add(k, v)
	}

	baseUrl.RawQuery = params.Encode()

	//fmt.Println("network.HTTP GET URL: ", baseUrl.String())

	req, _ := http.NewRequest("GET", baseUrl.String(), strings.NewReader(""))

	req.Header.Add("Authorization", "User "+token)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, errors.NewAppError(err, "error trying to reach  API", -1, nil)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, errors.NewAppError(err, "error reading body response", -1, nil)
	}

	httpError := errors.HTTPError{}
	json.Unmarshal(body, &httpError)

	if httpError.Code != 0 {
		errStr, _ := errors.HTTPErrorTOJSON(httpError)
		return nil, errors.NewAppError(nil, " API returned an error: "+errStr, int64(httpError.Code), nil)
	}

	return body, nil
}

// simple HTTP JSON post request taking data as a map OR byte array as JSON body
// by default bytes will be overrided if keyValues is passed
func Post(url string, keyValues map[string]interface{}, bytes []byte) ([]byte, *errors.AppError) {

	//fmt.Println("post keyValues: ",keyValues)
	//fmt.Println("post raw: "+string(bytes))

	var err error

	if keyValues != nil {
		bytes, err = json.Marshal(&keyValues)
	}

	if err != nil {
		return nil, errors.NewAppError(err, "error marshalling params", -1, nil)
	}

	req, _ := http.NewRequest("POST", url, strings.NewReader(string(bytes)))

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, errors.NewAppError(err, "error trying to reach  API", -1, nil)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, errors.NewAppError(err, "error reading body response", -1, nil)
	}

	httpError := errors.HTTPError{}
	json.Unmarshal(body, &httpError)

	if httpError.Code != 0 {
		errStr, _ := errors.HTTPErrorTOJSON(httpError)
		return nil, errors.NewAppError(nil, " API returned an error: "+errStr, int64(httpError.Code), nil)
	}

	return body, nil
}

// simple HTTP JSON post request taking data as JSON body
func PostRawData(url string, raw string, token string) ([]byte, *errors.AppError) {

	//fmt.Println("post raw data: " + raw)

	// force quotes if not json or array object
	if !strings.HasPrefix(raw,"\"") && !strings.HasPrefix(raw,"{") && !strings.HasPrefix(raw,"[") {
		raw = "\""+raw+"\""
	}

	req, _ := http.NewRequest("POST", url, strings.NewReader(raw))

	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "User "+token)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, errors.NewAppError(err, "error trying to reach  API", -1, nil)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, errors.NewAppError(err, "error reading body response", -1, nil)
	}

	httpError := errors.HTTPError{}
	json.Unmarshal(body, &httpError)

	if httpError.Code != 0 {
		errStr, _ := errors.HTTPErrorTOJSON(httpError)
		return nil, errors.NewAppError(nil, "API returned an error: "+errStr, int64(httpError.Code), nil)
	}

	return body, nil
}
