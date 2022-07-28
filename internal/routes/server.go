package routes

import (
	"context"
	"log"
	"net/http"

	v1 "github.com/jrbeverly/github-app-for-code-change/pkg/config/v1"
	"github.com/jrbeverly/github-app-for-code-change/pkg/githubwebapp"
)

var (
	middleware *githubwebapp.GithubMiddleware
	ctx        = context.Background()
	config     *v1.GitHubConfiguration
)

func Start(cfg *v1.GitHubConfiguration) {
	middleware = githubwebapp.NewGithubMiddleware(cfg.Authentication.AppIdentifier, cfg.Authentication.PrivateKey, cfg.Authentication.WebhookSecret)
	stdChain := middleware.NewGithubInterface()
	config = cfg

	http.HandleFunc("/health", health)
	http.HandleFunc("/commit", commit)
	http.Handle("/event_handler", stdChain.Then(http.HandlerFunc(eventHandler)))

	log.Println("Server listening...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
