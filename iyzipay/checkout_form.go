package iyzipay

type CheckoutForm struct{}

func (checkoutForm CheckoutForm) Retrieve(request RetrieveCheckoutFormRequest, options Options) *CheckoutFormResponse {
	response := makeRequest(request, "POST", "/payment/iyzipos/checkoutform/auth/ecom/detail", options)
	return decodeResponse(response, &CheckoutFormResponse{}).(*CheckoutFormResponse)
}
