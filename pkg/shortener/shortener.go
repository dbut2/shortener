package shortener

import (
	"cloud.google.com/go/datastore"
	"context"
	"errors"
)

type Url struct {
	Code string
	Url string
}

func Shorten(url string, code string) error {
	ctx := context.Background()


	client, err := datastore.NewClient(ctx, "dbut-0")
	if err != nil {
		return err
	}

	k := datastore.IncompleteKey("Url", nil)

	e := new(Url)

	e.Code = code
	e.Url = url

	_, err = client.Put(ctx, k, e)
	if err != nil {
		return err
	}
	return nil
}

func Lengthen(code string) (string, error) {
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, "dbut-0")
	if err != nil {
		return "", err
	}

	q := datastore.NewQuery("Url").Filter("Code =", code).Limit(1)

	it := client.Run(ctx, q)

	var url Url
	_, err = it.Next(&url)

	if url.Url == "" {
		return "", errors.New("Code not exist")
	}
	return url.Url, nil
}

func CodeExists(code string) bool {
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, "dbut-0")
	if err != nil {
		panic(err.Error())
	}

	q := datastore.NewQuery("Url").Filter("Code =", code).Limit(1)

	it := client.Run(ctx, q)

	var url Url
	_, err = it.Next(&url)

	if url.Url == "" {
		return false
	}
	return true
}
