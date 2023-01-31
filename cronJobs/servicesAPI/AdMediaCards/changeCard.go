package AdMediaCards

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type ChangeCardPayload struct {
	Type      string  `json:"Type"`
	Limit     float64 `json:"Limit"`
	Status    bool    `json:"Status"`
	FirstName string  `json:"FirstName"`
	LastName  string  `json:"LastName"`
	Address1  string  `json:"Address1"`
	City      string  `json:"City"`
	State     string  `json:"State"`
	Country   string  `json:"Country"`
	Zip       string  `json:"Zip"`
}

type ChangeCardResponse struct {
	Status int  `json:"Status"`
	Good   bool `json:"Good"`
	Result struct {
		IsApproved      bool `json:"IsApproved"`
		ChangeRequestID int  `json:"ChangeRequestID"`
	} `json:"Result"`
}

func ChangeCard(id int, payload ChangeCardPayload) (*ChangeCardResponse, error) {
	postBody, _ := json.Marshal(&payload)
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
