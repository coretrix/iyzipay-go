package iyzipay

type PaymentPostAuth struct{}

func (paymentPostAuth PaymentPostAuth) Create(request CreatePaymentPostAuthRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/postauth", options)
	return response
}
