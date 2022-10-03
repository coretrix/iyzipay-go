package iyzipay

type CreatePeccoInitializeRequest struct {
	Locale          string       `json:"locale,omitempty"`
	ConversationId  string       `json:"conversationId,omitempty"`
	Price           string       `json:"price,omitempty"`
	BasketId        string       `json:"basketId,omitempty"`
	PaymentGroup    string       `json:"paymentGroup,omitempty"`
	Buyer           Buyer        `json:"buyer,omitempty"`
	ShippingAddress Address      `json:"billingAddress,omitempty"`
	BillingAddress  Address      `json:"shippingAddress,omitempty"`
	BasketItems     []BasketItem `json:"basketItems,omitempty"`
	CallbackUrl     string       `json:"callbackUrl,omitempty"`
	PaymentSource   string       `json:"paymentSource,omitempty"`
	PaidPrice       string       `json:"paidPrice,omitempty"`
}

func (request CreatePeccoInitializeRequest) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationId)
	pkiBuilder.appendPrice("price", request.Price)
	pkiBuilder.append("basketId", request.BasketId)
	pkiBuilder.append("paymentGroup", request.PaymentGroup)
	pkiBuilder.append("buyer", request.Buyer.toPkiString())
	pkiBuilder.append("shippingAddress", request.ShippingAddress.toPkiString())
	pkiBuilder.append("billingAddress", request.BillingAddress.toPkiString())

	basketItemsPki := []string{}
	for i := range request.BasketItems {
		basketItemsPki = append(basketItemsPki, request.BasketItems[i].toPkiString())
	}
	pkiBuilder.appendArray("basketItems", basketItemsPki)

	pkiBuilder.append("callbackUrl", request.CallbackUrl)
	pkiBuilder.append("paymentSource", request.PaymentSource)
	pkiBuilder.appendPrice("paidPrice", request.PaidPrice)

	return pkiBuilder.getRequestString()
}
