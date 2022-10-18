package iyzipay

type APITest struct{}

func (apiTest APITest) Retrieve(options Options) string {
	return connect("GET", "/payment/test", options, "", "")
}
