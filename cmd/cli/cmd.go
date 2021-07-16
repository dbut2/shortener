package main

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/dbut2/shortener/pkg/model"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use: "shortener",
	}

	cmd.AddCommand(shorten())

	return cmd
}

func shorten() *cobra.Command {
	cmd := &cobra.Command{
		Use: "shorten",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("supply a url")
			}
			url := args[0]
			code := randomCode()

			fmt.Println(code)

			data, _ := json.Marshal(model.Shorten{
				Url:  url,
				Code: code,
			})

			buf := &bytes.Buffer{}
			buf.Write(data)

			resp, err := http.Post("https://shortener-dot-dbut-0.ts.r.appspot.com/api/shorten", "application/json", buf)
			cmd.Println(resp.StatusCode)
			return err
		},
	}

	return cmd
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
