package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dezer32/maze-client/internal/core/logger"
)

type Journal struct {
	Id          int         `json:"id"`
	User        int         `json:"user"`
	Ticket      int         `json:"ticket"`
	TicketTitle string      `json:"ticketTitle"`
	TicketScope string      `json:"ticketScope"`
	Date        JournalTime `json:"date"`
	Minutes     int         `json:"minutes"`
	Comment     string      `json:"comment"`
}

type JournalTime struct {
	time.Time
}

func (c *JournalTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) // get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value) // parse time
	if err != nil {
		return err
	}
	*c = JournalTime{t} // set result using the pointer
	return nil
}

func (c JournalTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + c.Time.Format("2006-01-02") + `"`), nil
}

func (c *Client) Journal(from time.Time, to time.Time) []*Journal {

	reqUrl := c.Url.JoinPath("/maze-api/journals/")
	q := reqUrl.Query()
	q.Add("from", from.Format("2006-01-02"))
	q.Add("to", to.Format("2006-01-02"))
	reqUrl.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		logger.Log.WithError(err).Error("Can't create request for journal.")
		return nil
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Authorization.Token))

	resp, err := c.Do(req)
	if err != nil {
		logger.Log.WithError(err).Error("Can't get journal.")
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		logger.Log.Fatal(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.WithError(err).Error("Can't read body.")
		return nil
	}

	var journals []*Journal
	err = json.Unmarshal(body, &journals)
	if err != nil {
		logger.Log.WithError(err).Error("Can't convert journal body to var.")
		return nil
	}

	logger.Log.WithField("lines", len(journals)).Info("Loaded journals.")

	return journals
}
