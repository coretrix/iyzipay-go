package iyzipay

type PeccoInitialize struct{}

func (peccoInitialize PeccoInitialize) Create(request CreatePeccoInitializeRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/pecco/initialize", options)
	return response
}
