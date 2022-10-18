package iyzipay

type Refund struct{}

func (refund Refund) Create(request CreateRefundRequest, options Options) *RefundResponse {
	return decodeResponse(makeRequest(request, "POST", "/payment/refund", options), &RefundResponse{}).(*RefundResponse)
}
