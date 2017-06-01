package qcxclient

import "testing"

func TestBalance(t *testing.T) {
	_, err := getAccountBalance()
	if err != nil {
		t.Fail()
	}
}
func TestOpenOrders(t *testing.T) {
	_, err := getOpenOrders()
	if err != nil {
		t.Fail()
	}
}

func TestCancelOrder(t *testing.T) {
	_, err := cancelOrder("someid")
	if err != nil {
		t.Fail()
	}
}
