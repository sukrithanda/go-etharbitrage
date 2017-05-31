package qcxclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fmt"

	validator "gopkg.in/go-playground/validator.v9"
)

func getOrderBook() (*OrderBookresp, error) {

	url := Qcxapibase + OrderBookurl

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error received status %s", res.Status)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var respstruct OrderBookresp
	err = json.Unmarshal(body, &respstruct)
	if err != nil {
		return nil, err
	}

	validate = validator.New()

	err = validate.Struct(respstruct)
	if err != nil {
		return nil, err
	}

	return &respstruct, nil
}

func getTransactions() (*Transactions, error) {

	url := Qcxapibase + Transactionsurl

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error received status %s", res.Status)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	tx := make([]Transaction, 0)
	err = json.Unmarshal(body, &tx)
	if err != nil {
		return nil, err
	}
	var txs Transactions
	txs.Transactions = tx

	validate = validator.New()

	err = validate.Struct(txs)
	if err != nil {
		return nil, err
	}

	return &txs, nil
}
