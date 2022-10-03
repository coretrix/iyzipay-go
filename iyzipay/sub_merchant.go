package iyzipay

type SubMerchant struct{}

func (subMerchant SubMerchant) Create(request CreateSubMerchantRequest, options Options) *SubMerchantResponse {
	response := makeRequest(request, "POST", "/onboarding/submerchant", options)
	return decodeResponse(response, &SubMerchantResponse{}).(*SubMerchantResponse)
}

func (subMerchant SubMerchant) Update(request UpdateSubMerchantRequest, options Options) *SubMerchantResponse {
	response := makeRequest(request, "PUT", "/onboarding/submerchant", options)
	return decodeResponse(response, &SubMerchantResponse{}).(*SubMerchantResponse)
}

func (subMerchant SubMerchant) Retrieve(request RetrieveSubMerchantRequest, options Options) *SubMerchantResponse {
	response := makeRequest(request, "POST", "/onboarding/submerchant/detail", options)
	return decodeResponse(response, &SubMerchantResponse{}).(*SubMerchantResponse)
}
