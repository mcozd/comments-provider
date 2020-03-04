package main

import (
	"log"
	"net/http"
	"strconv"
	_ "strings"
)

const ServerPort = 8080

func main() {
	http.HandleFunc("/user/", UserFullInfoHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(ServerPort), nil))
}
