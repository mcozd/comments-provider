package main

import (
	ownUser "comments-provider/pkg/user"
	"log"
	"net/http"
	"strconv"
	_ "strings"

)

const serverPort = 8080

func main() {
	http.HandleFunc("/users/", ownUser.UserFullInfoHandler)
	error := http.ListenAndServe(":"+strconv.Itoa(serverPort), nil)
	if error != nil {
		log.Fatal(error)
	} else {
		println("Server started successfully on port:" + strconv.Itoa(serverPort))
	}
}
