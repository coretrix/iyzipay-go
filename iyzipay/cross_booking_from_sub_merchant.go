package iyzipay

type CrossBookingFromSubMerchant struct{}

func (crossBookingFromSubMerchant CrossBookingFromSubMerchant) Create(request CreateCrossBookingRequest, options Options) string {

	response := makeRequest(request, "POST", "/crossbooking/receive", options)
	return response
}
