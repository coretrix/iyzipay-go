package builder

import (
	"encoding/json"
	iyzipay "github.com/kurtulussahin/iyzipay-go/iyzipay"
)

func CreateCardBuilder() map[string]interface{} {
	options := iyzipay.Options{}
	options = CreateSampleOptions()

	cardInformation := iyzipay.CardInformation{
		CardAlias:      "card alias",
		CardHolderName: "John Doe",
		CardNumber:     "5528790000000008",
		ExpireYear:     "2030",
		ExpireMonth:    "12",
	}

	request := iyzipay.CreateCardRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		Email:          "email@email.com",
		ExternalId:     "external id",
		Card:           cardInformation,
	}

	rawResponse := iyzipay.Card{}.Create(request, options)

	var response map[string]interface{}
	json.Unmarshal([]byte(rawResponse), &response)

	return response
}
