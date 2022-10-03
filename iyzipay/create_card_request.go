package iyzipay

type CreateCardRequest struct {
	Locale         string          `json:"locale,omitempty"`
	ConversationId string          `json:"conversationId,omitempty"`
	ExternalId     string          `json:"externalId,omitempty"`
	Email          string          `json:"email,omitempty"`
	CardUserKey    string          `json:"cardUserKey,omitempty"`
	Card           CardInformation `json:"card,omitempty"`
}

func (request CreateCardRequest) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationId)
	pkiBuilder.append("externalId", request.ExternalId)
	pkiBuilder.append("email", request.Email)
	pkiBuilder.append("cardUserKey", request.CardUserKey)
	pkiBuilder.append("card", request.Card.toPkiString())

	return pkiBuilder.getRequestString()
}
