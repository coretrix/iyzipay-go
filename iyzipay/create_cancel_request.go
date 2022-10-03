package iyzipay

type CreateCancelRequest struct {
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
	PaymentId      string `json:"paymentId,omitempty"`
	Ip             string `json:"ip,omitempty"`
}

func (request CreateCancelRequest) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationId)
	pkiBuilder.append("paymentId", request.PaymentId)
	pkiBuilder.append("ip", request.Ip)

	return pkiBuilder.getRequestString()
}
