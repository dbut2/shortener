package client

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/dbut2/shortener/pkg/model"
)

type Client struct {
	Host string

}

func New(host string) *Client {
	return &Client{
		Host: host,
	}
}

func (c *Client) Shorten(url string) (string, error) {
	code := randomCode()

	data, _ := json.Marshal(model.Shorten{
		Url:  url,
		Code: code,
	})

	buf := &bytes.Buffer{}
	buf.Write(data)

	resp, err := http.Post(c.Host+"/api/shorten", "application/json", buf)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http not ok: %d", resp.StatusCode)
	}
	return fmt.Sprintf("%s/%s", c.Host, code), nil
}

func randomCode() string {
	chars := "abcdefghijklmnopqrstuvwxyz1234567890"
	str := ""
	for i := 0; i < 4; i++ {
		i, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			panic(err.Error())
		}
		str += string(chars[int(i.Int64())])
	}
	return str
}
