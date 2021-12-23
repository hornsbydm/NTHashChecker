package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const URL = "https://api.nthashes.com/search/"

func queryHash(h string) string {

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	res, err := client.Get(URL + h[:5])

	if err != nil {
		log.Fatal("Bad Request")
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal("Invalid Response Body")
	}

	return string(body)
}
