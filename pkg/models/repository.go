package models

import (
	"net/url"
)

// Repository on GitHub or Gitlab.
type Repository struct {
	ID          string
	Name        string
	Owner       string
	Description string
	URL         url.URL
	Release     Release
}
