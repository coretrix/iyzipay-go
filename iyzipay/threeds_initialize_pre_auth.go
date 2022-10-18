package iyzipay

type ThreedsInitializePreAuth struct{}

func (threedsInitializePreAuth ThreedsInitializePreAuth) Create(request CreatePaymentRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/3dsecure/initialize/preauth", options)
}
