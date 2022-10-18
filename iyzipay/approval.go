package iyzipay

type Approval struct{}

func (approval Approval) Create(request CreateApprovalRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/iyzipos/item/approve", options)
}
