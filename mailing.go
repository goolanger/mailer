package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"

	"./src/mailer"
	"./src/scheduler"
	"./src/utils"
)

func main() {
	parser := argparse.NewParser("mailing", "Mailing provides a simple and robust email solution to newsletter subscriptions and notices. Implements a variety of algorithms that consist in concurrency tasks managements and fake mail detections.")

	token := parser.String("k", "token", &argparse.Options{Help: "JSON Web Token (JWT) for validation"})
	wait := parser.Int("w", "wait", &argparse.Options{Help: "Wait time between API updates (Given in Milliseconds)"})
	thr := parser.Int("t", "threads", &argparse.Options{Help: "Number of concurrent threads"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	utils.CleanParameterString(token, utils.GenerateToken())
	utils.CleanParameterInt(thr, 2)

	fmt.Println("Running on http://localhost:8080/mailing?token=" + *token)

	// Testing
	mail := "This is a test mail..."

	api := mailer.MailerApi{
		Get: "http://localhost:3000/mails/api",
		Put: "http://localhost:3000/mails/api/%s/put",
	}

	mails := []mailer.MailerConfig{
		mailer.MailerConfig{Mail: &mail, Email: "amauryuh@gmail.com", Fails: 2, Id: 1},
		mailer.MailerConfig{Mail: &mail, Email: "a.caballero@estudiantes.matcom.uh.cu", Fails: 2, Id: 2},
		mailer.MailerConfig{Mail: &mail, Email: "yavseny.roque@matcom.uh.cu", Fails: 2, Id: 2},
	}

	sched := scheduler.ScheduleConfig{
		ApiWait: *wait,
		Threads: *thr,
	}
	sched.Schedule(api, mails)

	fmt.Scanln()
}
