package iyzipay

type Options struct {
	apiKey    string
	secretKey string
	baseURL   string
}

func (options *Options) New(apiKey string, secretKey string, baseURL string) {
	options.apiKey = apiKey
	options.secretKey = secretKey
	options.baseURL = baseURL
}
