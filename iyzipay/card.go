package iyzipay

type Card struct{}

func (card Card) Create(request CreateCardRequest, options Options) string {
	return makeRequest(request, "POST", "/cardstorage/card", options)
}

func (card Card) Delete(request DeleteCardRequest, options Options) string {
	return makeRequest(request, "DELETE", "/cardstorage/card", options)
}
