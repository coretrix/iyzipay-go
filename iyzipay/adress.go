package iyzipay

type Address struct {
	ContactName string `json:"contactName,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	Address     string `json:"address,omitempty"`
	ZipCode     string `json:"zipCode,omitempty"`
}

func (request Address) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("address", request.Address)
	pkiBuilder.append("zipCode", request.ZipCode)
	pkiBuilder.append("contactName", request.ContactName)
	pkiBuilder.append("city", request.City)
	pkiBuilder.append("country", request.Country)

	return pkiBuilder.getRequestString()
}
