package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func getUpdates(offset int) ([]Update, error) {

	resp, err := http.Get(baseUrl + "/getUpdates?offset=" + strconv.Itoa(offset))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}

	var updates UpdateResponse

	err = json.Unmarshal(body, &updates)

	if err != nil {
		return nil, err
	}

	return updates.Result, nil
}
