package iyzipay

type CardInformation struct {
	CardAlias      string `json:"cardAlias,omitempty"`
	CardNumber     string `json:"cardNumber,omitempty"`
	ExpireYear     string `json:"expireYear,omitempty"`
	ExpireMonth    string `json:"expireMonth,omitempty"`
	CardHolderName string `json:"cardHolderName,omitempty"`
}

func (request CardInformation) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("cardAlias", request.CardAlias)
	pkiBuilder.append("cardNumber", request.CardNumber)
	pkiBuilder.append("expireYear", request.ExpireYear)
	pkiBuilder.append("expireMonth", request.ExpireMonth)
	pkiBuilder.append("cardHolderName", request.CardHolderName)

	return pkiBuilder.getRequestString()
}
