package iyzipay

type ThreedsInitialize struct{}

func (threedsInitialize ThreedsInitialize) Create(request CreatePaymentRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/3dsecure/initialize", options)
}
