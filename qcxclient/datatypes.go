package qcxclient

import validator "gopkg.in/go-playground/validator.v9"

const Qcxapibase string = "https://api.quadrigacx.com/v2/"
const OrderBookurl string = "ticker?book=eth_cad"
const Transactionsurl string = "transactions?book=eth_cad&time=hour"

var validate *validator.Validate

type OrderBookresp struct {
	High      string `json:"high" validate:"required"`
	Last      string `json:"last" validate:"required"`
	Timestamp string `json:"timestamp" validate:"required"`
	Volume    string `json:"volume" validate:"required"`
	Vwap      string `json:"vwap" validate:"required"`
	Low       string `json:"low" validate:"required"`
	Ask       string `json:"ask" validate:"required"`
	Bid       string `json:"bid" validate:"required"`
}

type Transaction struct {
	Date   string `json:"date" validate:"required"`
	Tid    int    `json:"tid" validate:"required"`
	Price  string `json:"price" validate:"required"`
	Amount string `json:"amount" validate:"required"`
	Side   string `json:"side" validate:"required"`
}

type Transactions struct {
	Transactions []Transaction `validate:"required"`
}
