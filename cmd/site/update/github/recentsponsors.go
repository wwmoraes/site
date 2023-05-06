package github

import (
	"context"
	"log"

	github "github.com/shurcooL/githubv4"
)

type recentSponsorsQuery struct {
	User struct {
		// Login                    github.String
		SponsorshipsAsMaintainer struct {
			// TotalCount github.Int
			Edges []struct {
				Cursor github.String
				Node   struct {
					CreatedAt     github.DateTime
					SponsorEntity struct {
						Typename     github.String `graphql:"__typename"`
						User         qlUser        `graphql:"... on User"`
						Organization qlUser        `graphql:"... on Organization"`
					}
				}
			}
		} `graphql:"sponsorshipsAsMaintainer(first: $count, orderBy: {field: CREATED_AT, direction: DESC})"`
	} `graphql:"user(login:$username)"`
}

type User struct {
	Login     string
	Name      string
	AvatarURL string
	URL       string
}

func (handler *Handler) GetRecentSponsors(ctx context.Context, count int) ([]User, error) {
	query := recentSponsorsQuery{}
	vars := Variables{
		"username": github.String(handler.username),
		"count":    github.Int(count),
	}

	err := handler.client.Query(ctx, &query, vars)
	if err != nil {
		return nil, err
	}

	sponsors := make([]User, 0, len(query.User.SponsorshipsAsMaintainer.Edges))

	for _, v := range query.User.SponsorshipsAsMaintainer.Edges {
		switch string(v.Node.SponsorEntity.Typename) {
		case "Organization":
			sponsors = append(sponsors, *v.Node.SponsorEntity.Organization.Unwrap())
		case "User":
			sponsors = append(sponsors, *v.Node.SponsorEntity.User.Unwrap())
		default:
			log.Println("unknown sponsor entity type named", v.Node.SponsorEntity.Typename)
		}
	}

	return sponsors, nil
}
