package goclitools

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	uuid "github.com/gofrs/uuid"
	"github.com/skratchdot/open-golang/open"

	"golang.org/x/oauth2"
)

// OAuth2GetToken Create local http server and use it to get access token from oauth authorization provider
func OAuth2GetToken(conf *oauth2.Config) (*oauth2.Token, error) {
	conf.RedirectURL = "http://localhost:3000"
	ctx := context.Background()
	state := uuid.Must(uuid.NewV4()).String()
	url := conf.AuthCodeURL(state, oauth2.AccessTypeOffline)

	serverChannel := make(chan string)
	srv := startHTTPServer(serverChannel)

	timeCountdown(2, "You will be redirected to OAuth provider in: ")
	if err := open.Run(url); err != nil {
		return nil, err
	}

	authError := <-serverChannel
	code := <-serverChannel
	stateInResponse := <-serverChannel

	if authError != "" {
		return nil, fmt.Errorf("Failed to obtain authorization code, reason: %s", authError)
	}

	if state != stateInResponse {
		return nil, errors.New("invalid state from response")
	}

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	fmt.Println("token received, shutting down server...")
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	srv.Shutdown(ctx)

	return tok, nil
}

func timeCountdown(seconds int, message string) {
	for i := 0; i <= seconds; i++ {
		fmt.Printf("\r%s%d", message, seconds-i)
		time.Sleep(time.Second)
	}
	fmt.Print("\n")
}

func startHTTPServer(c chan string) *http.Server {
	srv := &http.Server{Addr: ":3000"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c <- r.URL.Query().Get("error")
		c <- r.URL.Query().Get("code")
		c <- r.URL.Query().Get("state")
		io.WriteString(w, "Your information has been processed, you can close the window and return to script now.\n")
	})

	go func() {
		srv.ListenAndServe()
	}()

	return srv
}
