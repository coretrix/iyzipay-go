package iyzipay

type BinNumber struct{}

func (binNumber BinNumber) Retrieve(request RetrieveBinNumberRequest, options Options) string {
	return makeRequest(request, "POST", "/payment/bin/check", options)
}
