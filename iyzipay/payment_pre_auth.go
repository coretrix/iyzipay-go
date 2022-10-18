package iyzipay

type PaymentPreAuth struct{}

func (paymentPreAuth PaymentPreAuth) Create(request CreatePaymentRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/preauth", options)
}

func (paymentPreAuth PaymentPreAuth) Retrieve(request RetrievePaymentRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/detail", options)
}
