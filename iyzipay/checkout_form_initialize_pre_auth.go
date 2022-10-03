package iyzipay

type CheckoutFormInitializePreAuth struct{}

func (checkoutFormInitializePreAuth CheckoutFormInitializePreAuth) Create(request CreateCheckoutFormInitializeRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/iyzipos/checkoutform/initialize/preauth/ecom", options)
	return response
}
