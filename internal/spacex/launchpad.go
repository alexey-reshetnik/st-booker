package spacex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (c *Client) LaunchesForDate(launchpadID string, date time.Time) (int, error) {
	requestBody := fmt.Sprintf(`
{
  "query": {
      "launchpad": "%s",
      "$and": [{"date_utc": {"$gte": "%s"}}, {"date_utc": {"$lt": "%s"}}]
  }
}
`, launchpadID, date.Format(time.RFC3339), date.Add(24*time.Hour).Format(time.RFC3339))

	resp, err := http.Post(
		fmt.Sprintf("%s/v4/launches/query", c.apiURL),
		"application/json", bytes.NewBuffer([]byte(requestBody)),
	)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("got unexpected http status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	lqr := launchesQueryResult{}
	err = json.Unmarshal(body, &lqr)
	if err != nil {
		return 0, err
	}

	return lqr.TotalDocs, nil
}

type launchesQueryResult struct {
	TotalDocs int
}
