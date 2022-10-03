package iyzipay

type Card struct{}

func (card Card) Create(request CreateCardRequest, options Options) string {
	response := makeRequest(request, "POST", "/cardstorage/card", options)
	return response
}

func (card Card) Delete(request DeleteCardRequest, options Options) string {
	response := makeRequest(request, "DELETE", "/cardstorage/card", options)
	return response
}
