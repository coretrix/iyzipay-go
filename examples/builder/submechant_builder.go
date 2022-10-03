package builder

import (
	"encoding/json"
	iyzipay "github.com/kurtulussahin/iyzipay-go/iyzipay"
)

func CreateSubmerchantBuilder() map[string]interface{} {
	options := iyzipay.Options{}
	options = CreateSampleOptions()

	request := iyzipay.CreateSubMerchantRequest{
		Locale:                "tr",
		ConversationId:        "123456789",
		SubMerchantExternalId: "submerchant" + iyzipay.RandString(8),
		SubMerchantType:       "PERSONAL",
		Address:               "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
		ContactName:           "John",
		ContactSurname:        "Doe",
		Email:                 "email@submerchantemail.com",
		GsmNumber:             "+905350000000",
		Name:                  "John's market",
		Iban:                  "TR180006200119000006672315",
		IdentityNumber:        "31300864726",
		Currency:              "TRY",
	}

	rawResponse := iyzipay.SubMerchant{}.Create(request, options)

	var response map[string]interface{}
	json.Unmarshal([]byte(rawResponse), &response)

	return response
}
