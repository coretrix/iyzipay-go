package iyzipay

type CreatePaymentPostAuthRequest struct {
	Locale         string `json:"locale,omitempty"`
	ConversationID string `json:"conversationId,omitempty"`
	PaymentID      string `json:"paymentId,omitempty"`
	IP             string `json:"ip,omitempty"`
	PaidPrice      string `json:"paidPrice,omitempty"`
	Currency       string `json:"currency,omitempty"`
}

func (request CreatePaymentPostAuthRequest) toPkiString() string {
	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationID)
	pkiBuilder.append("paymentId", request.PaymentID)
	pkiBuilder.append("ip", request.IP)
	pkiBuilder.appendPrice("paidPrice", request.PaidPrice)
	pkiBuilder.append("currency", request.Currency)

	return pkiBuilder.getRequestString()
}
