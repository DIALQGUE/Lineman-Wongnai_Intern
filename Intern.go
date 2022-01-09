//intern.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type covid_case struct {
	ConfirmDate    string      `json:"confirmDate"`
	No             interface{} `json:"No"`
	Age            int         `json:"Age"`
	Gender         string      `json:"Gender"`
	GenderEn       string      `json:"GenderEn"`
	Nation         string      `json:"Nation"`
	NationEn       string      `json:"NationEn"`
	Province       string      `json:"Province"`
	ProvinceId     int         `json:"NationEnId"`
	District       string      `json:"District"`
	ProvinceEn     string      `json:"ProvinceEn"`
	StatQuarantine int         `json:"StatQuarantine"`
}

var covid_cases []covid_case

func readJSON(filename string, receiver *interface{}) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERROR: cannot open JSON file")
		return
	} else {
		fmt.Println("ALERT: open JSON file successfully")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), receiver)
}

func mapDecode(m interface{}, target *[]covid_case) {
	for _, j := range m.([]interface{}) {
		var result covid_case
		mapstructure.Decode(j, &result)
		covid_cases = append(covid_cases, result)
	}
}

func main() {
	//read file to interface
	var i interface{}
	readJSON("covid-case.json", &i)

	//extract Data map out of interface
	Data := i.(map[string]interface{})["Data"]

	mapDecode(Data, &covid_cases)

	router := gin.Default()

	router.Run("localhost:8080")
}
