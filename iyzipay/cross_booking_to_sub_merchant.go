package iyzipay

type CrossBookingToSubMerchant struct{}

func (crossBookingToSubMerchant CrossBookingToSubMerchant) Create(request CreateCrossBookingRequest, options Options) string {

	response := makeRequest(request, "POST", "/crossbooking/send", options)
	return response
}
