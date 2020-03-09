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
	http.HandleFunc("/user/", ownUser.UserFullInfoHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(serverPort), nil))
}
