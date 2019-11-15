# fawry-go

[![Build Status](https://travis-ci.com/ahmedshaaban/fawry-go.svg?branch=master)](https://travis-ci.com/ahmedshaaban/fawry-go)
[![Coverage Status](https://coveralls.io/repos/github/ahmedshaaban/fawry-go/badge.svg?branch=master)](https://coveralls.io/github/ahmedshaaban/fawry-go?branch=master)


## Description

fawry-go is a Go package interfacing with Fawry's payment gateway API. this package is inspired by Amr Bakry's ruby [gem](https://github.com/fawry-api/fawry "fawry-ruby§")

_**`important note`**_: You need to have a contract with fawry to use their service.

## Requirements

Go 1.8 or above.

## Getting Started

### Installation

Run the following command to install the package:
```
go get github.com/ahmedshaaban/fawry-go
```

### Charge Request

```go
package main

import (
	"fmt"
    "io/ioutil"

	"github.com/ahmedshaaban/fawry-go"
)

func main() {
	fc := fawry.Client{
		IsSandbox:      true,
		FawrySecureKey: "SecuredKeyProvidedByFawry",
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
				Description: "lorem",
				Price:       "15.20",
				Quantity:    1,
			},
		},
    }
    
    resp, err := fc.ChargeRequest(charge)
	if err != nil {
		fmt.Println(err)
		return
    }

    defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
    
```

### Refund Request

```go
package main

import (
	"fmt"
    "io/ioutil"

	"github.com/ahmedshaaban/fawry-go"
)

func main() {
	fc := fawry.Client{
		IsSandbox:      true,
		FawrySecureKey: "SecuredKeyProvidedByFawry",
	}

	refund := fawry.Refund{
		MerchantCode:    "1013969",
		ReferenceNumber: "322818",
		RefundAmount:    "100.00",
		Reason:          "Bad Quality ",
	}

    resp, err := fc.RefundRequest(refund)
	if err != nil {
		fmt.Println(err)
		return
	}

    defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
    
```

### Status Request

```go
package main

import (
	"fmt"
    "io/ioutil"

	"github.com/ahmedshaaban/fawry-go"
)

func main() {
	fc := fawry.Client{
		IsSandbox:      true,
		FawrySecureKey: "SecuredKeyProvidedByFawry",
	}

	status := fawry.Status{
		MerchantCode:   "is0N+YQzlE4=",
		MerchantRefNum: "99900642041",
	}

    resp, err := fc.StatusRequest(refund)
	if err != nil {
		fmt.Println(err)
		return
	}

    defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
    
```

## TODO:
- Read configuration keys (merchant code, secure key) from env vars
- Add public API documentation to README
- Increase code coverage
