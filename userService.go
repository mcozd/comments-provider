package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

//Example: https://jsonplaceholder.typicode.com/users/1
const UserInfoBaseUrl = "https://jsonplaceholder.typicode.com/users/"

//Example: https://jsonplaceholder.typicode.com/posts?userId=1
const UserCommentsBaseUrl = "https://jsonplaceholder.typicode.com/posts?userId="

const TimeoutSeconds = 5

var client = http.Client{
	Timeout: time.Duration(TimeoutSeconds * time.Second),
}

func getUserInfo(userId int, c chan userInfo) {
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
		log.Fatal(err)
	}
}
