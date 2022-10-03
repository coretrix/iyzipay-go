package iyzipay

type PaymentCard struct {
	CardHolderName string `json:"cardHolderName,omitempty"`
	CardNumber     string `json:"cardNumber,omitempty"`
	ExpireYear     string `json:"expireYear,omitempty"`
	ExpireMonth    string `json:"expireMonth,omitempty"`
	Cvc            string `json:"cvc,omitempty"`
	RegisterCard   string `json:"registerCard,omitempty"`
	CardAlias      string `json:"cardAlias,omitempty"`
	CardToken      string `json:"cardToken,omitempty"`
	CardUserKey    string `json:"cardUserKey,omitempty"`
}

func (request PaymentCard) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("cardHolderName", request.CardHolderName)
	pkiBuilder.append("cardNumber", request.CardNumber)
	pkiBuilder.append("expireYear", request.ExpireYear)
	pkiBuilder.append("expireMonth", request.ExpireMonth)
	pkiBuilder.append("cvc", request.Cvc)
	pkiBuilder.append("registerCard", request.RegisterCard)
	pkiBuilder.append("cardAlias", request.CardAlias)
	pkiBuilder.append("cardToken", request.CardToken)
	pkiBuilder.append("cardUserKey", request.CardUserKey)

	return pkiBuilder.getRequestString()
}
