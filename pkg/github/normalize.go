package github

import (
	"net/url"
	"path"
	"strings"
)

const (
	githubRawContentHost = "raw.githubusercontent.com"
	githubHost           = "github.com"
)

func NormalizeURLToContent(u *url.URL, extensions ...string) *url.URL {
	if u.Host != githubHost {
		return u
	}
	if len(extensions) > 0 && !extensionMatches(u, extensions) {
		return u
	}

	p := u.Path
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	parts := strings.Split(p, "/")
	if len(parts) < 4 {
		return u
	}

	normalized := *u
	normalized.Host = githubRawContentHost
	normalized.Path = "/" + path.Join(append(parts[:3], parts[4:]...)...)
	return &normalized
}

func extensionMatches(u *url.URL, extensions []string) bool {
	ext := path.Ext(u.Path)
	for _, e := range extensions {
		if strings.EqualFold(ext, "."+e) {
			return true
		}
	}
	return false
}
