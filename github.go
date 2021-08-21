package main

import (
	"net/url"
	"time"
)

type GithubQuery struct {
	Repository struct {
		ID          ID
		Name        String
		Description String
		URL         URI

		Releases struct {
			Edges []struct {
				Node struct {
					ID          ID
					Name        String
					Description String
					URL         URI
					PublishedAt DateTime
				}
			}
		} `graphql:"releases(last: 1)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

func (query GithubQuery) GetID() ID {
	return query.Repository.ID
}

func (query GithubQuery) GetName() String {
	return query.Repository.Name
}

func (query GithubQuery) GetDescription() String {
	return query.Repository.Description
}

func (query GithubQuery) GetURL() *url.URL {
	return query.Repository.URL.URL
}

func (query GithubQuery) GetReleasesCount() int {
	return len(query.Repository.Releases.Edges)
}

func (query GithubQuery) GetLatestReleaseID() ID {
	return query.Repository.Releases.Edges[0].Node.ID
}

func (query GithubQuery) GetLatestReleaseName() String {
	return query.Repository.Releases.Edges[0].Node.Name
}

func (query GithubQuery) GetLatestReleaseDescription() String {
	return query.Repository.Releases.Edges[0].Node.Description
}

func (query GithubQuery) GetLatestReleaseURL() *url.URL {
	return query.Repository.Releases.Edges[0].Node.URL.URL
}

func (query GithubQuery) GetLatestReleasePublishingDate() time.Time {
	return query.Repository.Releases.Edges[0].Node.PublishedAt.Time
}
