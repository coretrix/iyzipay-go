package iyzipay

type CreateCardRequest struct {
	Locale         string          `json:"locale,omitempty"`
	ConversationID string          `json:"conversationId,omitempty"`
	ExternalID     string          `json:"externalId,omitempty"`
	Email          string          `json:"email,omitempty"`
	CardUserKey    string          `json:"cardUserKey,omitempty"`
	Card           CardInformation `json:"card,omitempty"`
}

func (request CreateCardRequest) toPkiString() string {
	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationID)
	pkiBuilder.append("externalId", request.ExternalID)
	pkiBuilder.append("email", request.Email)
	pkiBuilder.append("cardUserKey", request.CardUserKey)
	pkiBuilder.append("card", request.Card.toPkiString())

	return pkiBuilder.getRequestString()
}
