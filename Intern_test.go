package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSummary(t *testing.T) {
	response, err := http.Get("http://localhost:8080/covid/summary")
	if err != nil || response.StatusCode != http.StatusOK {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseByte, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseMap map[string]interface{}
	json.Unmarshal([]byte(responseByte), &responseMap)

	var i interface{}
	readJSON("summary.json", &i)
	expected := i.(map[string]interface{})

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, responseMap)
}
