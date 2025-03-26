package gmail

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

const port = 9876

func Setup() {
	ctx := context.Background()
	b, err := os.ReadFile("./gmail/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	codeChan := make(chan string)

	server := &http.Server{Addr: fmt.Sprintf(":%d", port)}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code != "" {
			codeChan <- code

			fmt.Fprintf(w, "Authentication successful! You can close this window.")

			go func() {
				server.Shutdown(ctx)
			}()
		} else {
			fmt.Fprintf(w, "Error: no code found in the request")
		}
	})

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	config.RedirectURL = fmt.Sprintf("http://localhost:%d", port)

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Printf("Opening browser for authentication...\n")
	if err := openBrowser(authURL); err != nil {
		fmt.Printf("Failed to open browser: %v\n", err)
		fmt.Printf("Please open the following URL in your browser:\n%v\n", authURL)
	}

	fmt.Println("Waiting for authentication...")
	code := <-codeChan

	token, err := config.Exchange(ctx, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}

	tokenFile := "./gmail/token.json"
	saveToken(tokenFile, token)

	fmt.Printf("Token successfully retrieved and saved to %s\n", tokenFile)
}

func openBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
