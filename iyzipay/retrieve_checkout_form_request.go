package iyzipay

type RetrieveCheckoutFormRequest struct {
	Locale         string `json:"locale,omitempty"`
	ConversationID string `json:"conversationId,omitempty"`
	Token          string `json:"token,omitempty"`
}

func (request RetrieveCheckoutFormRequest) toPkiString() string {
	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationID)
	pkiBuilder.append("token", request.Token)

	return pkiBuilder.getRequestString()
}
