package github

import (
	github "github.com/shurcooL/githubv4"
)

type qlGist struct {
	IsPublic    github.Boolean
	Name        github.String
	Description github.String
	URL         github.String
	CreatedAt   github.DateTime
}

func (ql *qlGist) Unwrap() *Gist {
	return &Gist{
		Name:        string(ql.Name),
		Description: string(ql.Description),
		URL:         string(ql.URL),
		CreatedAt:   ql.CreatedAt.Time,
	}
}

type qlPullRequest struct {
	URL        github.String
	Title      github.String
	State      github.PullRequestState
	CreatedAt  github.DateTime
	Repository qlRepository
}

func (ql *qlPullRequest) Unwrap() *PullRequest {
	return &PullRequest{
		URL:        string(ql.URL),
		Title:      string(ql.Title),
		State:      string(ql.State),
		CreatedAt:  ql.CreatedAt.Time,
		Repository: *ql.Repository.Unwrap(),
	}
}

type qlRelease struct {
	Nodes []qlReleaseInfo
}

type qlUser struct {
	Login     github.String
	Name      github.String
	AvatarURL github.String
	URL       github.String
}

func (ql *qlUser) Unwrap() *User {
	return &User{
		Login:     string(ql.Login),
		Name:      string(ql.Name),
		AvatarURL: string(ql.AvatarURL),
		URL:       string(ql.URL),
	}
}

type qlRepository struct {
	NameWithOwner github.String
	URL           github.String
	Description   github.String
	IsPrivate     github.Boolean
	Stargazers    struct {
		TotalCount github.Int
	}
}

func (ql *qlRepository) Unwrap() *Repository {
	return &Repository{
		NameWithOwner: string(ql.NameWithOwner),
		URL:           string(ql.URL),
		Description:   string(ql.Description),
		Stargazers:    int(ql.Stargazers.TotalCount),
	}
}

type qlReleaseInfo struct {
	Name         github.String
	TagName      github.String
	PublishedAt  github.DateTime
	URL          github.String
	IsPrerelease github.Boolean
	IsDraft      github.Boolean
}

func (ql *qlReleaseInfo) Unwrap() *Release {
	return &Release{
		Name:        string(ql.Name),
		TagName:     string(ql.TagName),
		PublishedAt: ql.PublishedAt.Time,
		URL:         string(ql.URL),
	}
}
