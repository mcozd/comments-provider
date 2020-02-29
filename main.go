package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

//Example: https://jsonplaceholder.typicode.com/users/1
const UserInfoBaseUrl = "https://jsonplaceholder.typicode.com/users/"

//Example: https://jsonplaceholder.typicode.com/posts?userId=1
const UserCommentsBaseUrl = "https://jsonplaceholder.typicode.com/posts"

//EXAMPLE
//type people struct {
//	Number int `json:"number"`
//}

type userInfo struct {
	Id       int     `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  company `json:"company"`
}

//ToDo: Annahme, dass alle Infos interessieren
type address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     geo    `json:"geo"`
}

type geo struct {
	//ToDo: Float? Necessary for usecase?
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

func main() {
	getUserInfo(1)

	//ToDO: Server Stuff, temporaryly disabled
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//})
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUserInfo(userId int) {
	timeout := time.Duration(5 * time.Second)
	var client = http.Client{
		Timeout: timeout,
	}
	// ToDo: Better Concatenation?
	resp, err := client.Get(UserInfoBaseUrl + strconv.Itoa(userId))
	//fmt.Println(resp.Body)

	if err != nil {
		// ToDo: Do Something Better?
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}

	userInfo := userInfo{}
	unmarshallError := json.Unmarshal(body, &userInfo)

	if unmarshallError != nil {
		log.Fatal(unmarshallError)
	}

	fmt.Println(userInfo)
	//resp, err = http.Get( + userId)
	//resp
}
