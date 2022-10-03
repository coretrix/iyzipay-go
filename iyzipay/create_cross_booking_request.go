package iyzipay

type CreateCrossBookingRequest struct {
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
	Submerchantkey string `json:"submerchantkey,omitempty"`
	Price          string `json:"price,omitempty"`
	Reason         string `json:"reason,omitempty"`
	Currency       string `json:"currency,omitempty"`
}

func (request CreateCrossBookingRequest) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationId)
	pkiBuilder.append("submerchantkey", request.Submerchantkey)
	pkiBuilder.appendPrice("price", request.Price)
	pkiBuilder.append("reason", request.Reason)
	pkiBuilder.append("currency", request.Currency)

	return pkiBuilder.getRequestString()
}
