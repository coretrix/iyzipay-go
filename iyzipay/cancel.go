package iyzipay

type Cancel struct{}

func (cancel Cancel) Create(request CreateCancelRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/cancel", options)
	return response
}
