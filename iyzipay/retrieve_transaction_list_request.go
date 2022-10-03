package iyzipay

type RetrieveTransactionListRequest struct {
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
	Date           string `json:"date,omitempty"`
}

func (request RetrieveTransactionListRequest) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationId)
	pkiBuilder.append("date", request.Date)

	return pkiBuilder.getRequestString()
}
