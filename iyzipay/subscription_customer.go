package iyzipay

type SubscriptionCustomer struct {
	Locale          string               `json:"locale,omitempty"`
	ConversationId  string               `json:"conversationId,omitempty"`
	Name            string               `json:"name,omitempty"`
	Surname         string               `json:"surname,omitempty"`
	IdentityNumber  string               `json:"identityNumber,omitempty"`
	Email           string               `json:"email,omitempty"`
	GSMNumber       string               `json:"gsmNumber,omitempty"`
	BillingAddress  *SubscriptionAddress `json:"billingAddress,omitempty"`
	ShippingAddress *SubscriptionAddress `json:"shippingAddress,omitempty"`
}

type SubscriptionAddress struct {
	ContactName string `json:"contactName,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	Address     string `json:"address,omitempty"`
	ZipCode     string `json:"zipCode,omitempty"`
}

func (p *SubscriptionCustomer) Create(options *Options) (*SubscriptionCustomerResponse, error) {
	response, err := connectV2("POST", options.baseUrl+"/v2/subscription/customers", options.apiKey, options.secretKey, p)
	if err != nil {
		return nil, err
	}

	return decodeResponse(response, &SubscriptionCustomerResponse{}).(*SubscriptionCustomerResponse), nil
}

type SubscriptionCustomerResponse struct {
	Status     string                            `json:"status"` // success / failure
	SystemTime int64                             `json:"systemTime"`
	Data       *SubscriptionCustomerResponseData `json:"data"`
}

type SubscriptionCustomerResponseData struct {
	ReferenceCode string `json:"referenceCode"`
	Status        string `json:"status"` // ACTIVE
}
