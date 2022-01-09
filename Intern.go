//intern.go
package main

import (
	"net/http"

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

func mapDecode(m interface{}, target *[]covid_case) {
	for _, j := range m.([]interface{}) {
		var result covid_case
		mapstructure.Decode(j, &result)
		covid_cases = append(covid_cases, result)
	}
}

func GetSummary(c *gin.Context) {
	Province, AgeGroup := make(map[string]int), make(map[string]int)

	for _, Case := range covid_cases {
		p := Case.ProvinceEn
		a := Case.Age

		_, exist := Province[p]
		if exist {
			Province[p]++
		} else {
			Province[p] = 1
		}

		switch {
		case a == 0:
			AgeGroup["N/A"]++
		case a <= 30:
			AgeGroup["0-30"]++
		case a <= 60:
			AgeGroup["31-60"]++
		default:
			AgeGroup["61+"]++
		}
	}

	Province["N/A"] = Province[""]
	delete(Province, "")

	summary := make(map[string](map[string]int))
	summary["Province"] = Province
	summary["AgeGroup"] = AgeGroup

	c.IndentedJSON(http.StatusOK, summary)
}

func main() {
	//read file to interface
	var i interface{}
	readJSON("covid-case.json", &i)

	//extract Data map out of interface
	Data := i.(map[string]interface{})["Data"]

	mapDecode(Data, &covid_cases)

	router := gin.Default()
	router.GET("covid/summary", GetSummary)

	router.Run("localhost:8080")
}
