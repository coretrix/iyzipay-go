package iyzipay

type Bkm struct{}

func (bkm Bkm) Retrieve(request RetrieveBkmRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/iyzipos/checkoutform/auth/ecom/detail", options)
	return response
}
