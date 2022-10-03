package iyzipay

type Buyer struct {
	Id                  string `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	Surname             string `json:"surname,omitempty"`
	IdentityNumber      string `json:"identityNumber,omitempty"`
	Email               string `json:"email,omitempty"`
	GsmNumber           string `json:"gsmNumber,omitempty"`
	RegistrationDate    string `json:"registrationDate,omitempty"`
	LastLoginDate       string `json:"lastLoginDate,omitempty"`
	RegistrationAddress string `json:"registrationAddress,omitempty"`
	City                string `json:"city,omitempty"`
	Country             string `json:"country,omitempty"`
	ZipCode             string `json:"zipCode,omitempty"`
	Ip                  string `json:"ip,omitempty"`
}

func (request Buyer) toPkiString() string {

	pkiBuilder := PkiBuilder{}

	pkiBuilder.append("id", request.Id)
	pkiBuilder.append("name", request.Name)
	pkiBuilder.append("surname", request.Surname)
	pkiBuilder.append("identityNumber", request.IdentityNumber)
	pkiBuilder.append("email", request.Email)
	pkiBuilder.append("gsmNumber", request.GsmNumber)
	pkiBuilder.append("registrationDate", request.RegistrationDate)
	pkiBuilder.append("lastLoginDate", request.LastLoginDate)
	pkiBuilder.append("registrationAddress", request.RegistrationAddress)
	pkiBuilder.append("city", request.City)
	pkiBuilder.append("country", request.Country)
	pkiBuilder.append("zipCode", request.ZipCode)
	pkiBuilder.append("ip", request.Ip)

	return pkiBuilder.getRequestString()
}
