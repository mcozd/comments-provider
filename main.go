package main

import (
	"log"
	"net/http"
	"strconv"
	_ "strings"

)

const serverPort = 8080

func main() {
	http.HandleFunc("/user/", UserFullInfoHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(serverPort), nil))
}
