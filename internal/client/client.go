package client

import (
	"net/http"
	"net/url"
	"time"

	"github.com/dezer32/maze-client/internal/core/config"
	"github.com/dezer32/maze-client/internal/core/logger"
)

type Client struct {
	*http.Client
	Url           *url.URL
	Username      string
	Password      string
	Headers       map[string]string
	Authorization *config.Authorization
}

func NewClient(cfg *config.Config) *Client {
	u, err := url.Parse(cfg.Url)
	if err != nil {
		logger.Log.WithError(err).Fatal("Can't parse url.")
	}

	return &Client{
		Client: &http.Client{
			Timeout: 15 * time.Second,
		},
		Headers: map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/json",
		},
		Url:           u,
		Username:      cfg.Username,
		Password:      cfg.Password,
		Authorization: cfg.Authorization,
	}
}
