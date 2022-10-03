package iyzipay

type PeccoPayment struct{}

func (peccoPayment PeccoPayment) Create(request CreatePeccoPaymentRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/pecco/auth", options)
	return response
}
