package qcxclient

import "testing"

func TestOrderBook(t *testing.T) {
	_, err := getOrderBook()
	if err != nil {
		t.Fail()
	}
}

func TestTransactions(t *testing.T) {
	_, err := getTransactions()
	if err != nil {
		t.Fail()
	}
}
