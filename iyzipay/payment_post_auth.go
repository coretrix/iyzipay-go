package iyzipay

type PaymentPostAuth struct{}

func (paymentPostAuth PaymentPostAuth) Create(request CreatePaymentPostAuthRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/postauth", options)
}
