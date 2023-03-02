package service

import (
	"context"
	ghHooks "github.com/go-playground/webhooks/v6/github"
	"net/http"

	"github.com/deed-labs/gittips/bot/internal/entity"
	"github.com/deed-labs/gittips/bot/internal/repository"
	"github.com/deed-labs/gittips/bot/internal/ton"
)

type Owners interface {
	Exists(ctx context.Context, id int64) (bool, error)
	Create(ctx context.Context, id int64, login string, url string, avatarURL string, ownerType string) error
}

type Bounties interface {
	GetAll(ctx context.Context) ([]*entity.Bounty, error)
	Create(ctx context.Context, id int64, ownerID int64, title string, url string, body string) error
	Delete(ctx context.Context, id int64) error
}

type Comments interface {
	Process(ctx context.Context, senderId int64, ownerId int64, body string) error
}

type Github interface {
	ProcessOrganizationInstallation(ctx context.Context, payload ghHooks.InstallationPayload) error
	ProcessRepositoriesInstallation(ctx context.Context, payload ghHooks.InstallationRepositoriesPayload) error
	ProcessIssueEvent(ctx context.Context, payload ghHooks.IssuesPayload) error
	ProcessIssueComment(ctx context.Context, payload ghHooks.IssueCommentPayload) error
	ProcessNewPR(ctx context.Context, payload ghHooks.PullRequestPayload) error
	ProcessPRComment(ctx context.Context, payload ghHooks.PullRequestReviewCommentPayload) error
}

type Services struct {
	Owners   Owners
	Bounties Bounties
	Comments Comments
	Github   Github
}

type Deps struct {
	TON          *ton.TON
	GithubClient *http.Client
	Repository   repository.Repository
}

func New(deps *Deps) *Services {
	ownersSvc := NewOwnersService(deps.Repository)
	bountiesSvc := NewBountiesService(ownersSvc, deps.Repository)
	commentsSvc := NewCommentsService(deps.TON)
	githubSvc := NewGithubService(deps.GithubClient, ownersSvc, bountiesSvc)

	return &Services{
		Owners:   ownersSvc,
		Bounties: bountiesSvc,
		Comments: commentsSvc,
		Github:   githubSvc,
	}
}
