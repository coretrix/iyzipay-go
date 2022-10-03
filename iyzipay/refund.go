package iyzipay

type Refund struct{}

func (refund Refund) Create(request CreateRefundRequest, options Options) *RefundResponse {
	response := makeRequest(request, "POST", "/payment/refund", options)
	return decodeResponse(response, &RefundResponse{}).(*RefundResponse)
}
