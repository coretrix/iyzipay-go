package iyzipay

type ApiTest struct{}

func (apiTest ApiTest) Retrieve(options Options) string {

	response := connect("GET", "/payment/test", options, "", "")
	return response
}
