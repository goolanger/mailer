package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/akamensky/argparse"
	"github.com/gorilla/mux"

	"./src/mailer"
	"./src/scheduler"
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

	// PUT /set-token { token, newToken }
	router.HandleFunc("/set-token", func(res http.ResponseWriter, req *http.Request) {
		if !utils.ValidToken(req, *token) {
			utils.AccessDenied(res)
			return
		}

		body := json.NewDecoder(req.Body)
		data := struct{ Token string }{}
		err := body.Decode(&data)

		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		*token = data.Token

		json.NewEncoder(res).Encode(struct{ Token string }{Token: *token})
	}).Methods("PUT")

	// PUT /send { token, mails: [...], message, smtp.options, threads, retry }
	router.HandleFunc("/send", func(res http.ResponseWriter, req *http.Request) {
		if !utils.ValidToken(req, *token) {
			utils.AccessDenied(res)
			return
		}

		body := json.NewDecoder(req.Body)
		data := struct {
			SMTPOptions mailer.SMTPOption
			Mail        string
			Emails      []string
			Threads     int
			Retry       int
			Wait        int
		}{}
		err := body.Decode(&data)

		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		config := scheduler.Config{Wait: data.Wait, Threads: data.Threads}

		var mails []mailer.Config

		for i := range data.Emails {
			mails = append(mails, mailer.Config{Mail: &data.Mail, Email: data.Emails[i], Pending: data.Retry, Opt: &data.SMTPOptions})
		}

		fails := config.Schedule(&mails)

		json.NewEncoder(res).Encode(struct{ Failed []string }{Failed: mailer.FilterMails(&fails)})
	}).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+*port, router))
}
