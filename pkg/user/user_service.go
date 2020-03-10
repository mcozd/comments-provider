package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

//Example: https://jsonplaceholder.typicode.com/users/1
const UserInfoBaseURL = "https://jsonplaceholder.typicode.com/users/"

//Example: https://jsonplaceholder.typicode.com/posts?userId=1
const UserCommentsBaseURL = "https://jsonplaceholder.typicode.com/posts?userId="

const timeoutSeconds = 5

var client = http.Client{
	Timeout: time.Duration(timeoutSeconds * time.Second),
}

func getUserInfo(userID int, c chan userInfo) {
	resp, httpCallError := client.Get(UserInfoBaseURL + strconv.Itoa(userID))
	handleError(httpCallError)

	body, readError := ioutil.ReadAll(resp.Body)
	handleError(readError)

	userInfo := userInfo{}
	unmarshalError := json.Unmarshal(body, &userInfo)
	handleError(unmarshalError)

	c <- userInfo
}

func getUserComments(userID int, c chan []comment) {
	resp, err := client.Get(UserCommentsBaseURL + strconv.Itoa(userID))
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
