package iyzipay

type DeleteCardRequest struct {
	Locale         string `json:"locale,omitempty"`
	ConversationID string `json:"conversationId,omitempty"`
	CardUserKey    string `json:"cardUserKey,omitempty"`
	CardToken      string `json:"cardToken,omitempty"`
}

func (request DeleteCardRequest) toPkiString() string {
	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationID)
	pkiBuilder.append("cardUserKey", request.CardUserKey)
	pkiBuilder.append("cardToken", request.CardToken)

	return pkiBuilder.getRequestString()
}
