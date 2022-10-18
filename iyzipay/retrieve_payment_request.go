package iyzipay

type RetrievePaymentRequest struct {
	Locale                string `json:"locale,omitempty"`
	ConversationID        string `json:"conversationId,omitempty"`
	PaymentID             string `json:"paymentId,omitempty"`
	PaymentConversationID string `json:"paymentConversationId,omitempty"`
}

func (request RetrievePaymentRequest) toPkiString() string {
	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationID)
	pkiBuilder.append("paymentId", request.PaymentID)
	pkiBuilder.append("paymentConversationId", request.PaymentConversationID)

	return pkiBuilder.getRequestString()
}
