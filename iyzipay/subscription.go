package iyzipay

const (
	SubscriptionInitialStatusActive  = "ACTIVE"
	SubscriptionInitialStatusPending = "PENDING"
)

type Subscription struct {
	Locale                    string                `json:"locale,omitempty"`
	ConversationId            string                `json:"conversationId,omitempty"`
	CallbackUrl               string                `json:"callbackUrl,omitempty"`
	PricingPlanReferenceCode  string                `json:"pricingPlanReferenceCode,omitempty"`
	SubscriptionInitialStatus string                `json:"subscriptionInitialStatus,omitempty"` // PENDING or ACTIVE
	SubscriptionCustomer      *SubscriptionCustomer `json:"customer,omitempty"`
}

func (p *Subscription) Create(options *Options) (*SubscriptionResponse, error) {
	response, err := connectV2("POST", options.baseUrl+"/v2/subscription/checkoutform/initialize", options.apiKey, options.secretKey, p)
	if err != nil {
		return nil, err
	}

	return decodeResponse(response, &SubscriptionResponse{}).(*SubscriptionResponse), nil
}

type SubscriptionResponse struct {
	Errors
	Status              string `json:"status"` // success / failure
	SystemTime          int64  `json:"systemTime"`
	CheckoutFormContent string `json:"checkoutFormContent"`
	Token               string `json:"token"`
	TokenExpireTime     int64  `json:"tokenExpireTime"`
}

type SubscriptionCheckoutFormResult struct {
	Token string `json:"token,omitempty"`
}

func (p *SubscriptionCheckoutFormResult) Get(options *Options) (*SubscriptionCheckoutFormResponse, error) {
	response, err := connectV2("GET", options.baseUrl+"/v2/subscription/checkoutform/"+p.Token, options.apiKey, options.secretKey, nil)
	if err != nil {
		return nil, err
	}

	return decodeResponse(response, &SubscriptionCheckoutFormResponse{}).(*SubscriptionCheckoutFormResponse), nil
}

type SubscriptionCheckoutFormResponse struct {
	Errors
	Status     string                                `json:"status"` // success / failure
	SystemTime int64                                 `json:"systemTime"`
	Data       *SubscriptionCheckoutFormResponseData `json:"data"`
}

type SubscriptionCheckoutFormResponseData struct {
	ReferenceCode         string `json:"referenceCode"`
	CustomerReferenceCode string `json:"customerReferenceCode"`
	SubscriptionStatus    string `json:"subscriptionStatus"` // ACTIVE or PENDING
	TrialDays             int16  `json:"trialDays"`
	TrialStartDate        int64  `json:"trialStartDate"`
	TrialEndDate          int64  `json:"trialEndDate"`
	CreatedDate           int64  `json:"createdDate"`
	StartDate             int64  `json:"startDate"`
}

type SubscriptionUpgrade struct {
	SubscriptionReferenceCode   string `json:"subscriptionReferenceCode,omitempty"`
	NewPricingPlanReferenceCode string `json:"newPricingPlanReferenceCode,omitempty"`
	UpgradePeriod               string `json:"upgradePeriod,omitempty"`
	UseTrial                    bool   `json:"useTrial,omitempty"`
	ResetRecurrenceCount        bool   `json:"resetRecurrenceCount,omitempty"`
}

func (p *SubscriptionUpgrade) Upgrade(options *Options) (*SubscriptionUpgradeResponse, error) {
	response, err := connectV2("POST", options.baseUrl+"/v2/subscription/subscriptions/"+p.SubscriptionReferenceCode+"/upgrade", options.apiKey, options.secretKey, p)
	if err != nil {
		return nil, err
	}

	return decodeResponse(response, &SubscriptionUpgradeResponse{}).(*SubscriptionUpgradeResponse), nil
}

type SubscriptionUpgradeResponse struct {
	Errors
	Status     string                                `json:"status"` // success / failure
	SystemTime int64                                 `json:"systemTime"`
	Data       *SubscriptionCheckoutFormResponseData `json:"data"`
}

type SubscriptionCancel struct {
	Locale                    string `json:"locale,omitempty"`
	ConversationId            string `json:"conversationId,omitempty"`
	SubscriptionReferenceCode string `json:"subscriptionReferenceCode,omitempty"`
}

func (p *SubscriptionCancel) Cancel(options *Options) (*SubscriptionCancelResponse, error) {
	response, err := connectV2("POST", options.baseUrl+"/v2/subscription/subscriptions/"+p.SubscriptionReferenceCode+"/cancel", options.apiKey, options.secretKey, p)
	if err != nil {
		return nil, err
	}

	return decodeResponse(response, &SubscriptionCancelResponse{}).(*SubscriptionCancelResponse), nil
}

type SubscriptionCancelResponse struct {
	Errors
	Status     string `json:"status"` // success / failure
	SystemTime int64  `json:"systemTime"`
}

type SubscriptionUpdateCardInformation struct {
	Locale                    string `json:"locale,omitempty"`
	ConversationId            string `json:"conversationId,omitempty"`
	SubscriptionReferenceCode string `json:"subscriptionReferenceCode,omitempty"`
	CustomerReferenceCode     string `json:"customerReferenceCode,omitempty"`
	CallBackUrl               string `json:"callbackUrl,omitempty"`
}

func (p *SubscriptionUpdateCardInformation) UpdateSubscriptionCardInformation(options *Options) (*SubscriptionUpdateCardInformationResponse, error) {
	response, err := connectV2("POST", options.baseUrl+"/v2/subscription/card-update/checkoutform/initialize/with-subscription", options.apiKey, options.secretKey, p)
	if err != nil {
		return nil, err
	}

	return decodeResponse(response, &SubscriptionUpdateCardInformationResponse{}).(*SubscriptionUpdateCardInformationResponse), nil
}

func (p *SubscriptionUpdateCardInformation) UpdateCustomerCardInformation(options *Options) (*SubscriptionUpdateCardInformationResponse, error) {
	response, err := connectV2("POST", options.baseUrl+"/v2/subscription/card-update/checkoutform/initialize", options.apiKey, options.secretKey, p)
	if err != nil {
		return nil, err
	}

	return decodeResponse(response, &SubscriptionUpdateCardInformationResponse{}).(*SubscriptionUpdateCardInformationResponse), nil
}

type SubscriptionUpdateCardInformationResponse struct {
	Errors
	Status              string `json:"status"` // success / failure
	SystemTime          int64  `json:"systemTime"`
	TokenExpireTime     int64  `json:"tokenExpireTime"`
	Token               string `json:"token"`
	CheckoutFormContent string `json:"checkoutFormContent"`
}

type Errors struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	ErrorGroup   string `json:"errorGroup"`
}
