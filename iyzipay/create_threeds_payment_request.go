package iyzipay

type CreateThreedsPaymentRequest struct {
	Locale           string `json:"locale,omitempty"`
	ConversationID   string `json:"conversationId,omitempty"`
	PaymentID        string `json:"paymentId,omitempty"`
	ConversationData string `json:"conversationData,omitempty"`
}

func (request CreateThreedsPaymentRequest) toPkiString() string {
	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationID)
	pkiBuilder.append("paymentId", request.PaymentID)
	pkiBuilder.append("conversationData", request.ConversationData)

	return pkiBuilder.getRequestString()
}
