package builder

import (
	"encoding/json"
	iyzipay "github.com/kurtulussahin/iyzipay-go/iyzipay"
)

func CreateApprovalBuilder(paymentTransactionId string) map[string]interface{} {
	options := iyzipay.Options{}
	options = CreateSampleOptions()

	request := iyzipay.CreateApprovalRequest{
		Locale:               "tr",
		ConversationId:       "123456789",
		PaymentTransactionId: paymentTransactionId,
	}

	rawResponse := iyzipay.Approval{}.Create(request, options)

	var response map[string]interface{}
	json.Unmarshal([]byte(rawResponse), &response)

	return response
}
