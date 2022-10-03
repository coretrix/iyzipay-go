package iyzipay

type RetrieveBinNumberRequest struct {
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
	BinNumber      string `json:"binNumber,omitempty"`
}

func (request RetrieveBinNumberRequest) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationId)
	pkiBuilder.append("binNumber", request.BinNumber)

	return pkiBuilder.getRequestString()
}
