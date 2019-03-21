package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status int
}

func ValidToken(req *http.Request, token string) bool {
	usertoken := req.Header.Get("User-Token")
	fmt.Println(token, usertoken, token == usertoken)
	return token == usertoken
}

func AccessDenied(res http.ResponseWriter) {
	res.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(res).Encode(Response{Status: 401})
}
