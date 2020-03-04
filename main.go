package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	_ "strings"
)

const ServerPort = 8080

//ToDo: Already used in golang pkg net/url !?
type userInfo struct {
	Id       int     `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  company `json:"company"`
}

//ToDo: assumption every field is important
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

type comment struct {
	UserId int    `json:"userId"`
	Id     int    `json:"Id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type userFullInfo struct {
	UserInfo userInfo  `json:"userInfo"`
	Comments []comment `json:"comments"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		userId := parseId(r)
		user := collectUserFullInfo(userId)
		err := json.NewEncoder(w).Encode(user)

		handleError(err)
	})
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(ServerPort), nil))
}

func parseId(r *http.Request) int {
	//ToDo: Be more robust?
	path := strings.Split(r.URL.Path, "/")
	userId, _ := strconv.Atoi(path[1])
	return userId
}

func collectUserFullInfo(userId int) userFullInfo {
	c := make(chan userInfo)
	cc := make(chan []comment)
	go getUserInfo(userId, c)
	go getUserComments(userId, cc)

	return userFullInfo{<-c, <-cc}
}
