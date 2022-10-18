package iyzipay

import (
	"bytes"
	//nolint //G505: Blocklisted import crypto/sha1: weak cryptographic primitive
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Request interface {
	toPkiString() string
}

func makeRequest(request Request, methodType string, endpoint string, options Options) string {
	requestJSON, _ := json.Marshal(request)

	return connect(methodType, endpoint, options, string(requestJSON), request.toPkiString())
}

func decodeResponse(response string, into interface{}) interface{} {
	if err := json.Unmarshal([]byte(response), into); err != nil {
		panic(err)
	}

	return into
}

func connect(method string, url string, options Options, requestString, pkiString string) string {
	body := bytes.NewBufferString(requestString)

	randomHeaderValue := RandString(8)

	request, err := http.NewRequest(method, options.baseURL+url, body)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", prepareAuthString(options, randomHeaderValue, pkiString))
	request.Header.Set("x-iyzi-rnd", randomHeaderValue)

	client := http.Client{}

	rawResponse, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer rawResponse.Body.Close()

	bodyByte, err := io.ReadAll(rawResponse.Body)
	if err != nil {
		panic(err)
	}

	return string(bodyByte)
}

func prepareAuthString(options Options, randomStr string, pkiString string) string {
	return formatHeaderString(
		options.apiKey,
		generateHash(options.apiKey, options.secretKey, randomStr, pkiString))
}

func formatHeaderString(apiKey string, hashed string) string {
	return "IYZWS " + apiKey + ":" + hashed
}

func generateHash(apiKey string, secretKey string, randomString string, pkiString string) string {
	//nolint // G401: Use of weak cryptographic primitive
	hasher := sha1.New()
	hasher.Write([]byte(apiKey + randomString + secretKey + pkiString))
	hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	hash = strings.Replace(hash, "_", "/", -1)
	hash = strings.Replace(hash, "-", "+", -1)

	return hash
}

func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	random := make([]rune, n)
	for i := range random {
		//nolint // G404: Use of weak random number generator (math/rand instead of crypto/rand)
		random[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(random)
}
