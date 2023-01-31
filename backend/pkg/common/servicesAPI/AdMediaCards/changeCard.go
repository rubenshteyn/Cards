package AdMediaCards

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type ChangeCardLimitPayload struct {
	Type  string  `json:"Type"`
	Limit float64 `json:"Limit"`
}

type ChangeCardStatusPayload struct {
	Type   string `json:"Type"`
	Status bool   `json:"Status"`
}

type ChangeCardNameAddressPayload struct {
	Type      string `json:"Type"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Address1  string `json:"Address1"`
	City      string `json:"City"`
	State     string `json:"State"`
	Country   string `json:"Country"`
	Zip       string `json:"Zip"`
}

type ChangeCardResponse struct {
	Status       int    `json:"Status"`
	Good         bool   `json:"Good"`
	ErrorMessage string `json:"ErrorMessage"`
	Result       struct {
		IsApproved      bool `json:"IsApproved"`
		ChangeRequestID int  `json:"ChangeRequestID"`
	} `json:"Result"`
}

func ChangeCard(id int, postBody []byte) (*ChangeCardResponse, error) {
	resp, err := doRequest("PUT", "open/cards/"+strconv.Itoa(id), postBody)
	if err != nil {
		return nil, err
	}
	binAnswer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := &ChangeCardResponse{}
	err = json.Unmarshal(binAnswer, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ChangeCardLimit(id int, payload ChangeCardLimitPayload) (*ChangeCardResponse, error) {
	postBody, _ := json.Marshal(&payload)
	fmt.Println(string(postBody))
	card, err := ChangeCard(id, postBody)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func ChangeCardStatus(id int, payload ChangeCardStatusPayload) (*ChangeCardResponse, error) {
	postBody, _ := json.Marshal(&payload)
	fmt.Println(string(postBody))
	card, err := ChangeCard(id, postBody)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func ChangeCardNameAddress(id int, payload ChangeCardNameAddressPayload) (*ChangeCardResponse, error) {
	postBody, _ := json.Marshal(&payload)
	fmt.Println(string(postBody))
	card, err := ChangeCard(id, postBody)
	if err != nil {
		return nil, err
	}
	return card, nil
}
