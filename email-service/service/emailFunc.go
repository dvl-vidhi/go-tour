package service

import (
	"context"
	"crypto/tls"
	mail "email-service/emailModel"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	gomail "gopkg.in/mail.v2"
)

type Connection struct {
	Server     string
	Database   string
	Collection string
}

var Collection *mongo.Collection
var ctx = context.TODO()

func (e *Connection) Connect() {
	clientOptions := options.Client().ApplyURI(e.Server)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database(e.Database).Collection(e.Collection)
}

func (e *Connection) SendEmail(mail mail.Mail) (string, error) {

	err := sendMail2(mail)
	fmt.Println(err)
	if err != nil {
		return "", err
	}
	fmt.Println("Mail Sent Succefully")
	mail.Time = time.Now()
	insert, err := Collection.InsertOne(ctx, mail)
	fmt.Println(insert)
	if err != nil {
		return "", errors.New("Unable To Insert New Record")
	}
	return "Email Sent Successfully", nil
}

func sendMail2(mail mail.Mail) error {
	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress("vidhi.goel@gridinfocom.com", "Vidhi")},
		"To":      mail.MailSendTo,
		"Cc":      mail.MailSendCC,
		"Subject": mail.MailSubject,
	})

	m.SetBody("text/plain", mail.MailBody.Salutation+"\n\n"+mail.MailBody.Message+"\n\n"+mail.MailBody.Closing+"\n"+mail.SenderName)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp-relay.sendinblue.com", 587, "vidhi.goel@gridinfocom.com", "pGL756txPrWkSBX4")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (e *Connection) SearchFilter(search mail.Search) ([]*mail.Mail, error) {
	var searchData []*mail.Mail

	filter := bson.D{}

	if search.MailSendTo != "" {
		filter = append(filter, primitive.E{Key: "to", Value: bson.M{"$regex": search.MailSendTo}})
	}
	if search.MailSendCC != "" {
		filter = append(filter, primitive.E{Key: "cc", Value: bson.M{"$regex": search.MailSendCC}})
	}
	if search.MailSendBCC != "" {
		filter = append(filter, primitive.E{Key: "bcc", Value: bson.M{"$regex": search.MailSendBCC}})
	}
	if search.MailSubject != "" {
		filter = append(filter, primitive.E{Key: "subject", Value: bson.M{"$regex": search.MailSubject}})
	}

	t, _ := time.Parse("2006-01-02", search.Date)
	if search.Date != "" {
		filter = append(filter, primitive.E{Key: "time", Value: bson.M{
			"$gte": primitive.NewDateTimeFromTime(t)}})
	}

	result, err := Collection.Find(ctx, filter)

	if err != nil {
		return searchData, err
	}

	for result.Next(ctx) {
		var data mail.Mail
		err := result.Decode(&data)
		if err != nil {
			return searchData, err
		}
		searchData = append(searchData, &data)
	}

	if searchData == nil {
		return searchData, errors.New("No mail found for the given search criteria!")
	}

	return searchData, err
}
