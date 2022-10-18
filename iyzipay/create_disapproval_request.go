package iyzipay

type CreateDisapprovalRequest struct {
	Locale               string `json:"locale,omitempty"`
	ConversationID       string `json:"conversationId,omitempty"`
	PaymentTransactionID string `json:"paymentTransactionId,omitempty"`
}

func (request CreateDisapprovalRequest) toPkiString() string {
	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationID)
	pkiBuilder.append("paymentTransactionId", request.PaymentTransactionID)

	return pkiBuilder.getRequestString()
}
