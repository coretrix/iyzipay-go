package iyzipay

import (
	"fmt"
	"strconv"
)

type PkiBuilder struct {
	request_string string
}

func (pkiBuilder *PkiBuilder) append(key string, value string) {
	if value != "" {
		pkiBuilder.appendKeyValue(key, value)
	}
}

func (pkiBuilder *PkiBuilder) appendArray(key string, array []string) {
	if array != nil {
		appended_value := ""
		for _, value := range array {
			appended_value = appended_value + value + ", "
		}
		pkiBuilder.appendKeyValueArray(key, appended_value)
	}
}

func (pkiBuilder *PkiBuilder) appendKeyValue(key string, value string) {
	pkiBuilder.request_string = pkiBuilder.request_string + key + "=" + value + ","
}

func (pkiBuilder *PkiBuilder) appendKeyValueArray(key string, appended_value string) {

	appended_value = appended_value[:len(appended_value)-2]
	pkiBuilder.request_string = pkiBuilder.request_string + key + "=[" + appended_value + "],"

}

func (pkiBuilder *PkiBuilder) appendPrice(key string, value string) {
	if value != "" {
		value_to_float, _ := strconv.ParseFloat(value, 64)
		raunded_string_value := string(fmt.Sprintf("%.2f", value_to_float))

		if (raunded_string_value[len(raunded_string_value)-1:]) == "0" {
			raunded_string_value = raunded_string_value[:len(raunded_string_value)-1]
		}

		pkiBuilder.appendKeyValue(key, raunded_string_value)
	}
}

func (pkiBuilder *PkiBuilder) getRequestString() string {
	pkiBuilder.removeTrailingComma()
	pkiBuilder.appendPrefix()
	return pkiBuilder.request_string
}

func (pkiBuilder *PkiBuilder) removeTrailingComma() {
	pkiBuilder.request_string = pkiBuilder.request_string[:len(pkiBuilder.request_string)-1]

}

func (pkiBuilder *PkiBuilder) appendPrefix() {
	pkiBuilder.request_string = "[" + pkiBuilder.request_string + "]"

}
