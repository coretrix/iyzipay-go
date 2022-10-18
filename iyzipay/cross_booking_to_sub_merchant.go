package iyzipay

type CrossBookingToSubMerchant struct{}

func (crossBookingToSubMerchant CrossBookingToSubMerchant) Create(request CreateCrossBookingRequest, options Options) string {
	return makeRequest(request, "POST", "/crossbooking/send", options)
}
