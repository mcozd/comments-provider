package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var client = http.Client{
	Timeout: time.Duration(5 * time.Second),
}

func getUserInfo(userId int, c chan userInfo) {
	// ToDo: Better Concatenation?
	resp, httpCallError := client.Get(UserInfoBaseUrl + strconv.Itoa(userId))
	handleError(httpCallError)

	body, readError := ioutil.ReadAll(resp.Body)
	handleError(readError)

	userInfo := userInfo{}
	unmarshalError := json.Unmarshal(body, &userInfo)
	handleError(unmarshalError)

	c <- userInfo
}

func getUserComments(userId int, c chan []comment) {
	//ToDo: Pass query parameter differently?
	resp, err := client.Get(UserCommentsBaseUrl + strconv.Itoa(userId))
	handleError(err)

	body, readError := ioutil.ReadAll(resp.Body)
	handleError(readError)

	comments := []comment{}
	unmarshalError := json.Unmarshal(body, &comments)
	handleError(unmarshalError)

	c <- comments
}

func handleError(err error) {
	if err != nil {
		// ToDo: Do Something Better?
		log.Fatal(err)
	}
}
