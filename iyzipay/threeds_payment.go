package iyzipay

type ThreedsPayment struct{}

func (threedsPayment ThreedsPayment) Create(request CreateThreedsPaymentRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/3dsecure/auth", options)
}

func (threedsPayment ThreedsPayment) Retrieve(request RetrievePaymentRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/detail", options)
}
