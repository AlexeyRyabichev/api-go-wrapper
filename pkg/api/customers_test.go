package api

import (
	"context"
	"testing"
)

//works
func TestCustomerManager(t *testing.T) {
	const (
		//fill your data here
		sk                    = ""
		cc                    = ""
		someAvailableUsername = ""
	)
	//and here
	var (
		testingCustomer = &CustomerRequest{
			CompanyName: "",
			Username:    "",
			Password:    "",
		}
		ctx = context.Background()
	)

	cli, err := NewClient(sk, cc, nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Run("test get customers", func(t *testing.T) {
		resp, err := cli.GetCustomers(ctx, map[string]string{})
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(resp)
	})
	//works
	t.Run("test post customer", func(t *testing.T) {

		report, err := cli.PostCustomer(ctx, testingCustomer)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(report)
	})
	t.Run("test verifyCustomerUser", func(t *testing.T) {

		isAvailable, err := cli.VerifyCustomerUser(ctx, testingCustomer.Username, testingCustomer.Password)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(isAvailable)
	})
	t.Run("test validation of the username", func(t *testing.T) {
		isAvailable, err := cli.ValidateCustomerUsername(ctx, someAvailableUsername)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(isAvailable)
	})
}
