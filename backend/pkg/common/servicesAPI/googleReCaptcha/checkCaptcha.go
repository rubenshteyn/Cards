package googleReCaptcha

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type captchaResponse struct {
	Success     bool          `json:"success"`
	ChallengeTs string        `json:"challenge_ts"`
	Hostname    string        `json:"hostname"`
	ErrorCodes  []interface{} `json:"error-codes"`
}

func CheckCaptcha(token string, ip string) (*captchaResponse, error) {
	secretKey := getSecretKey()
	// siteKey := getSiteKey()
	client := &http.Client{}
	urlEncode := "?secret=" + url.QueryEscape(secretKey) + "&response=" + url.QueryEscape(token) + "&remoteip=" + url.QueryEscape(ip)
	urlDomain := getDomain() + urlEncode
	fmt.Println(urlDomain)
	req, err := http.NewRequest("POST", urlDomain, strings.NewReader(""))
	req.Header.Set("Content-Type", "application/x-www-from-urlencoded; charset=utf-8")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	binAnswer, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	data := &captchaResponse{}
	err = json.Unmarshal(binAnswer, data)
	fmt.Println(data)
	return data, nil
}
