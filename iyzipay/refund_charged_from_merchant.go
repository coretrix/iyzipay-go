package iyzipay

type RefundChargedFromMerchant struct{}

func (refund RefundChargedFromMerchant) Create(request CreateRefundRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/iyzipos/refund/merchant/charge", options)
	return response
}
