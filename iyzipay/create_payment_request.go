package iyzipay

type CreatePaymentRequest struct {
	Locale          string       `json:"locale,omitempty"`
	ConversationID  string       `json:"conversationId,omitempty"`
	Price           string       `json:"price,omitempty"`
	PaidPrice       string       `json:"paidPrice,omitempty"`
	BasketID        string       `json:"basketId,omitempty"`
	PaymentGroup    string       `json:"paymentGroup,omitempty"`
	PaymentCard     PaymentCard  `json:"paymentCard,omitempty"`
	Buyer           Buyer        `json:"buyer,omitempty"`
	ShippingAddress Address      `json:"billingAddress,omitempty"`
	BillingAddress  Address      `json:"shippingAddress,omitempty"`
	BasketItems     []BasketItem `json:"basketItems,omitempty"`
	PaymentSource   string       `json:"paymentSource,omitempty"`
	Currency        string       `json:"currency,omitempty"`
	PosOrderID      string       `json:"posOrderId,omitempty"`
	ConnectorName   string       `json:"connectorName,omitempty"`
	CallbackURL     string       `json:"callbackUrl,omitempty"`
}

func (request CreatePaymentRequest) toPkiString() string {
	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("locale", request.Locale)
	pkiBuilder.append("conversationId", request.ConversationID)
	pkiBuilder.appendPrice("price", request.Price)
	pkiBuilder.appendPrice("paidPrice", request.PaidPrice)
	pkiBuilder.append("basketId", request.BasketID)
	pkiBuilder.append("paymentGroup", request.PaymentGroup)
	pkiBuilder.append("paymentCard", request.PaymentCard.toPkiString())
	pkiBuilder.append("buyer", request.Buyer.toPkiString())
	pkiBuilder.append("shippingAddress", request.ShippingAddress.toPkiString())
	pkiBuilder.append("billingAddress", request.BillingAddress.toPkiString())

	basketItemsPki := []string{}
	for i := range request.BasketItems {
		basketItemsPki = append(basketItemsPki, request.BasketItems[i].toPkiString())
	}

	pkiBuilder.appendArray("basketItems", basketItemsPki)

	pkiBuilder.append("paymentSource", request.PaymentSource)
	pkiBuilder.append("currency", request.Currency)
	pkiBuilder.append("posOrderId", request.PosOrderID)
	pkiBuilder.append("connectorName", request.ConnectorName)
	pkiBuilder.append("callbackUrl", request.CallbackURL)

	return pkiBuilder.getRequestString()
}
