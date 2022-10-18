package iyzipay

type CreateCheckoutFormInitializeRequest struct {
	Locale              string       `json:"locale,omitempty"`
	ConversationID      string       `json:"conversationId,omitempty"`
	Price               string       `json:"price,omitempty"`
	BasketID            string       `json:"basketId,omitempty"`
	PaymentGroup        string       `json:"paymentGroup,omitempty"`
	Buyer               Buyer        `json:"buyer,omitempty"`
	ShippingAddress     Address      `json:"billingAddress,omitempty"`
	BillingAddress      Address      `json:"shippingAddress,omitempty"`
	BasketItems         []BasketItem `json:"basketItems,omitempty"`
	CallbackURL         string       `json:"callbackUrl,omitempty"`
	PaymentSource       string       `json:"paymentSource,omitempty"`
	Currency            string       `json:"currency,omitempty"`
	PosOrderID          string       `json:"posOrderId,omitempty"`
	PaidPrice           string       `json:"paidPrice,omitempty"`
	ForceThreeDS        string       `json:"forceThreeDS,omitempty"`
	CardUserKey         string       `json:"cardUserKey,omitempty"`
	EnabledInstallments []string     `json:"enabledInstallments,omitempty"`
}

func (request CreateCheckoutFormInitializeRequest) toPkiString() string {
	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationID)
	pkiBuilder.appendPrice("price", request.Price)
	pkiBuilder.append("basketId", request.BasketID)
	pkiBuilder.append("paymentGroup", request.PaymentGroup)
	pkiBuilder.append("buyer", request.Buyer.toPkiString())
	pkiBuilder.append("shippingAddress", request.ShippingAddress.toPkiString())
	pkiBuilder.append("billingAddress", request.BillingAddress.toPkiString())

	basketItemsPki := []string{}
	for i := range request.BasketItems {
		basketItemsPki = append(basketItemsPki, request.BasketItems[i].toPkiString())
	}

	pkiBuilder.appendArray("basketItems", basketItemsPki)

	pkiBuilder.append("callbackUrl", request.CallbackURL)
	pkiBuilder.append("paymentSource", request.PaymentSource)
	pkiBuilder.append("currency", request.Currency)
	pkiBuilder.append("posOrderId", request.PosOrderID)
	pkiBuilder.appendPrice("paidPrice", request.PaidPrice)
	pkiBuilder.append("forceThreeDS", request.ForceThreeDS)
	pkiBuilder.append("cardUserKey", request.CardUserKey)
	pkiBuilder.appendArray("enabledInstallments", request.EnabledInstallments)

	return pkiBuilder.getRequestString()
}
