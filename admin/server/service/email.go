package service

import (
	"admin/core/log"
	"admin/server/models"
	"fmt"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"net/smtp"
)

type Email struct {
	ID uint

	Host     string
	Port     int
	Sender   string
	Password string
}

var emailCache *Email

func (e *Email) Save() error {
	data := make(map[string]interface{})
	data["host"] = e.Host
	data["port"] = e.Port
	data["sender"] = e.Sender
	data["password"] = e.Password

	if e.ID > 0 {
		emailCache = nil
		return models.UpdateEmail(e.ID, data)
	}

	return models.AddEmail(data)
}

func (e *Email) Get() (*models.Email, error) {
	return models.GetEmail()
}

func SendMail(receiver, sub string, body []byte) error {
	if emailCache == nil {
		emailConfig, err := models.GetEmail()
		if err != nil {
			return err
		}
		emailCache = new(Email)
		emailCache.Host = emailConfig.Host
		emailCache.Port = emailConfig.Port
		emailCache.Sender = emailConfig.Sender
		emailCache.Password = emailConfig.Password
	}

	go func() {
		auth := smtp.PlainAuth("", emailCache.Sender, emailCache.Password, emailCache.Host)
		e := &email.Email{
			From:    fmt.Sprintf("%s", emailCache.Sender),
			To:      []string{receiver},
			Subject: sub,
			Text:    body,
		}
		err := e.Send(emailCache.Host+":"+fmt.Sprintf("%d", emailCache.Port), auth)
		if err != nil {
			log.Logger.Error("email", zap.String("err", err.Error()))
		}
	}()

	return nil
}
