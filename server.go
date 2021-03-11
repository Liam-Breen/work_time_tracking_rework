package main

import (
	"fmt"
	"net/http"
	"strings"
)

func UserServer(w http.ResponseWriter, r *http.Request) {

	user := strings.TrimPrefix(r.URL.Path, "/user/")
	fmt.Fprint(w, GetUserHours(user))

}

func GetUserHours(user string) string {
	if user == "LiamBreen" {
		return "10"
	}

	if user == "JamesBos" {
		return "20"
	}

	return ""
}
