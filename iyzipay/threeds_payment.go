package iyzipay

type ThreedsPayment struct{}

func (threedsPayment ThreedsPayment) Create(request CreateThreedsPaymentRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/3dsecure/auth", options)
	return response
}

func (threedsPayment ThreedsPayment) Retrieve(request RetrievePaymentRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/detail", options)
	return response
}
