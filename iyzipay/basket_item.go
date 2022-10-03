package iyzipay

type BasketItem struct {
	Id               string `json:"id,omitempty"`
	Price            string `json:"price,omitempty"`
	Name             string `json:"name,omitempty"`
	Category1        string `json:"category1,omitempty"`
	Category2        string `json:"category2,omitempty"`
	ItemType         string `json:"itemType,omitempty"`
	SubMerchantKey   string `json:"subMerchantKey,omitempty"`
	SubMerchantPrice string `json:"subMerchantPrice,omitempty"`
}

func (request BasketItem) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("id", request.Id)
	pkiBuilder.appendPrice("price", request.Price)
	pkiBuilder.append("name", request.Name)
	pkiBuilder.append("category1", request.Category1)
	pkiBuilder.append("category2", request.Category2)
	pkiBuilder.append("itemType", request.ItemType)
	pkiBuilder.append("subMerchantKey", request.SubMerchantKey)
	pkiBuilder.appendPrice("subMerchantPrice", request.SubMerchantPrice)

	return pkiBuilder.getRequestString()
}
