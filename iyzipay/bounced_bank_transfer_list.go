package iyzipay

type BouncedBankTransferList struct{}

func (bouncedBankTransferList BouncedBankTransferList) Retrieve(request RetrieveTransactionListRequest, options Options) string {
	return makeRequest(request, "POST", "/reporting/settlement/bounced", options)
}
