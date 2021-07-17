package cli

import (
	"errors"

	"github.com/dbut2/shortener/pkg/client"
	"github.com/spf13/cobra"
)

type Cli struct {
	client *client.Client
	*cobra.Command
}

func New(host string) *Cli {
	return &Cli{
		client: client.New(host),
	}
}

func (c *Cli) Shorten() *cobra.Command {
	cmd := &cobra.Command{
		Use: "shorten",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("supply a url")
			}
			url := args[0]

			short, err := c.client.Shorten(url)
			if err != nil {
				return err
			}
			cmd.Printf("short url: %s\n", short)
			return nil
		},
	}

	return cmd
}
