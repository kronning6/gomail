# gmail-screener

## Overview

This project serves two purposes:

1. Create a TUI program to interact with Gmail to provide email screening, making it easier
   to manage an unruly inbox
1. Get more familiar with Go

The future of this project may expand into using AI to help with email screening. It could also
expand into an email screening service, but for now it's focused on being a power tool that needs
to run in a CLI.

## Getting Started

This project was started from: https://developers.google.com/gmail/api/quickstart/go

To get started, follow the instructions to enable the Gmail API on your account and populate
`credentials.json` from the `credentials-eample.json` file.

Run `go run gms.go`. The first time you run it you'll have to manually parse out the token and paste it back
into the terminal.

## Roadmap

- Get all emails received today
- Keep track of senders and screen new senders with Yes or No
- Select email and load all emails from sender and be able to review
- Mass delete selected emails by sender
- See if any emails from the sender are important
- See if any screened out emails are important
- Unsubscribe from emails
