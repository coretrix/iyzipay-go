package iyzipay

type RetrieveCardListRequest struct {
	Locale         string `json:"locale,omitempty"`
	ConversationId string `json:"conversationId,omitempty"`
	CardUserKey    string `json:"cardUserKey,omitempty"`
}

func (request RetrieveCardListRequest) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationId)
	pkiBuilder.append("cardUserKey", request.CardUserKey)

	return pkiBuilder.getRequestString()
}
