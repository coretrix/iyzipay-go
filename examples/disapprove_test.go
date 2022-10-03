package iyzipay_examples

import (
	"encoding/json"
	"fmt"
	builder "github.com/kurtulussahin/iyzipay-go/examples/builder"
	iyzipay "github.com/kurtulussahin/iyzipay-go/iyzipay"
	"reflect"
)

func ExampleDisapprove() {
	options := iyzipay.Options{}
	options = builder.CreateSampleOptions()

	payment := builder.CreateMarketplacePaymentBuilder()
	basketItem1 := reflect.ValueOf(payment["itemTransactions"]).Index(0).Interface().(map[string]interface{})
	builder.CreateApprovalBuilder(basketItem1["paymentTransactionId"].(string))

	request := iyzipay.CreateDisapprovalRequest{
		Locale:               "tr",
		ConversationId:       "123456789",
		PaymentTransactionId: basketItem1["paymentTransactionId"].(string),
	}

	rawResponse := iyzipay.Disapproval{}.Create(request, options)

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
