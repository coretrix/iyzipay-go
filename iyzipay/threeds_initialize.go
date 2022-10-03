package iyzipay

type ThreedsInitialize struct{}

func (threedsInitialize ThreedsInitialize) Create(request CreatePaymentRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/3dsecure/initialize", options)
	return response
}
