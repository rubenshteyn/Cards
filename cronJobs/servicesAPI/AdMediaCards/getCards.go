package AdMediaCards

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type GetCardsResponse struct {
	Status int  `json:"Status"`
	Good   bool `json:"Good"`
	Result []struct {
		CardID      int         `json:"CardID"`
		Name        string      `json:"Name"`
		Limit       float64     `json:"Limit"`
		BIN         string      `json:"BIN"`
		Last4       string      `json:"Last4"`
		Balance     float64     `json:"Balance"`
		Address     string      `json:"Address"`
		IsActive    bool        `json:"IsActive"`
		ClientNote  interface{} `json:"ClientNote"`
		Currency    string      `json:"Currency"`
		DateEntered string      `json:"DateEntered"`
		ExpMonth    string      `json:"ExpMonth"`
		ExpYear     string      `json:"ExpYear"`
		Spend       float64     `json:"Spend"`
	} `json:"Result"`
	Metadata struct {
		Count      int  `json:"Count"`
		TotalCount int  `json:"TotalCount"`
		TotalPages int  `json:"TotalPages"`
		StartIndex int  `json:"StartIndex"`
		EndIndex   int  `json:"EndIndex"`
		IsMore     bool `json:"IsMore"`
	} `json:"Metadata"`
}

// GetCards get cards from API; if id is 0 get all cards with search params
func GetCards(id int, search ...string) (*GetCardsResponse, error) {
	uri := "open/cards"
	if id != 0 {
		uri += "/" + strconv.Itoa(id)
	}
	if len(search) > 0 {
		uri += "?" + search[0]
	}
	resp, err := doRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	binAnswer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := &GetCardsResponse{}
	err = json.Unmarshal(binAnswer, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type GetCardsCVVResponse struct {
	Status int  `json:"Status"`
	Good   bool `json:"Good"`
	Result struct {
		CcNum    string `json:"ccNum"`
		Cvx2     string `json:"cvx2"`
		ExpMonth string `json:"ExpMonth"`
		ExpYear  string `json:"ExpYear"`
	} `json:"Result"`
}

func GetCardCVV(id int) (*GetCardsCVVResponse, error) {
	uri := "open/cards"
	uri += "/" + strconv.Itoa(id) + "/showpan"
	resp, err := doRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	binAnswer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := &GetCardsCVVResponse{}
	err = json.Unmarshal(binAnswer, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
