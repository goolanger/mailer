package scheduler

import (
	"time"

	"../mailer"
)

type ScheduleConfig struct {
	Threads int
	ApiWait int
}

func (config ScheduleConfig) Schedule(api mailer.MailerApi, mails []mailer.MailerConfig) {
	bufferSize := len(mails)

	// Create params channel
	params := make(chan mailer.MailerConfig, bufferSize)
	// Create results channel
	results := make(chan struct {
		mailer.MailerConfig
		bool
	}, bufferSize)

	for i := 0; i < config.Threads; i++ {
		go worker(params, results)
	}

	for mail := range mails {
		params <- mails[mail]
	}
	close(params)

	for i := 0; i < bufferSize; i++ {
		result := <-results
		if !result.bool {
			api.Update(result.MailerConfig)
			time.Sleep(time.Millisecond * time.Duration(config.ApiWait))
		}
	}

}

func worker(params <-chan mailer.MailerConfig, results chan<- struct {
	mailer.MailerConfig
	bool
}) {
	for param := range params {
		results <- struct {
			mailer.MailerConfig
			bool
		}{param, param.SendMail()}
	}
}
