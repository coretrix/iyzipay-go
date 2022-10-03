package iyzipay

type CreatePaymentPostAuthRequest struct {
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
	PaymentId      string `json:"paymentId,omitempty"`
	Ip             string `json:"ip,omitempty"`
	PaidPrice      string `json:"paidPrice,omitempty"`
	Currency       string `json:"currency,omitempty"`
}

func (request CreatePaymentPostAuthRequest) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationId)
	pkiBuilder.append("paymentId", request.PaymentId)
	pkiBuilder.append("ip", request.Ip)
	pkiBuilder.appendPrice("paidPrice", request.PaidPrice)
	pkiBuilder.append("currency", request.Currency)

	return pkiBuilder.getRequestString()
}
