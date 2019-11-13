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

const API_PATH = "/ECommerceWeb/Fawry/payments/"
const BASE_URL = "https://www.atfawry.com"
const SANDBOX_BASE_URL = "https://atfawry.fawrystaging.com"

type FawryClient struct {
	IsSandbox      bool
	FawrySecureKey string
}

func (fc FawryClient) getUrl(isSandbox bool) string {
	if isSandbox {
		return SANDBOX_BASE_URL + API_PATH
	}
	return BASE_URL + API_PATH
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



	url := fc.getUrl(fc.IsSandbox) + "charge"

	signatureArray := []string{charge.MerchantCode,
		charge.MerchantRefNum, charge.CustomerProfileId,
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
