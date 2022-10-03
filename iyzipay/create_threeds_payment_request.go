package iyzipay

type CreateThreedsPaymentRequest struct {
	Locale           string `json:"locale,omitempty"`
	ConversationId   string `json:"conversationId,omitempty"`
	PaymentId        string `json:"paymentId,omitempty"`
	ConversationData string `json:"conversationData,omitempty"`
}

func (request CreateThreedsPaymentRequest) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationId)
	pkiBuilder.append("paymentId", request.PaymentId)
	pkiBuilder.append("conversationData", request.ConversationData)

	return pkiBuilder.getRequestString()
}
