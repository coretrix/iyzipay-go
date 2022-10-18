package iyzipay

type CreateCancelRequest struct {
	Locale         string `json:"locale,omitempty"`
	ConversationID string `json:"conversationId,omitempty"`
	PaymentID      string `json:"paymentId,omitempty"`
	IP             string `json:"ip,omitempty"`
}

func (request CreateCancelRequest) toPkiString() string {
	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationID)
	pkiBuilder.append("paymentId", request.PaymentID)
	pkiBuilder.append("ip", request.IP)

	return pkiBuilder.getRequestString()
}
