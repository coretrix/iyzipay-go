package iyzipay

type CardList struct{}

func (cardList CardList) Retrieve(request RetrieveCardListRequest, options Options) string {
	return makeRequest(request, "POST", "/cardstorage/cards", options)
}
