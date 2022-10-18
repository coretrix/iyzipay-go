package iyzipay

type Payment struct{}

func (payment Payment) Create(request CreatePaymentRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/auth", options)
}

func (payment Payment) Retrieve(request RetrievePaymentRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/detail", options)
}
