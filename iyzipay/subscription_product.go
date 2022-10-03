package iyzipay

type SubscriptionProduct struct {
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
}

func (p *SubscriptionProduct) Create(options *Options) (*SubscriptionProductResponse, error) {
	response, err := connectV2("POST", options.baseUrl+"/v2/subscription/products", options.apiKey, options.secretKey, p)
	if err != nil {
		return nil, err
	}

	return decodeResponse(response, &SubscriptionProductResponse{}).(*SubscriptionProductResponse), nil
}

type SubscriptionProductResponse struct {
	Errors
	Status     string                           `json:"status"` // success / failure
	SystemTime int64                            `json:"systemTime"`
	Data       *SubscriptionProductResponseData `json:"data"`
}

type SubscriptionProductResponseData struct {
	ReferenceCode string `json:"referenceCode"`
}
