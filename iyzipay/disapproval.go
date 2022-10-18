package iyzipay

type Disapproval struct{}

func (approval Disapproval) Create(request CreateDisapprovalRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/iyzipos/item/disapprove", options)
}
