# iyzipay-go

You can sign up for an iyzico account at https://iyzico.com

# Requirements

Go 1.7 or newer

# Installation

To install iyzipay, use `go get`:

    go get github.com/iyzico/iyzipay-go/iyzipay

Import the `github.com/iyzico/iyzipay-go/iyzipay` package into your code

# Usage

```go
    api_key := "sandbox-afXhZPW0MQlE4dCUUlHcEopnMBgXnAZI"
	secret_key := "sandbox-wbwpzKIiplZxI3hh5ALI4FJyAcZKL6kq"
	base_url := "https://sandbox-api.iyzipay.com"

	options := iyzipay.Options{}
	options.New(api_key, secret_key, base_url)

	paymentCard := iyzipay.PaymentCard{
		CardHolderName: "John Doe",
		CardNumber:     "5528790000000008",
		ExpireMonth:    "12",
		ExpireYear:     "2030",
		Cvc:            "123",
		RegisterCard:   "0",
	}

	buyer := iyzipay.Buyer{
		Id:                  "BY789",
		Name:                "John",
		Surname:             "Doe",
		IdentityNumber:      "74300864791",
		Email:               "email@email.com",
		GsmNumber:           "+905350000000",
		RegistrationDate:    "2013-04-21 15:12:09",
		LastLoginDate:       "2015-10-05 12:43:35",
		RegistrationAddress: "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
		City:                "Istanbul",
		Country:             "Turkey",
		ZipCode:             "34732",
		Ip:                  "::1",
	}

	address := iyzipay.Address{
		ContactName: "Jane Doe",
		City:        "Istanbul",
		Country:     "Turkey",
		Address:     "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
		ZipCode:     "34742",
	}

	basketItems := []iyzipay.BasketItem{
		iyzipay.BasketItem{
			Id:        "BI101",
			Name:      "Binocular",
			Category1: "Collectibles",
			Category2: "Accessories",
			ItemType:  "PHYSICAL",
			Price:     "0.3",
		},
		iyzipay.BasketItem{
			Id:        "BI102",
			Name:      "Game code",
			Category1: "Game",
			Category2: "Online Game Items",
			ItemType:  "VIRTUAL",
			Price:     "0.5",
		},
		iyzipay.BasketItem{
			Id:        "BI103",
			Name:      "Usb",
			Category1: "Electronics",
			Category2: "Usb / Cable",
			ItemType:  "PHYSICAL",
			Price:     "0.2",
		},
	}

	request := iyzipay.CreatePaymentRequest{
		Locale:          "tr",
		ConversationId:  "123456789",
		Price:           "1",
		PaidPrice:       "1.2",
		BasketId:        "B67832",
		PaymentGroup:    "LISTING",
		PaymentCard:     paymentCard,
		Currency:        "TRY",
		Buyer:           buyer,
		ShippingAddress: address,
		BillingAddress:  address,
		BasketItems:     basketItems,
	}

	payment := iyzipay.Payment{}.Create(request, options)
```

See other examples under `examples` folder.

### Testable examples

Run examples: `go test -v ./examples`

Or to run an individual example file: `go test -v ./examples/create_payment_test.go`

### Mock test cards

Test cards that can be used to simulate a *successful* payment:

Card Number      | Bank                       | Card Type
-----------      | ----                       | ---------
5890040000000016 | Akbank                     | Master Card (Debit)
5526080000000006 | Akbank                     | Master Card (Credit)
4766620000000001 | Denizbank                  | Visa (Debit)
4603450000000000 | Denizbank                  | Visa (Credit)
4729150000000005 | Denizbank Bonus            | Visa (Credit)
4987490000000002 | Finansbank                 | Visa (Debit)
5311570000000005 | Finansbank                 | Master Card (Credit)
9792020000000001 | Finansbank                 | Troy (Debit)
9792030000000000 | Finansbank                 | Troy (Credit)
5170410000000004 | Garanti Bankası            | Master Card (Debit)
5400360000000003 | Garanti Bankası            | Master Card (Credit)
374427000000003  | Garanti Bankası            | American Express
4475050000000003 | Halkbank                   | Visa (Debit)
5528790000000008 | Halkbank                   | Master Card (Credit)
4059030000000009 | HSBC Bank                  | Visa (Debit)
5504720000000003 | HSBC Bank                  | Master Card (Credit)
5892830000000000 | Türkiye İş Bankası         | Master Card (Debit)
4543590000000006 | Türkiye İş Bankası         | Visa (Credit)
4910050000000006 | Vakıfbank                  | Visa (Debit)
4157920000000002 | Vakıfbank                  | Visa (Credit)
5168880000000002 | Yapı ve Kredi Bankası      | Master Card (Debit)
5451030000000000 | Yapı ve Kredi Bankası      | Master Card (Credit)

*Cross border* test cards:

Card Number      | Country
-----------      | -------
4054180000000007 | Non-Turkish (Debit)
5400010000000004 | Non-Turkish (Credit)

Test cards to get specific *error* codes:

Card Number       | Description
-----------       | -----------
5406670000000009  | Success but cannot be cancelled, refund or post auth
4111111111111129  | Not sufficient funds
4129111111111111  | Do not honour
4128111111111112  | Invalid transaction
4127111111111113  | Lost card
4126111111111114  | Stolen card
4125111111111115  | Expired card
4124111111111116  | Invalid cvc2
4123111111111117  | Not permitted to card holder
4122111111111118  | Not permitted to terminal
4121111111111119  | Fraud suspect
4120111111111110  | Pickup card
4130111111111118  | General error
4131111111111117  | Success but mdStatus is 0
4141111111111115  | Success but mdStatus is 4
4151111111111112  | 3dsecure initialize failed
