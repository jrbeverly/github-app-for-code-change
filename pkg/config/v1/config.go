package v1

type GitHubConfiguration struct {
	Version        string               `yaml:"version"`
	Authentication GitHubAuthentication `yaml:"auth"`
	Author         GitHubAuthorship     `yaml:"git"`
	Repository     GitHubRepository     `yaml:"repo"`
}

type GitHubAuthentication struct {
	InstallationKey int    `yaml:"installation"`
	AppIdentifier   int64  `yaml:"app-identifier"`
	PrivateKey      string `yaml:"private-key"`
	WebhookSecret   string `yaml:"webhook-secret"`
}

type GitHubRepository struct {
	Org    string `yaml:"org"`
	Repo   string `yaml:"repo"`
	Branch string `yaml:"branch"`
}

type GitHubAuthorship struct {
	Name   string `yaml:"name"`
	Email  string `yaml:"email"`
	Commit string `yaml:"commit"`
}
