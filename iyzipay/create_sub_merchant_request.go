package iyzipay

type CreateSubMerchantRequest struct {
	Locale                string `json:"locale,omitempty"`
	ConversationId        string `json:"conversationId,omitempty"`
	Name                  string `json:"name,omitempty"`
	Email                 string `json:"email,omitempty"`
	GsmNumber             string `json:"gsmNumber,omitempty"`
	Address               string `json:"address,omitempty"`
	Iban                  string `json:"iban,omitempty"`
	TaxOffice             string `json:"taxOffice,omitempty"`
	ContactName           string `json:"contactName,omitempty"`
	ContactSurname        string `json:"contactSurname,omitempty"`
	LegalCompanyTitle     string `json:"legalCompanyTitle,omitempty"`
	SwiftCode             string `json:"swiftCode,omitempty"`
	Currency              string `json:"currency,omitempty"`
	SubMerchantExternalId string `json:"subMerchantExternalId,omitempty"`
	IdentityNumber        string `json:"identityNumber,omitempty"`
	TaxNumber             string `json:"taxNumber,omitempty"`
	SubMerchantType       string `json:"subMerchantType,omitempty"`
}

func (request CreateSubMerchantRequest) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationId)
	pkiBuilder.append("name", request.Name)
	pkiBuilder.append("email", request.Email)
	pkiBuilder.append("gsmNumber", request.GsmNumber)
	pkiBuilder.append("address", request.Address)
	pkiBuilder.append("iban", request.Iban)
	pkiBuilder.append("taxOffice", request.TaxOffice)
	pkiBuilder.append("contactName", request.ContactName)
	pkiBuilder.append("contactSurname", request.ContactSurname)
	pkiBuilder.append("legalCompanyTitle", request.LegalCompanyTitle)
	pkiBuilder.append("swiftCode", request.SwiftCode)
	pkiBuilder.append("currency", request.Currency)
	pkiBuilder.append("subMerchantExternalId", request.SubMerchantExternalId)
	pkiBuilder.append("identityNumber", request.IdentityNumber)
	pkiBuilder.append("taxNumber", request.TaxNumber)
	pkiBuilder.append("subMerchantType", request.SubMerchantType)

	return pkiBuilder.getRequestString()
}
