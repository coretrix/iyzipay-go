package iyzipay

type RetrieveSubMerchantRequest struct {
	Locale                string `json:"locale,omitempty"`
	ConversationId        string `json:"conversationId,omitempty"`
	SubMerchantExternalId string `json:"subMerchantExternalId,omitempty"`
}

func (request RetrieveSubMerchantRequest) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationId)
	pkiBuilder.append("subMerchantExternalId", request.SubMerchantExternalId)

	return pkiBuilder.getRequestString()
}
