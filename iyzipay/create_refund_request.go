package iyzipay

type CreateRefundRequest struct {
	Locale               string `json:"locale,omitempty"`
	ConversationID       string `json:"conversationId,omitempty"`
	PaymentTransactionID string `json:"paymentTransactionId,omitempty"`
	Price                string `json:"price,omitempty"`
	IP                   string `json:"ip,omitempty"`
	Currency             string `json:"currency,omitempty"`
}

func (request CreateRefundRequest) toPkiString() string {
	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationID)
	pkiBuilder.append("paymentTransactionId", request.PaymentTransactionID)
	pkiBuilder.appendPrice("price", request.Price)
	pkiBuilder.append("ip", request.IP)
	pkiBuilder.append("currency", request.Currency)

	return pkiBuilder.getRequestString()
}
