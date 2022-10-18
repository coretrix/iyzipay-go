package iyzipay

type PayoutCompletedTransactionList struct{}

func (payoutCompletedTransactionList PayoutCompletedTransactionList) Retrieve(request RetrieveTransactionListRequest, options Options) string {
	return makeRequest(request, "POST", "/reporting/settlement/payoutcompleted", options)
}
