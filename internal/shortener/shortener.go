package shortener

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"github.com/dbut2/shortener/pkg/url"
	"math/rand"
)

type Shortener struct {
	client *datastore.Client
}

func NewShortener() (*Shortener, error) {
	client, err := datastore.NewClient(context.Background(), "dbut-butla")
	if err != nil {
		return nil, err
	}

	return &Shortener{
		client: client,
	}, nil
}

func (s *Shortener) Shorten(u url.URL, code string) (string, error) {
	if code == "" {
		code = s.randomCode()
	}
	key := datastore.NameKey("urlcode", code, nil)
	blank := url.URL{}
	err := s.client.Get(context.Background(), key, &blank)
	if err != datastore.ErrNoSuchEntity {
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("code taken")
	}
	_, err = s.client.Put(context.Background(), key, &u)
	if err != nil {
		return "", err
	}
	return code, nil
}

func (s *Shortener) Lengthen(code string) (url.URL, error) {
	u := url.URL{}
	key := datastore.NameKey("urlcode", code, nil)
	err := s.client.Get(context.Background(), key, &u)
	if err != nil {
		return url.URL{}, err
	}
	return u, nil
}

func (s *Shortener) randomCode() string {
	u := url.URL{}
	charset := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var code string
	var err error
	for err != datastore.ErrNoSuchEntity {
		c := make([]rune, 4)
		for i := range c {
			c[i] = charset[rand.Intn(len(charset))]
		}
		code = string(c)
		key := datastore.NameKey("urlcode", code, nil)
		err = s.client.Get(context.Background(), key, &u)
	}
	return code
}
