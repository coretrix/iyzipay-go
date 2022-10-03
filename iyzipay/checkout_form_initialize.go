package iyzipay

type CheckoutFormInitialize struct{}

func (checkoutFormInitialize CheckoutFormInitialize) Create(request CreateCheckoutFormInitializeRequest, options Options) *CheckoutFormInitializeResource {
	response := makeRequest(request, "POST", "/payment/iyzipos/checkoutform/initialize/auth/ecom", options)
	return decodeResponse(response, &CheckoutFormInitializeResource{}).(*CheckoutFormInitializeResource)
}
