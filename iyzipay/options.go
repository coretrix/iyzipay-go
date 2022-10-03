package iyzipay

import ()

type Options struct {
	apiKey    string
	secretKey string
	baseUrl   string
}

func (options *Options) New(apiKey string, secretKey string, baseUrl string) {
	options.apiKey = apiKey
	options.secretKey = secretKey
	options.baseUrl = baseUrl
}
