package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"math"
	"net/http"

	"github.com/System-Glitch/goyave/v3/config"
)

// MakePayment makes tries to make payment request to the external payment API.
func MakePayment(username string, amount float64) (paymentID string, err error) {
	paymentURL := config.GetString("app.paymentAPI")
	paymentType := "DEBIT"
	if amount > 0 {
		paymentType = "CREDIT"
	}
	payload, err := json.Marshal(map[string]interface{}{
		"user_name":    username,
		"payment_type": paymentType,
		"amount":       math.Abs(float64(amount)),
	})
	if err != nil {
		return
	}
	resp, err := http.Post(paymentURL, "application/json", bytes.NewReader(payload))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = errors.New("PaymentError1")
		return
	}
	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return
	}
	if res["status"] == "FAILIURE" {
		err = errors.New("PaymentError2")
		return
	}
	paymentID = res["payment_id"].(string)
	return
}
