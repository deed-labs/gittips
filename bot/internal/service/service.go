package service

import (
	"context"
	"net/http"

	"github.com/deed-labs/gittips/bot/internal/entity"
	"github.com/deed-labs/gittips/bot/internal/repository"
	"github.com/deed-labs/gittips/bot/internal/ton"
	ghHooks "github.com/go-playground/webhooks/v6/github"
)

type Owners interface {
	Exists(ctx context.Context, id int64) (bool, error)
	Create(ctx context.Context, id int64, login string, url string, avatarURL string, ownerType string) error
	LinkWithWallet(ctx context.Context, ownerId int64, walletAddress string) error
}

type Bounties interface {
	GetAll(ctx context.Context) ([]*entity.Bounty, error)
	Create(ctx context.Context, id int64, ownerID int64, title string, url string, body string) error
	Delete(ctx context.Context, id int64) error
}

type Commands interface {
	Parse(command string) interface{}
}

// GitHub is service for processing github events and interacting with github API.
// TODO: replace third-party types with our own models?
type GitHub interface {
	ProcessOrganizationInstallation(ctx context.Context, payload ghHooks.InstallationPayload) error
	ProcessRepositoriesInstallation(ctx context.Context, payload ghHooks.InstallationRepositoriesPayload) error
	ProcessIssueEvent(ctx context.Context, payload ghHooks.IssuesPayload) error
	ProcessIssueComment(ctx context.Context, payload ghHooks.IssueCommentPayload) error
	ProcessNewPR(ctx context.Context, payload ghHooks.PullRequestPayload) error
	ProcessPRComment(ctx context.Context, payload ghHooks.PullRequestReviewCommentPayload) error
	ProcessInstallationSetup(ctx context.Context, installationId int64, walletAddress string) error
}

type Services struct {
	Owners   Owners
	Bounties Bounties
	Commands Commands
	GitHub   GitHub
}

type Deps struct {
	TON          *ton.TON
	GitHubClient *http.Client
	Repository   repository.Repository
}

func New(deps *Deps) *Services {
	ownersSvc := NewOwnersService(deps.Repository)
	bountiesSvc := NewBountiesService(ownersSvc, deps.Repository)
	commandsSvc := NewCommandsService(deps.TON, deps.Repository)
	githubSvc := NewGitHubService(deps.GitHubClient, ownersSvc, bountiesSvc, commandsSvc)

	return &Services{
		Owners:   ownersSvc,
		Bounties: bountiesSvc,
		Commands: commandsSvc,
		GitHub:   githubSvc,
	}
}
