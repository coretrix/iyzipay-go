package iyzipay_examples

import (
	"encoding/json"
	"fmt"
	builder "github.com/kurtulussahin/iyzipay-go/examples/builder"
	iyzipay "github.com/kurtulussahin/iyzipay-go/iyzipay"
)

func ExampleCreateCard() {
	options := iyzipay.Options{}
	options = builder.CreateSampleOptions()

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
