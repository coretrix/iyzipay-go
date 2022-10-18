package iyzipay

type InstallmentInfo struct{}

func (installmentInfo InstallmentInfo) Retrieve(request RetrieveInstallmentInfoRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/iyzipos/installment", options)
}
