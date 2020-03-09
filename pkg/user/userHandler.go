package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func parseID(r *http.Request) int {
	path := strings.Split(r.URL.Path, "/")
	userID, _ := strconv.Atoi(path[2])
	return userID
}

func collectUserFullInfo(userID int) userFullInfo {
	c := make(chan userInfo)
	cc := make(chan []comment)
	go getUserInfo(userID, c)
	go getUserComments(userID, cc)

	return userFullInfo{<-c, <-cc}
}

func UserFullInfoHandler(w http.ResponseWriter, r *http.Request) {
	userID := parseID(r)
	user := collectUserFullInfo(userID)
	err := json.NewEncoder(w).Encode(user)
	handleError(err)
}
