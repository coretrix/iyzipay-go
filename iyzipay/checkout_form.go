package iyzipay

type CheckoutForm struct{}

func (checkoutForm CheckoutForm) Retrieve(request RetrieveCheckoutFormRequest, options Options) *CheckoutFormResponse {
	return decodeResponse(
		makeRequest(request, "POST", "/payment/iyzipos/checkoutform/auth/ecom/detail", options),
		&CheckoutFormResponse{},
	).(*CheckoutFormResponse)
}
