package iyzipay

type CheckoutFormInitialize struct{}

func (checkoutFormInitialize CheckoutFormInitialize) Create(
	request CreateCheckoutFormInitializeRequest,
	options Options,
) *CheckoutFormInitializeResource {
	return decodeResponse(
		makeRequest(request, "POST", "/payment/iyzipos/checkoutform/initialize/auth/ecom", options),
		&CheckoutFormInitializeResource{},
	).(*CheckoutFormInitializeResource)
}
