package AdMediaCards

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CardCreatePayload struct {
	BIN         int     `json:"BIN"`
	Limit       float64 `json:"Limit"`
	FirstName   string  `json:"FirstName"`
	LastName    string  `json:"LastName"`
	Address1    string  `json:"Address1"`
	City        string  `json:"City"`
	State       string  `json:"State"`
	Zip         string  `json:"Zip"`
	CountryIso  string  `json:"CountryIso"`
	ExpMonth    string  `json:"ExpMonth"`
	ExpYear     string  `json:"ExpYear"`
	PhoneNumber string  `json:"PhoneNumber"`
}

type CardCreateResponse struct {
	Status       int    `json:"Status"`
	Good         bool   `json:"Good"`
	Log          string `json:"Log"`
	ErrorMessage string `json:"ErrorMessage"`
	Result       struct {
		IsApproved      bool `json:"IsApproved"`
		CardID          int  `json:"CardID"`
		ChangeRequestID int  `json:"ChangeRequestID"`
	} `json:"Result"`
}

func CardCreate(payload CardCreatePayload) (*CardCreateResponse, error) {
	postBody, _ := json.Marshal(&payload)
	fmt.Println(string(postBody))
	resp, err := doRequest("POST", "open/cards", postBody)
	if err != nil {
		return nil, err
	}
	binAnswer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := &CardCreateResponse{}
	err = json.Unmarshal(binAnswer, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
