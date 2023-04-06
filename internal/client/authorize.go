package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/dezer32/maze-client/internal/core/config"
)

func (c *Client) Authorize() {
	reqBody, err := json.Marshal(&struct {
		Login    string `json:"login,omitempty"`
		Password string `json:"password,omitempty"`
	}{
		Login:    c.Username,
		Password: c.Password,
	})
	if err != nil {
		logrus.WithError(err).Error("Can't marshal json.")
	}

	reqUrl := c.Url.JoinPath("/maze-api/login/")
	req, err := http.NewRequest("POST", reqUrl.String(), bytes.NewReader(reqBody))
	if err != nil {
		logrus.WithError(err).Error("Can't make request.")
	}

	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	resp, err := c.Do(req)
	if err != nil {
		logrus.WithError(err).Error("Can't authorize.")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err).Error("Can't read body.")
	}

	c.Authorization = &config.Authorization{}
	err = json.Unmarshal(body, c.Authorization)
	if err != nil {
		logrus.WithError(err).Error()
	}

	logrus.WithFields(logrus.Fields{
		"token": fmt.Sprintf(
			"%s*******%s",
			string(c.Authorization.Token[0]),
			string(c.Authorization.Token[len(c.Authorization.Token)-1]),
		),
		"user": c.Authorization.UserId,
	}).Info("Authorized.")
}
