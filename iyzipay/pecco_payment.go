package iyzipay

type PeccoPayment struct{}

func (peccoPayment PeccoPayment) Create(request CreatePeccoPaymentRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/pecco/auth", options)
}
