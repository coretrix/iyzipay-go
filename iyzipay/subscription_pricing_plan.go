package iyzipay

const (
	PricingPlanIntervalDaily   = "DAILY"
	PricingPlanIntervalWeekly  = "WEEKLY"
	PricingPlanIntervalMonthly = "MONTHLY"
	PricingPlanIntervalYearly  = "YEARLY"

	PricingPlanTypeRecurring = "RECURRING"
)

type PricingPlan struct {
	Locale               string  `json:"locale,omitempty"`
	ConversationId       string  `json:"conversationId,omitempty"`
	ProductReferenceCode string  `json:"productReferenceCode,omitempty"`
	Name                 string  `json:"name,omitempty"`
	Price                float64 `json:"price,omitempty"`
	CurrencyCode         string  `json:"currencyCode,omitempty"`
	PaymentInterval      string  `json:"paymentInterval,omitempty"`
	PaymentIntervalCount string  `json:"paymentIntervalCount,omitempty"`
	TrialPeriodDays      string  `json:"trialPeriodDays,omitempty"`
	PlanPaymentType      string  `json:"planPaymentType,omitempty"`
	RecurrenceCount      string  `json:"recurrenceCount,omitempty"`
}

func (p *PricingPlan) Create(options *Options) (*PricingPlanResponse, error) {
	response, err := connectV2("POST", options.baseUrl+"/v2/subscription/products/"+p.ProductReferenceCode+"/pricing-plans", options.apiKey, options.secretKey, p)
	if err != nil {
		return nil, err
	}

	return decodeResponse(response, &PricingPlanResponse{}).(*PricingPlanResponse), nil
}

type PricingPlanResponse struct {
	Errors
	Status     string                   `json:"status"` // success / failure
	SystemTime int64                    `json:"systemTime"`
	Data       *PricingPlanResponseData `json:"data"`
}

type PricingPlanResponseData struct {
	ReferenceCode string `json:"referenceCode"`
	Status        string `json:"status"` // RECURRING
}
