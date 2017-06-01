package qcxclient

import validator "gopkg.in/go-playground/validator.v9"

const Qcxapibase string = "https://api.quadrigacx.com/v2/"
const OrderBookurl string = "ticker?book=eth_cad"
const Transactionsurl string = "transactions?book=eth_cad&time=hour"
const Balanceurl string = "balance"
const OpenOrdersurl string = "open_orders"
const CancelOrderurl string = "cancel_order"

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

type Balanceresp struct {
	CAD string `json:"cad_balance" validate:"required"`
	ETH string `json:"eth_balance" validate:"required"`
	FEE string `json:"fee" validate:"required"`
}

type Balancepayload struct {
	Key       string `json:"key" validate:"required"`
	Signature string `json:"signature" validate:"required"`
	Nonce     string `json:"nonce" validate:"required"`
}

type OpenOrder struct {
	ID       string `json:"id" validate:"required"`
	Datetime string `json:"datetime" validate:"required"`
	Type     string `json:"type" validate:"required"`
	Price    string `json:"price" validate:"required"`
	Amount   string `json:"amount" validate:"required"`
	Status   string `json:"status" validate:"required"`
}

type OpenOrders struct {
	OpenOrders []OpenOrder `validate:"required"`
}

type OpenOrderspayload struct {
	Key       string `json:"key" validate:"required"`
	Signature string `json:"signature" validate:"required"`
	Nonce     string `json:"nonce" validate:"required"`
	Book      string `json:"book" validate:"required"`
}

type CancelOrderpayload struct {
	Key       string `json:"key" validate:"required"`
	Signature string `json:"signature" validate:"required"`
	Nonce     string `json:"nonce" validate:"required"`
	OrderID   string `json:"id" validate:"required"`
}
