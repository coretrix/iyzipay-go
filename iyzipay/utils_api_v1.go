package iyzipay

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Request interface {
	toPkiString() string
}

func makeRequest(request Request, methodType string, endpoint string, options Options) string {
	pkiString := request.toPkiString()
	requestJson, _ := json.Marshal(request)
	requestJsonString := string(requestJson)

	response := connect(methodType, endpoint, options, requestJsonString, pkiString)

	return response
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

	request, err := http.NewRequest(method, options.baseUrl+url, body)
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
	bodyByte, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		panic(err)
	}
	return string(bodyByte)
}

func prepareAuthString(options Options, random_str string, pki_string string) string {
	hashed := generateHash(options.apiKey, options.secretKey, random_str, pki_string)
	return formatHeaderString(options.apiKey, hashed)
}

func formatHeaderString(api_key string, hashed string) string {
	return "IYZWS " + api_key + ":" + hashed
}

func generateHash(apiKey string, secretKey string, random_string string, pki_string string) string {
	hash_str := apiKey + random_string + secretKey + pki_string
	hasher := sha1.New()
	hasher.Write([]byte(hash_str))
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
		random[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(random)
}
