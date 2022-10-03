package iyzipay

type BkmInitialize struct{}

func (bkmInitialize BkmInitialize) Create(request CreateBkmInitializeRequest, options Options) string {
	response := makeRequest(request, "POST", "/payment/bkm/initialize", options)
	return response
}
