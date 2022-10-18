package iyzipay

type RefundChargedFromMerchant struct{}

func (refund RefundChargedFromMerchant) Create(request CreateRefundRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/iyzipos/refund/merchant/charge", options)
}
