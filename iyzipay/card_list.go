package iyzipay

type CardList struct{}

func (cardList CardList) Retrieve(request RetrieveCardListRequest, options Options) string {
	response := makeRequest(request, "POST", "/cardstorage/cards", options)
	return response
}
