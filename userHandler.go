package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func UserFullInfoHandler(w http.ResponseWriter, r *http.Request) {
	userId := parseId(r)
	user := collectUserFullInfo(userId)
	err := json.NewEncoder(w).Encode(user)
	handleError(err)
}

func parseId(r *http.Request) int {
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
