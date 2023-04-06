package container

import "github.com/dezer32/maze-client/internal/client"

var c *client.Client

func Client() *client.Client {
	if c == nil {
		c = client.NewClient(Config())
	}

	return c
}
