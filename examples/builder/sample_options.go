package builder

import (
	iyzipay "github.com/kurtulussahin/iyzipay-go/iyzipay"
)

func CreateSampleOptions() iyzipay.Options {

	api_key := "sandbox-afXhZPW0MQlE4dCUUlHcEopnMBgXnAZI"
	secret_key := "sandbox-wbwpzKIiplZxI3hh5ALI4FJyAcZKL6kq"
	base_url := "https://sandbox-api.iyzipay.com"

	options := iyzipay.Options{}
	options.New(api_key, secret_key, base_url)

	return options
}
