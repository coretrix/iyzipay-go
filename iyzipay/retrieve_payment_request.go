package iyzipay

type RetrievePaymentRequest struct {
	Locale                string `json:"locale,omitempty"`
	ConversationId        string `json:"conversationId,omitempty"`
	PaymentId             string `json:"paymentId,omitempty"`
	PaymentConversationId string `json:"paymentConversationId,omitempty"`
}

func (request RetrievePaymentRequest) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationId)
	pkiBuilder.append("paymentId", request.PaymentId)
	pkiBuilder.append("paymentConversationId", request.PaymentConversationId)

	return pkiBuilder.getRequestString()
}
