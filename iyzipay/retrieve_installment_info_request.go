package iyzipay

type RetrieveInstallmentInfoRequest struct {
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
	BinNumber      string `json:"binNumber,omitempty"`
	Price          string `json:"price,omitempty"`
}

func (request RetrieveInstallmentInfoRequest) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationId)
	pkiBuilder.append("binNumber", request.BinNumber)
	pkiBuilder.appendPrice("price", request.Price)

	return pkiBuilder.getRequestString()
}
