package iyzipay

import (
	"fmt"
	"strconv"
)

type PkiBuilder struct {
	requestString string
}

func (pkiBuilder *PkiBuilder) append(key string, value string) {
	if value != "" {
		pkiBuilder.appendKeyValue(key, value)
	}
}

func (pkiBuilder *PkiBuilder) appendArray(key string, array []string) {
	if array != nil {
		appendedValue := ""
		for _, value := range array {
			appendedValue = appendedValue + value + ", "
		}

		pkiBuilder.appendKeyValueArray(key, appendedValue)
	}
}

func (pkiBuilder *PkiBuilder) appendKeyValue(key string, value string) {
	pkiBuilder.requestString = pkiBuilder.requestString + key + "=" + value + ","
}

func (pkiBuilder *PkiBuilder) appendKeyValueArray(key string, appendedValue string) {
	appendedValue = appendedValue[:len(appendedValue)-2]
	pkiBuilder.requestString = pkiBuilder.requestString + key + "=[" + appendedValue + "],"
}

func (pkiBuilder *PkiBuilder) appendPrice(key string, value string) {
	if value != "" {
		valueToFloat, _ := strconv.ParseFloat(value, 64)
		roundedStringValue := fmt.Sprintf("%.2f", valueToFloat)

		if (roundedStringValue[len(roundedStringValue)-1:]) == "0" {
			roundedStringValue = roundedStringValue[:len(roundedStringValue)-1]
		}

		pkiBuilder.appendKeyValue(key, roundedStringValue)
	}
}

func (pkiBuilder *PkiBuilder) getRequestString() string {
	pkiBuilder.removeTrailingComma()
	pkiBuilder.appendPrefix()

	return pkiBuilder.requestString
}

func (pkiBuilder *PkiBuilder) removeTrailingComma() {
	pkiBuilder.requestString = pkiBuilder.requestString[:len(pkiBuilder.requestString)-1]
}

func (pkiBuilder *PkiBuilder) appendPrefix() {
	pkiBuilder.requestString = "[" + pkiBuilder.requestString + "]"
}
