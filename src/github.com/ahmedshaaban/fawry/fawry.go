package fawry

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const APIPath = "/ECommerceWeb/Fawry/payments/"
const BaseUrl = "https://www.atfawry.com"
const SandboxBaseUrl = "https://atfawry.fawrystaging.com"

type FawryClient struct {
	IsSandbox      bool
	FawrySecureKey string
}

func (fc FawryClient) getURL() string {
	if fc.IsSandbox {
		return SandboxBaseUrl + APIPath
	}
	return BaseUrl + APIPath
}

func (fc FawryClient) getSignature(inputs []string) string {
	sum := sha256.Sum256([]byte(strings.Join(inputs[:], ",")))
	return hex.EncodeToString(sum[:])
}

func (fc FawryClient) ChargeRequest(charge Charge) (*http.Response, error) {
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

	fmt.Println(string(jsonBytes))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (fc FawryClient) RefundRequest(refund Refund) (*http.Response, error) {
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

	fmt.Println(string(jsonBytes))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (fc FawryClient) StatusRequest(status Status) (*http.Response, error) {
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

	fmt.Println(req.URL.String())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
