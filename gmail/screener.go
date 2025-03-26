package gmail

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func Screener(show bool) {
	if !show {
		return
	}
	ctx := context.Background()
	b, err := os.ReadFile("./gmail/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client, err := getClient(config)
	if err != nil {
		log.Fatalf("Could not get client. Please run:\n\n`gomail setup`\n")
	}

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	today := time.Now().Format("2006/01/02")
	query := fmt.Sprintf("after:%s before:%s", today, time.Now().AddDate(0, 0, 1).Format("2006/01/02"))

	// List messages from today
	user := "me" // "me" refers to the authenticated user
	r, err := srv.Users.Messages.List(user).Q(query).MaxResults(100).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve messages: %v", err)
	}

	fmt.Printf("Found %d messages from today\n\n", len(r.Messages))
	for _, msg := range r.Messages {
		fullMsg, err := srv.Users.Messages.Get(user, msg.Id).Do()
		if err != nil {
			log.Printf("Unable to retrieve message %s: %v", msg.Id, err)
			continue
		}

		subject := ""
		sender := ""

		for _, header := range fullMsg.Payload.Headers {
			if header.Name == "Subject" {
				subject = header.Value
			} else if header.Name == "From" {
				sender = header.Value
			}

			if subject != "" && sender != "" {
				break
			}
		}

		internalDate := time.Unix(fullMsg.InternalDate/1000, 0)
		formattedDate := internalDate.Format("2006-01-02 15:04:05")

		fmt.Printf("Date: %s, From: %s, Subject: %s\n", formattedDate, sender, subject)
	}
}

func getClient(config *oauth2.Config) (*http.Client, error) {
	tokFile := "./gmail/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		return nil, err
	}
	return config.Client(context.Background(), tok), nil
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}
