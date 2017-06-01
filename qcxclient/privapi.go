package qcxclient

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/spf13/viper"

	"fmt"

	"time"

	validator "gopkg.in/go-playground/validator.v9"
)

func getAccountBalance() (*Balanceresp, error) {
	apikey, nonce, sig := signtransaction()
	url := Qcxapibase + Balanceurl

	reqbody := Balancepayload{
		Key:       apikey,
		Signature: sig,
		Nonce:     nonce,
	}
	jsonstr, _ := json.Marshal(reqbody)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonstr))
	req.Header.Set("Content-Type", "application/json")

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

	var respstruct Balanceresp
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

func getOpenOrders() (*OpenOrders, error) {
	apikey, nonce, sig := signtransaction()
	url := Qcxapibase + OpenOrdersurl

	reqbody := OpenOrderspayload{
		Key:       apikey,
		Signature: sig,
		Nonce:     nonce,
		Book:      "eth_cad",
	}
	jsonstr, _ := json.Marshal(reqbody)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonstr))
	req.Header.Set("Content-Type", "application/json")

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
	ord := make([]OpenOrder, 0)
	err = json.Unmarshal(body, &ord)
	if err != nil {
		return nil, err
	}
	var ords OpenOrders
	ords.OpenOrders = ord

	validate = validator.New()

	err = validate.Struct(ords)
	if err != nil {
		return nil, err
	}

	return &ords, nil
}

func cancelOrder(orderid string) (bool, error) {
	apikey, nonce, sig := signtransaction()
	url := Qcxapibase + CancelOrderurl

	reqbody := CancelOrderpayload{
		Key:       apikey,
		Signature: sig,
		Nonce:     nonce,
		OrderID:   orderid,
	}
	jsonstr, _ := json.Marshal(reqbody)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonstr))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	if res.StatusCode != http.StatusOK {
		return false, fmt.Errorf("API error received status %s", res.Status)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	result, _ := strconv.ParseBool(string(body))

	return result, nil
}

func signtransaction() (string, string, string) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("sampleauth")
	viper.ReadInConfig()

	clientid := viper.GetString("qcx.clientid")
	apikey := viper.GetString("qcx.apikey")
	secret := viper.GetString("qcx.secret")
	nonce := strconv.FormatInt(time.Now().Unix(), 10)

	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(nonce + clientid + apikey))

	return apikey, nonce, hex.EncodeToString(h.Sum(nil))

}
