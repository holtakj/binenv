package list

import (
	"context"
)

// Lister should return a list of available release versions
type Lister interface {
	Get(ctx context.Context) ([]string, error)
}

// List contains list definition
type List struct {
	Type        string `yaml:"type"`
	Prefix      string `yaml:"prefix"`
	VersionFrom string `yaml:"version_from"`
	URL         string `yaml:"url"`
}

// Factory returns instances that comply to Lister interface
func (l List) Factory() Lister {
	switch l.Type {
	case "github-releases":
		return GithubRelease{
			url:         l.URL,
			prefix:      l.Prefix,
			versionFrom: l.VersionFrom,
		}
	}

	return nil
}
