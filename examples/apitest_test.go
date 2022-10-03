package iyzipay_examples

import (
	"encoding/json"
	"fmt"
	builder "github.com/kurtulussahin/iyzipay-go/examples/builder"
	iyzipay "github.com/kurtulussahin/iyzipay-go/iyzipay"
)

func ExampleApiTest() {
	options := iyzipay.Options{}
	options = builder.CreateSampleOptions()

	rawResponse := iyzipay.ApiTest{}.Retrieve(options)

	var response map[string]interface{}
	json.Unmarshal([]byte(rawResponse), &response)

	fmt.Println(response["status"])
	fmt.Println(response["locale"])
	//Output:
	//success
	//tr
}
