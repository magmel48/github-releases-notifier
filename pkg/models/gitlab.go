package models

import (
	"net/url"
	"time"
)

type GitlabQuery struct {
	Repository struct {
		ID          ID
		Name        String
		Description String
		URL         URI `graphql:"webUrl"`

		Releases struct {
			Edges []struct {
				Node struct {
					Name        String `graphql:"tagName"`
					Description String
					PublishedAt DateTime `graphql:"releasedAt"`
					Commit      struct {
						Sha String
					}
					Links struct {
						SelfURL URI `graphql:"selfUrl"`
					}
				}
			}
		} `graphql:"releases(sort: RELEASED_AT_DESC, first: 1)"`
	} `graphql:"project(fullPath: $fullPath)"`
}

func (query GitlabQuery) GetID() ID {
	return query.Repository.ID
}

func (query GitlabQuery) GetName() String {
	return query.Repository.Name
}

func (query GitlabQuery) GetDescription() String {
	return query.Repository.Description
}

func (query GitlabQuery) GetURL() *url.URL {
	return query.Repository.URL.URL
}

func (query GitlabQuery) GetReleasesCount() int {
	return len(query.Repository.Releases.Edges)
}

func (query GitlabQuery) GetLatestReleaseID() ID {
	return query.Repository.Releases.Edges[0].Node.Commit.Sha
}

func (query GitlabQuery) GetLatestReleaseName() String {
	return query.Repository.Releases.Edges[0].Node.Name
}

func (query GitlabQuery) GetLatestReleaseDescription() String {
	return query.Repository.Releases.Edges[0].Node.Description
}

func (query GitlabQuery) GetLatestReleaseURL() *url.URL {
	return query.Repository.Releases.Edges[0].Node.Links.SelfURL.URL
}

func (query GitlabQuery) GetLatestReleasePublishingDate() time.Time {
	return query.Repository.Releases.Edges[0].Node.PublishedAt.Time
}
