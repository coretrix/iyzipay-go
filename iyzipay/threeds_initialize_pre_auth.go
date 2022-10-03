package iyzipay

type ThreedsInitializePreAuth struct{}

func (threedsInitializePreAuth ThreedsInitializePreAuth) Create(request CreatePaymentRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/3dsecure/initialize/preauth", options)
	return response
}
