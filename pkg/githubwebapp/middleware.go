package githubwebapp

import (
	"log"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
	"github.com/jrbeverly/github-app-for-code-change/internal/storage"
	"github.com/justinas/alice"
)

type GithubMiddleware struct {
	WebEvent      interface{}
	Payload       interface{}
	AppIdentifier int64
	PrivateKey    string
	WebhookSecret string
}

func NewGithubMiddleware(id int64, key string, webhook string) *GithubMiddleware {
	middleware := GithubMiddleware{
		AppIdentifier: id,
		PrivateKey:    key,
		WebhookSecret: webhook,
	}
	return &middleware
}

func (mid *GithubMiddleware) NewGithubInterface() alice.Chain {
	return alice.New(mid.validatePayload, mid.authenticate)
}

func (mid *GithubMiddleware) validatePayload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, err := github.ValidatePayload(r, []byte(mid.WebhookSecret))
		if err != nil {
			log.Printf("could not validate payload: err=%s\n", err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		defer r.Body.Close()

		event, err := github.ParseWebHook(github.WebHookType(r), p)
		if err != nil {
			log.Printf("could not parse webhook: err=%s\n", err)
			return
		}

		mid.WebEvent = event
		next.ServeHTTP(w, r)
	})
}

func (mid *GithubMiddleware) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		event := mid.WebEvent
		switch e := event.(type) {
		case *github.IssueCommentEvent:
			result, err := mid.processIssueCommentEvent(e)
			if err != nil {
				return
			}
			mid.Payload = result
		case *github.PushEvent:
			result, ok, err := mid.processPushEvent(e)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			if !ok {
				http.Error(w, http.StatusText(http.StatusAccepted), http.StatusAccepted)
				return
			}
			mid.Payload = result
		}
		next.ServeHTTP(w, r)
	})
}

func (mid *GithubMiddleware) processIssueCommentEvent(event *github.IssueCommentEvent) (storage.TestTriggerEvent, error) {
	return storage.TestTriggerEvent{Key: *event.Comment.Body}, nil
}

func (mid *GithubMiddleware) processPushEvent(event *github.PushEvent) (storage.ConfigChangeEvent, bool, error) {
	var result storage.ConfigChangeEvent
	return result, false, nil
}

func (mid *GithubMiddleware) NewGitHubInstallation(installationId int) (*github.Client, error) {
	transport := http.DefaultTransport
	itr, err := ghinstallation.New(
		transport,
		mid.AppIdentifier,
		int64(installationId),
		[]byte(mid.PrivateKey),
	)

	if err != nil {
		return nil, err
	}

	client := github.NewClient(&http.Client{Transport: itr})
	return client, nil
}
