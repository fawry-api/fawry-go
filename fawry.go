package fawry

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const apiPath = "/ECommerceWeb/Fawry/payments/"
const baseURL = "https://www.atfawry.com"
const sandboxBaseURL = "https://atfawry.fawrystaging.com"

// ErrKeyMissing indicates security key missing
var ErrKeyMissing = errors.New("Fawry security key missing")

// Client Struct
type Client struct {
	IsSandbox      bool
	FawrySecureKey string
}

// NewClientFromEnv creates new client from FAWRY_SECURE_KEY environment
func NewClientFromEnv(sandbox bool) (*Client, error) {
	securityKey, ok := os.LookupEnv("FAWRY_SECURE_KEY")
	if !ok {
		return nil, ErrKeyMissing
	}
	return &Client{
		FawrySecureKey: securityKey,
		IsSandbox: sandbox,
	}
}

// NewClientFromEnv creates new client
func NewClient(securityKey string, sandbox bool) (*Client, error) {
	if securityKey == "" {
		return nil, ErrKeyMissing
	}
	
	return &Client{
		FawrySecureKey: securityKey,
		IsSandbox: sandbox,
	}
}

func (fc Client) getURL() string {
	if fc.IsSandbox {
		return sandboxBaseURL + apiPath
	}
	return baseURL + apiPath
}

func (fc Client) getSignature(inputs []string) string {
	sum := sha256.Sum256([]byte(strings.Join(inputs[:], ",")))
	return hex.EncodeToString(sum[:])
}

// ChargeRequest could be used to charge the customer with different payment methods.
// 	It also might be used to create a reference number to be paid at Fawry's outlets or
//	it can be used to direct debit the customer card using card token.
func (fc Client) ChargeRequest(charge Charge) (*http.Response, error) {
	err := charge.Validate()
	if err != nil {
		return nil, err
	}

	url := fc.getURL() + "charge"

	signatureArray := []string{charge.MerchantCode,
		charge.MerchantRefNum, charge.CustomerProfileID,
		charge.PaymentMethod, charge.Amount, charge.CardToken, fc.FawrySecureKey}

	jsonBytes, err := json.Marshal(struct {
		Charge
		Signature string `json:"signature"`
	}{Charge: charge,
		Signature: fc.getSignature(signatureArray)})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RefundRequest  can refund the payment again to the customer
func (fc Client) RefundRequest(refund Refund) (*http.Response, error) {
	err := refund.Validate()
	if err != nil {
		return nil, err
	}

	url := fc.getURL() + "refund"

	signatureArray := []string{refund.MerchantCode,
		refund.ReferenceNumber, refund.RefundAmount,
		refund.Reason, fc.FawrySecureKey}

	jsonBytes, err := json.Marshal(struct {
		Refund
		Signature string `json:"signature"`
	}{Refund: refund,
		Signature: fc.getSignature(signatureArray)})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// StatusRequest can use Get Payment Status Service to retrieve the payment status for the charge request
func (fc Client) StatusRequest(status Status) (*http.Response, error) {
	err := status.Validate()
	if err != nil {
		return nil, err
	}

	signatureArray := []string{status.MerchantCode, status.MerchantRefNum, fc.FawrySecureKey}

	url := fc.getURL() + "status" + fmt.Sprintf("?merchantCode=%s&merchantRefNumber=%s&signature=%s",
		status.MerchantCode,
		status.MerchantRefNum,
		fc.getSignature(signatureArray))

	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
