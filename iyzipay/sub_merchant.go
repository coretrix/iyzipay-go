package iyzipay

type SubMerchant struct{}

func (subMerchant SubMerchant) Create(request CreateSubMerchantRequest, options Options) *SubMerchantResponse {
	return decodeResponse(
		makeRequest(request, "POST", "/onboarding/submerchant", options),
		&SubMerchantResponse{},
	).(*SubMerchantResponse)
}

func (subMerchant SubMerchant) Update(request UpdateSubMerchantRequest, options Options) *SubMerchantResponse {
	return decodeResponse(
		makeRequest(request, "PUT", "/onboarding/submerchant", options),
		&SubMerchantResponse{},
	).(*SubMerchantResponse)
}

func (subMerchant SubMerchant) Retrieve(request RetrieveSubMerchantRequest, options Options) *SubMerchantResponse {
	return decodeResponse(
		makeRequest(request, "POST", "/onboarding/submerchant/detail", options),
		&SubMerchantResponse{},
	).(*SubMerchantResponse)
}
