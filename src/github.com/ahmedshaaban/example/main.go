package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ahmedshaaban/fawry"
)

func main() {
	fc := fawry.Client{
		IsSandbox:      true,
		FawrySecureKey: "ay haga",
	}

	charge := fawry.Charge{
		MerchantCode:      "is0N+YQzlE4=",
		MerchantRefNum:    "9990064204",
		CustomerProfileID: "9990064204",
		CustomerMobile:    "01000000200",
		CustomerEmail:     "77@test.com",
		PaymentMethod:     "PAYATFAWRY",
		Amount:            "20.10",
		CurrencyCode:      "EGP",
		Description:       "the charge request description",
		PaymentExpiry:     1516554874077,
		ChargeItems: []fawry.ChargeItem{
			fawry.ChargeItem{
				ItemID:      "897fa8e81be26df25db592e81c31c",
				Description: "asdasd",
				Price:       "15.20",
				Quantity:    1,
			},
		},
	}

	// refund := fawry.Refund{
	// 	MerchantCode:    "1013969",
	// 	ReferenceNumber: "322818",
	// 	RefundAmount:    "100.00",
	// 	Reason:          "Bad Quality ",
	// }

	// status := fawry.Status{
	// 	MerchantCode:   "is0N+YQzlE4=",
	// 	MerchantRefNum: "99900642041",
	// }

	resp, err := fc.ChargeRequest(charge)
	if err != nil {
		fmt.Println(err)
		return
	}

	// resp, err := fc.RefundRequest(refund)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// resp, err := fc.StatusRequest(status)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
