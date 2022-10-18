package iyzipay

type Cancel struct{}

func (cancel Cancel) Create(request CreateCancelRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/cancel", options)
}
