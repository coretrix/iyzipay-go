package iyzipay

type BkmInitialize struct{}

func (bkmInitialize BkmInitialize) Create(request CreateBkmInitializeRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/bkm/initialize", options)
}
