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
	startupError := http.ListenAndServe(":"+strconv.Itoa(serverPort), nil)

	if startupError != nil {
		log.Fatal(startupError)
	} else {
		log.Printf("Server started successfully on port: %s", strconv.Itoa(serverPort))
	}
}
