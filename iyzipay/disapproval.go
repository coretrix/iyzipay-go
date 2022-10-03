package iyzipay

type Disapproval struct{}

func (approval Disapproval) Create(request CreateDisapprovalRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/iyzipos/item/disapprove", options)
	return response
}
