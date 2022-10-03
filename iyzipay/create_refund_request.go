package iyzipay

type CreateRefundRequest struct {
	Locale               string `json:"locale,omitempty"`
	ConversationId       string `json:"conversationId,omitempty"`
	PaymentTransactionId string `json:"paymentTransactionId,omitempty"`
	Price                string `json:"price,omitempty"`
	Ip                   string `json:"ip,omitempty"`
	Currency             string `json:"currency,omitempty"`
}

func (request CreateRefundRequest) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationId)
	pkiBuilder.append("paymentTransactionId", request.PaymentTransactionId)
	pkiBuilder.appendPrice("price", request.Price)
	pkiBuilder.append("ip", request.Ip)
	pkiBuilder.append("currency", request.Currency)

	return pkiBuilder.getRequestString()
}
