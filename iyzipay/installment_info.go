package iyzipay

type InstallmentInfo struct{}

func (installmentInfo InstallmentInfo) Retrieve(request RetrieveInstallmentInfoRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/iyzipos/installment", options)
	return response
}
