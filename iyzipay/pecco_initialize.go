package iyzipay

type PeccoInitialize struct{}

func (peccoInitialize PeccoInitialize) Create(request CreatePeccoInitializeRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/pecco/initialize", options)
}
