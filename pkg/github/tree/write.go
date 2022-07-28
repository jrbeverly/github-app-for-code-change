package tree

import (
	"context"
	"time"

	"github.com/google/go-github/github"
	"github.com/jrbeverly/github-app-for-code-change/internal/storage"
	v1 "github.com/jrbeverly/github-app-for-code-change/pkg/config/v1"
)

const DEFAULT_FILE_MODE = "100644"
const FILE_TYPE_BLOB = "blob"

func GetLatestCommit(client *github.Client, ctx context.Context, config *v1.GitHubConfiguration) (ref *github.Reference, err error) {
	var branchRef *github.Reference
	if branchRef, _, err = client.Git.GetRef(ctx, config.Repository.Org, config.Repository.Repo, config.Repository.Branch); err == nil {
		return branchRef, nil
	}

	commitRef := &github.Reference{
		Ref: github.String(config.Repository.Branch),
		Object: &github.GitObject{
			SHA: branchRef.Object.SHA,
		},
	}
	ref, _, err = client.Git.CreateRef(ctx, config.Repository.Org, config.Repository.Repo, commitRef)
	return ref, err
}

func NewCommitContentFromFiles(client *github.Client, ctx context.Context, config *v1.GitHubConfiguration, ref *github.Reference, files *storage.GeneratedFiles) (tree *github.Tree, err error) {
	entries := []github.TreeEntry{}

	for _, fileArg := range files.Files {
		entry := github.TreeEntry{
			Path:    github.String(fileArg.Path),
			Type:    github.String(FILE_TYPE_BLOB),
			Content: github.String(string(fileArg.Contents)),
			Mode:    github.String(DEFAULT_FILE_MODE),
		}
		entries = append(entries, entry)
	}

	tree, _, err = client.Git.CreateTree(ctx, config.Repository.Org, config.Repository.Repo, *ref.Object.SHA, entries)
	return tree, err
}

func newAuthor(config *v1.GitHubConfiguration) *github.CommitAuthor {
	date := time.Now()
	author_name := config.Author.Name
	author_email := config.Author.Email
	return &github.CommitAuthor{Date: &date, Name: &author_name, Email: &author_email}
}

func PushCommit(client *github.Client, ctx context.Context, config *v1.GitHubConfiguration, ref *github.Reference, tree *github.Tree) (err error) {
	parent, _, err := client.Repositories.GetCommit(ctx, config.Repository.Org, config.Repository.Repo, *ref.Object.SHA)
	if err != nil {
		return err
	}

	parent.Commit.SHA = parent.SHA

	author := newAuthor(config)

	commit_msg := config.Author.Commit
	commit := &github.Commit{Author: author, Message: &commit_msg, Tree: tree, Parents: []github.Commit{*parent.Commit}}
	newCommit, _, err := client.Git.CreateCommit(ctx, config.Repository.Org, config.Repository.Repo, commit)
	if err != nil {
		return err
	}

	ref.Object.SHA = newCommit.SHA
	_, _, err = client.Git.UpdateRef(ctx, config.Repository.Org, config.Repository.Repo, ref, false)
	return err
}
