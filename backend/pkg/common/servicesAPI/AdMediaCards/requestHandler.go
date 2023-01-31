package AdMediaCards

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func doRequest(method string, uri string, payload []byte) (*http.Response, error) {
	domain := getDomain()
	apiKey := getAPIKey()
	url := domain + uri
	fmt.Println("method = " + method)
	fmt.Println("url = " + url)
	client := &http.Client{}
	if method == "GET" {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("content-type", "application/json")
		fmt.Println("content-type: application/json")
		req.Header.Set("Authorization", "ApiKey "+apiKey)
		fmt.Println("Authorization: ApiKey" + apiKey)
		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")
	fmt.Println("content-type: application/json")
	req.Header.Set("Authorization", "ApiKey "+apiKey)
	fmt.Println("Authorization: ApiKey" + apiKey)
	if method == "POST" {
		uuidKey := uuid.New().String()
		req.Header.Set("Idempotency-Key", uuidKey)
		fmt.Println("Idempotency-Key: " + uuidKey)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
