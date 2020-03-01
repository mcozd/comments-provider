package main

import (
	"fmt"
)

//Example: https://jsonplaceholder.typicode.com/users/1
const UserInfoBaseUrl = "https://jsonplaceholder.typicode.com/users/"

//Example: https://jsonplaceholder.typicode.com/posts?userId=1
const UserCommentsBaseUrl = "https://jsonplaceholder.typicode.com/posts?userId="

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
	i := 1
	c := make(chan userInfo)
	cc := make(chan []comment)
	go getUserInfo(i, c)
	go getUserComments(i, cc)

	var response = userFullInfo{<-c, <-cc}
	fmt.Println(response)

	//ToDO: Server Stuff, temporaryly disabled
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//})
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
