package fawry

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

const API_PATH = "/ECommerceWeb/Fawry/payments/"
const BASE_URL = "https://www.atfawry.com"
const SANDBOX_BASE_URL = "https://atfawry.fawrystaging.com"

func getUrl(isSandbox bool) string {
	if isSandbox {
		return SANDBOX_BASE_URL + API_PATH
	}
	return BASE_URL + API_PATH
}

func getSignature(inputs []string) string {
	sum := sha256.Sum256([]byte(strings.Join(inputs[:], ",")))
	return hex.EncodeToString(sum[:])
}

func ChargeRequest(charge Charge) {
	err := charge.Validate()
	// to be continued
	fmt.Println(err)
}

func Hello() {
	charge := Charge{
		MerchantCode: "ahmed",
	}
	ChargeRequest(charge)

	fmt.Println("Hello, World!")
}
