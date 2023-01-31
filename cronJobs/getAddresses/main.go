package getAddresses

import (
	"encoding/json"
	"fmt"
	"github.com/qrds1/cronJobs/models"
	"github.com/qrds1/cronJobs/updateCards"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type address struct {
}

// Addresses struct which contains
// an array of addresses
type Addresses struct {
	Addresses []Address `json:"addresses"`
}

// Address struct
type Address struct {
	Number   int         `json:"number"`
	Street   string      `json:"street"`
	City     string      `json:"city"`
	Region   string      `json:"region"`
	Postcode interface{} `json:"postcode"`
}

func AddressesByJSON(h *updateCards.Handler) {
	// Open our jsonFile
	jsonFile, err := os.Open("getAddresses/csvjson-2.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Opened csvjson-2.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(jsonFile)
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// we initialize our addresses array
	var addresses Addresses
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &addresses)
	if err != nil {
		fmt.Println(err)
		return
	}
	areaCodes := []string{"602", "520", "480", "623", "928"}
	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i, address := range addresses.Addresses {
		fmt.Println(strconv.Itoa(i))
		var ourAddress models.Address
		address1 := strconv.Itoa(address.Number) + " " + address.Street
		/*if result := h.DB.Where("address1 = ?", address1).First(&ourAddress); result.Error == nil {
			continue
		}*/
		ourAddress.Address1 = address1
		ourAddress.Zip = strconv.Itoa(int(address.Postcode.(float64)))
		ourAddress.State = address.Region
		ourAddress.City = address.City
		rand.Seed(time.Now().UnixNano())
		v := rand.Intn(8999999) + 1000000
		ourAddress.PhoneNumber = "+1" + areaCodes[(i+5)%5] + strconv.Itoa(v)
		if result := h.DB.Create(&ourAddress); result.Error != nil {
			fmt.Println("address not created " + result.Error.Error())
		}
	}
}
