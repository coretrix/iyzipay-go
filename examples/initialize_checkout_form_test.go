package iyzipay_examples

import (
	"encoding/json"
	"fmt"
	builder "github.com/kurtulussahin/iyzipay-go/examples/builder"
	iyzipay "github.com/kurtulussahin/iyzipay-go/iyzipay"
)

func ExampleInitializeCheckoutForm() {
	options := iyzipay.Options{}
	options = builder.CreateSampleOptions()

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

	request := iyzipay.CreateCheckoutFormInitializeRequest{
		Locale:              "tr",
		ConversationId:      "123456789",
		Price:               "1",
		PaidPrice:           "1.2",
		BasketId:            "B67832",
		PaymentGroup:        "LISTING",
		Currency:            "TRY",
		Buyer:               buyer,
		ShippingAddress:     address,
		BillingAddress:      address,
		BasketItems:         basketItems,
		CallbackUrl:         "https://iyzi-sandbox-kurtulussahin.c9users.io/iyzipay-php-develop/samples/CheckoutFormAuthSample.php",
		EnabledInstallments: []string{"2", "3", "6", "9"},
	}

	rawResponse := iyzipay.CheckoutFormInitialize{}.Create(request, options)

	var response map[string]interface{}
	json.Unmarshal([]byte(rawResponse), &response)

	fmt.Println(response["status"])
	fmt.Println(response["locale"])
	fmt.Println(response["conversationId"])
	fmt.Println(response["errorCode"])
	fmt.Println(response["errorMessage"])
	fmt.Println(response["errorGroup"])
	//Output:
	//success
	//tr
	//123456789
	//<nil>
	//<nil>
	//<nil>
}
