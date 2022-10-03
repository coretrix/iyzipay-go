package iyzipay

type Payment struct{}

func (payment Payment) Create(request CreatePaymentRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/auth", options)
	return response
}

func (payment Payment) Retrieve(request RetrievePaymentRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/detail", options)
	return response
}
