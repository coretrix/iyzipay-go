package iyzipay

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func connectV2(method, uri, apiKey, secretKey string, req interface{}) (string, error) {
	randomString := uuid.New().String()
	authHeader, err := prepareAuthorizationStringV2(uri, apiKey, secretKey, randomString, req)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return "", err
	}

	if req != nil {
		marshaled, err := json.Marshal(req)
		if err != nil {
			return "", err
		}

		request, err = http.NewRequest(method, uri, bytes.NewBuffer(marshaled))
		if err != nil {
			return "", err
		}
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", authHeader)

	client := http.Client{}
	rawResponse, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer rawResponse.Body.Close()

	bodyByte, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return "", err
	}

	return string(bodyByte), nil
}

func prepareAuthorizationStringV2(uri, apiKey, secretKey, randomString string, request interface{}) (string, error) {
	authContent, err := generateAuthContent(uri, apiKey, secretKey, randomString, request)
	if err != nil {
		return "", err
	}
	return "IYZWSv2 " + authContent, nil
}

func generateAuthContent(uri, apiKey, secretKey, randomString string, request interface{}) (string, error) {
	hmac256, err := getHmacSHA256Signature(uri, secretKey, randomString, request)
	if err != nil {
		return "", err
	}

	input := "apiKey:" + apiKey + "&randomKey:" + randomString + "&signature:" + hmac256
	return base64.StdEncoding.EncodeToString([]byte(input)), nil
}

func getHmacSHA256Signature(uri, secretKey, randomString string, request interface{}) (string, error) {
	payload, err := getPayload(uri, request)
	if err != nil {
		return "", err
	}

	dataToSign := randomString + payload

	key := []byte(secretKey)

	sig := hmac.New(sha256.New, key)
	sig.Write([]byte(dataToSign))

	return strings.ToLower(hex.EncodeToString(sig.Sum(nil))), nil
}

func getPayload(uri string, request interface{}) (string, error) {
	startIndex := strings.Index(uri, "/v2")
	endIndex := strings.Index(uri, "?")

	uriPath := ""
	if endIndex == -1 {
		uriPath = uri[startIndex:]
	} else {
		uriPath = uri[startIndex:endIndex]
	}

	if request != nil {
		marshaled, err := json.Marshal(request)
		if err != nil {
			return "", err
		}

		return uriPath + string(marshaled), nil
	}

	return uriPath, nil
}
