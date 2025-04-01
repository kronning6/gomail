# gomail

gomail is a CLI for Gmail enabling email screening power tools.

## Overview

This project serves two purposes:

1. Create a TUI program to interact with Gmail to provide email screening, making it easier
   to manage an unruly inbox
1. Get more familiar with Go

The future of this project may expand into using AI to help with email screening. It could also
expand into an email screening service, but for now it's focused on being a power tool that needs
to run in a CLI.

## Getting Started

This project was started from: https://developers.google.com/gmail/api/quickstart/go.

To get started, follow the instructions to enable the Gmail API on your account and populate
`credentials.json` from the `credentials-eample.json` file.

To authenticate run `go run main.goi setup`.

## Roadmap

- [x] Get all emails received today
- [ ] Display emails using Charm
- [ ] Select email and load all emails from sender and be able to review
- [ ] Mass delete selected emails by sender
- [ ] Find all emails for a set of senders
- [ ] Keep track of senders and screen (isolate) new senders with Yes or No
- [ ] Store and manage sender status in Turso DB
- [ ] See if any emails from the sender are important
- [ ] See if any screened out emails are important
- [ ] Unsubscribe from emails
- [ ] AI

## Core Problems

- Emails end up in Promotions that shouldn't
- Important emails end up in Promotions making it harder to mass delete

## Inspiration

- https://superhuman.com
- https://www.shortwave.com
- https://clickup.com/p/ai-agents/gmail
- https://www.fyxer.com/
