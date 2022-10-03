package iyzipay_examples

import (
	"encoding/json"
	"fmt"
	builder "github.com/kurtulussahin/iyzipay-go/examples/builder"
	iyzipay "github.com/kurtulussahin/iyzipay-go/iyzipay"
)

func ExampleUpdatePersonalSubMerchant() {
	options := iyzipay.Options{}
	options = builder.CreateSampleOptions()

	request := iyzipay.UpdateSubMerchantRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		SubMerchantKey: "bGf96xpbXusSTaW0fg/YCbdhHIg=",
		Address:        "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
		ContactName:    "John",
		ContactSurname: "Doe",
		Email:          "email@submerchantemail.com",
		GsmNumber:      "+905350000000",
		Name:           "John's market",
		Iban:           "TR180006200119000006672315",
		IdentityNumber: "31300864726",
		Currency:       "TRY",
	}

	rawResponse := iyzipay.SubMerchant{}.Update(request, options)

	var response map[string]interface{}
	json.Unmarshal([]byte(rawResponse), &response)

	fmt.Println(response["status"])
	fmt.Println(response["locale"])
	fmt.Println(response["conversationId"])
	fmt.Println(response["errorMessage"])
	//Output:
	//success
	//tr
	//123456789
	//<nil>
}

func ExampleUpdatePrivateSubMerchant() {
	options := iyzipay.Options{}
	options = builder.CreateSampleOptions()

	request := iyzipay.UpdateSubMerchantRequest{
		Locale:            "tr",
		ConversationId:    "123456789",
		SubMerchantKey:    "eXfZayHK/kXpZWJwkaDujkAFNeQ=",
		Address:           "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
		TaxOffice:         "Tax office",
		LegalCompanyTitle: "John Doe inc",
		Email:             "email@submerchantemail.com",
		GsmNumber:         "+905350000000",
		Name:              "John's market",
		Iban:              "TR180006200119000006672315",
		IdentityNumber:    "31300864726",
		Currency:          "TRY",
	}

	rawResponse := iyzipay.SubMerchant{}.Update(request, options)

	var response map[string]interface{}
	json.Unmarshal([]byte(rawResponse), &response)

	fmt.Println(response["status"])
	fmt.Println(response["locale"])
	fmt.Println(response["conversationId"])
	fmt.Println(response["errorMessage"])
	//Output:
	//success
	//tr
	//123456789
	//<nil>
}

func ExampleUpdateLimitedCompanySubMerchant() {
	options := iyzipay.Options{}
	options = builder.CreateSampleOptions()

	request := iyzipay.UpdateSubMerchantRequest{
		Locale:            "tr",
		ConversationId:    "123456789",
		SubMerchantKey:    "YJz6b789M3t5imkUTyDns8JsDeQ=",
		Address:           "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
		TaxOffice:         "Tax office",
		TaxNumber:         "9261877",
		LegalCompanyTitle: "John Doe inc",
		Email:             "email@submerchantemail.com",
		GsmNumber:         "+905350000000",
		Name:              "John's market",
		Iban:              "TR180006200119000006672315",
		IdentityNumber:    "31300864726",
		Currency:          "TRY",
	}

	rawResponse := iyzipay.SubMerchant{}.Update(request, options)

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
