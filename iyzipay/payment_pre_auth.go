package iyzipay

type PaymentPreAuth struct{}

func (paymentPreAuth PaymentPreAuth) Create(request CreatePaymentRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/preauth", options)
	return response
}

func (paymentPreAuth PaymentPreAuth) Retrieve(request RetrievePaymentRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/detail", options)
	return response
}
