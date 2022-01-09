package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

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
