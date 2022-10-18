package iyzipay

type RetrieveSubMerchantRequest struct {
	Locale                string `json:"locale,omitempty"`
	ConversationID        string `json:"conversationId,omitempty"`
	SubMerchantExternalID string `json:"subMerchantExternalId,omitempty"`
}

func (request RetrieveSubMerchantRequest) toPkiString() string {
	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationID)
	pkiBuilder.append("subMerchantExternalId", request.SubMerchantExternalID)

	return pkiBuilder.getRequestString()
}
