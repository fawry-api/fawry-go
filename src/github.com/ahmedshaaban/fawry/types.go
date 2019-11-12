package fawry

import (
	"regexp"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ChargeItem struct {
	ItemID      string `json:"itemId"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Quantity    int    `json:"quantity"`
}

func (chargeItem ChargeItem) Validate() error {
	return validation.ValidateStruct(&chargeItem,
		validation.Field(&chargeItem.ItemID, validation.Required),
		validation.Field(&chargeItem.Description, validation.Required),
		validation.Field(&chargeItem.Price, validation.Required, validation.Match(regexp.MustCompile(`^\d+\.\d\d$`))),
		validation.Field(&chargeItem.Quantity, validation.Required),
	)
}

type Charge struct {
	MerchantCode      string       `json:"merchantCode"`
	MerchantRefNum    string       `json:"merchantRefNum"`
	CustomerProfileId string       `json:"customerProfileId"`
	Amount            string       `json:"amount"`
	Description       string       `json:"description"`
	CustomerMobile    string       `json:"customerMobile"`
	ChargeItems       []ChargeItem `json:"chargeItems"`
	CurrencyCode      string       `json:"currencyCode"`
	CardToken         string       `json:"cardToken"`
	CustomerEmail     string       `json:"customerEmail"`
	PaymentMethod     string       `json:"paymentMethod"`
	PaymentExpiry     int          `json:"paymentExpiry"`
}

func (charge Charge) Validate() error {
	return validation.ValidateStruct(&charge,
		validation.Field(&charge.MerchantCode, validation.Required),
		validation.Field(&charge.MerchantRefNum, validation.Required),
		validation.Field(&charge.CustomerProfileId, validation.Required),
		validation.Field(&charge.Amount, validation.Required, validation.Match(regexp.MustCompile(`^\d+\.\d\d$`))),
		validation.Field(&charge.Description, validation.Required),
		validation.Field(&charge.CustomerMobile, validation.Required),
		validation.Field(&charge.ChargeItems),
		validation.Field(&charge.CurrencyCode),
		validation.Field(&charge.CardToken),
		validation.Field(&charge.CustomerEmail, is.Email),
		validation.Field(&charge.PaymentMethod),
		validation.Field(&charge.PaymentExpiry),
	)
}
