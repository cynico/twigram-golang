package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"twigram-go/data"
	"twigram-go/handlers"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

// This function construct the URL the user will visit to give/deny permission to the application.
// An example of such url:

// https://twitter.com/i/oauth2/authorize?response_type=code&client_id=$CLIENT_ID&\
// redirect_uri=$REDIRECT_URL&scope=tweet.read%20users.read%20follows.read+follows.write%20\
// offline.access&state=state&code_challenge=$CODE_CHALLENGE&code_challenge_method=plain
func ConstructAutorizationURL() string {
	config := &oauth2.Config{
		ClientID:     data.ClientID,
		ClientSecret: data.ClientSecret,
		Scopes:       []string{"follows.read", "offline.access"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://twitter.com/i/oauth2/authorize",
			TokenURL: "https://api.twitter.com/oauth/request_token",
		},
		RedirectURL: data.RedirectURI,
	}

	url := config.AuthCodeURL("state")
	url += fmt.Sprintf("&code_challenge=%s&code_challenge_method=plain", data.CodeVerifier)

	return url
}

func StartHTTPServer() {

	fmt.Printf("Visit the URL for the auth dialog: %v\n", ConstructAutorizationURL())

	sm := mux.NewRouter()
	callbackRouter := sm.Methods(http.MethodGet).Subrouter()
	callbackRouter.HandleFunc("/callback", handlers.HandleCallBack).Queries("code")
	callbackRouter.HandleFunc("/callback", handlers.HandleCallBack)

	server := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  5 * time.Second,  // max time to read request from the client
		WriteTimeout: 10 * time.Second, // max time to write response to the client
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Println("Starting to listen on port 9090")
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Error starting server on port 9090: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
