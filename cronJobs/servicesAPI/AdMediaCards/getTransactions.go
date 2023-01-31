package AdMediaCards

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type GetTransactionsResponse struct {
	Status int  `json:"Status"`
	Good   bool `json:"Good"`
	Result []struct {
		TransactionID     uint        `json:"TransactionID"`
		ClientID          uint        `json:"ClientID"`
		CardID            uint        `json:"CardID"`
		BIN               string      `json:"BIN"`
		Last4             string      `json:"Last4"`
		Merchant          string      `json:"Merchant"`
		TransactionDate   int64       `json:"TransactionDate"`
		Amount            float64     `json:"Amount"`
		AmountUi          float64     `json:"AmountUi"`
		Response          interface{} `json:"Response"`
		ResponseText      string      `json:"ResponseText"`
		Name              string      `json:"Name"`
		Address           string      `json:"Address"`
		LocalAmount       int         `json:"LocalAmount"`
		LocalCurrency     string      `json:"LocalCurrency"`
		Type              string      `json:"Type"`
		IsTemp            bool        `json:"IsTemp"`
		Country           string      `json:"Country"`
		IsSettleAvailable bool        `json:"IsSettleAvailable"`
		CompanyName       string      `json:"CompanyName"`
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

// GetTransactions get transactions from API; if id is 0 gets all transactions
func GetTransactions(id int, search ...string) (*GetTransactionsResponse, error) {
	uri := "open/transactions"
	if id != 0 {
		uri = "open/cards/" + strconv.Itoa(id) + "/transactions/"
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
	data := &GetTransactionsResponse{}
	err = json.Unmarshal(binAnswer, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
