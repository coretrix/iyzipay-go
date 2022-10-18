package iyzipay

type CrossBookingFromSubMerchant struct{}

func (crossBookingFromSubMerchant CrossBookingFromSubMerchant) Create(request CreateCrossBookingRequest, options Options) string {
	return makeRequest(request, "POST", "/crossbooking/receive", options)
}
