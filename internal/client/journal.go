package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Journal struct {
	Id          int    `json:"id"`
	User        int    `json:"user"`
	Ticket      int    `json:"ticket"`
	TicketTitle string `json:"ticketTitle"`
	TicketScope string `json:"ticketScope"`
	Date        string `json:"date"`
	Minutes     int    `json:"minutes"`
	Comment     string `json:"comment"`
}

func (c *Client) Journal(from time.Time, to time.Time) []*Journal {

	reqUrl := c.Url.JoinPath("/maze-api/journals/")
	q := reqUrl.Query()
	q.Add("from", from.Format("2006-01-02"))
	q.Add("to", to.Format("2006-01-02"))
	reqUrl.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		logrus.WithError(err).Error("Can't create request for journal.")
		return nil
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Authorization.Token))

	resp, err := c.Do(req)
	if err != nil {
		logrus.WithError(err).Error("Can't get journal.")
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err).Error("Can't read body.")
		return nil
	}

	var journals []*Journal
	err = json.Unmarshal(body, &journals)
	if err != nil {
		logrus.WithError(err).Error("Can't convert journal body to var.")
		return nil
	}

	logrus.WithField("lines", len(journals)).Info("Loaded journals.")

	return journals
}
