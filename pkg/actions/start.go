package actions

import (
	"github.com/jrbeverly/github-app-for-code-change/internal/routes"
	v1 "github.com/jrbeverly/github-app-for-code-change/pkg/config/v1"
)

func Start(cfg *v1.GitHubConfiguration) {
	routes.Start(cfg)
}
