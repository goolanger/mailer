package scheduler

import (
	"time"

	"../mailer"
)

// Config ...
type Config struct {
	Threads int
	Wait    int
}

// Schedule ...
func (config Config) Schedule(mails *[]mailer.Config) []mailer.Config {
	bufferSize := len(*mails)
	var failedMails []mailer.Config

	// Create params channel
	params := make(chan mailer.Config, bufferSize)
	// Create results channel
	results := make(chan struct {
		mailer.Config
		bool
	}, bufferSize)

	for i := 0; i < config.Threads; i++ {
		go worker(params, results)
	}

	for mail := range *mails {
		params <- (*mails)[mail]
	}
	close(params)

	for i := 0; i < bufferSize; i++ {
		result := <-results
		if !result.bool {
			result.Config.Pending--
			failedMails = append(failedMails, result.Config)
		}
	}

	if stillPending(&failedMails) {
		time.Sleep(time.Second * time.Duration(config.Wait))
		return config.Schedule(&failedMails)
	}

	return failedMails
}

func worker(params <-chan mailer.Config, results chan<- struct {
	mailer.Config
	bool
}) {
	for param := range params {
		results <- struct {
			mailer.Config
			bool
		}{param, param.SendMail()}
	}
}

func stillPending(mails *[]mailer.Config) bool {
	for i := range *mails {
		if (*mails)[i].Pending > 0 {
			return true
		}
	}
	return false
}
