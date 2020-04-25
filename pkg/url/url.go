package url

import "github.com/dbut2/shortener/pkg/shortener"

type URL struct {
	Protocol string
	Hostname string
	Port int
	Path string
	Queries []*Query
}

func (u *URL) String() string {
	return ""
}

func (u *URL) Shorten(code string) URL {
	return shortener.Shorten(u, code)
}

func URLFromString(s string) URL {

}
