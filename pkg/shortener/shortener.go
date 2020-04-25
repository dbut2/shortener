package shortener

import "github.com/dbut2/shortener/pkg/url"

func Shorten(url *url.URL, code string) url.URL {
	return url.Shorten(code)
}
