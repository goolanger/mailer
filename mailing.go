package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/akamensky/argparse"
	"github.com/gorilla/mux"

	"./src/utils"
)

func main() {
	parser := argparse.NewParser("mailing", "Mailing provides a simple and robust email solution to newsletter subscriptions and notices. Implements a variety of algorithms that consist in concurrency tasks managements and fake mail detections.")

	token := parser.String("k", "token", &argparse.Options{Help: "JSON Web Token (JWT) for validation"})
	port := parser.String("p", "port", &argparse.Options{Help: "HTTP TCP Port to comunicate with the API"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	utils.CleanParameterString(token, utils.GenerateToken())
	utils.CleanParameterString(port, "5500")

	router := mux.NewRouter()

	// GET /token { token }
	router.HandleFunc("/token", func(res http.ResponseWriter, req *http.Request) {
		if !utils.ValidToken(req, *token) {
			utils.AccessDenied(res)
			return
		}

		json.NewEncoder(res).Encode(struct{ Token string }{Token: *token})
	}).Methods("GET")

	// PUT /token { token, newToken }
	router.HandleFunc("/set-token", func(res http.ResponseWriter, req *http.Request) {

	}).Methods("PUT")

	// PUT /send { token, mails: [...], message, smtp.options, threads, retry }
	router.HandleFunc("/send", func(res http.ResponseWriter, req *http.Request) {

	}).Methods("PUT")

	log.Fatal(http.ListenAndServe(":"+*port, router))
}
