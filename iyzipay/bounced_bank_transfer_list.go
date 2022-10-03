package iyzipay

type BouncedBankTransferList struct{}

func (bouncedBankTransferList BouncedBankTransferList) Retrieve(request RetrieveTransactionListRequest, options Options) string {
	response := makeRequest(request, "POST", "/reporting/settlement/bounced", options)
	return response
}
