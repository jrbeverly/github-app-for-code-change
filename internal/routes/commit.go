package routes

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/jrbeverly/github-app-for-code-change/internal/storage"
	"github.com/jrbeverly/github-app-for-code-change/pkg/github/tree"
)

func yieldTemplatedFiles() *storage.GeneratedFiles {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return &storage.GeneratedFiles{
		Files: []*storage.GeneratedFile{
			{Path: "a/b/c.txt", Contents: []byte(strconv.Itoa(r1.Intn(100)))},
			{Path: "b/c.txt", Contents: []byte(strconv.Itoa(r1.Intn(100)))},
			{Path: "b/t/c.txt", Contents: []byte(strconv.Itoa(r1.Intn(100)))},
		},
	}
}

func commit(w http.ResponseWriter, r *http.Request) {
	client, err := middleware.NewGitHubInstallation(config.Authentication.InstallationKey)
	if err != nil {
		log.Fatal(err)
	}

	ref, err := tree.GetLatestCommit(client, ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	files := yieldTemplatedFiles()

	commitContent, err := tree.NewCommitContentFromFiles(client, ctx, config, ref, files)
	if err != nil {
		log.Fatalf("Unable to create the tree based on the provided files: %s\n", err)
	}

	if err := tree.PushCommit(client, ctx, config, ref, commitContent); err != nil {
		log.Fatalf("Unable to create the commit: %s\n", err)
	}
}
