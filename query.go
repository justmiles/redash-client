package redash

import "fmt"
import "errors"
import "strconv"

// Query is returned for all /api/queries
type Query struct {
	DataSourceID      int     `json:"data_source_id"`
	LastModifiedByID  int     `json:"last_modified_by_id"`
	LatestQueryDataID int     `json:"latest_query_data_id"`
	Schedule          string  `json:"schedule"`
	IsArchived        bool    `json:"is_archived"`
	RetrievedAt       string  `json:"retrieved_at"`
	UpdatedAt         string  `json:"updated_at"`
	User              User    `json:"user"`
	Query             string  `json:"query"`
	IsDraft           bool    `json:"is_draft"`
	ID                int     `json:"id"`
	Description       string  `json:"description"`
	Runtime           float64 `json:"runtime"`
	Name              string  `json:"name"`
	CreatedAt         string  `json:"created_at"`
	Version           int     `json:"version"`
	QueryHash         string  `json:"query_hash"`
	APIKey            string  `json:"api_key"`
	Options           struct {
		Parameters []struct {
			Global bool   `json:"global"`
			Type   string `json:"type"`
			Name   string `json:"name"`
			Value  string `json:"value"`
			Title  string `json:"title"`
		} `json:"parameters"`
	} `json:"options"`
}

type ListQueriesResponse struct {
	Count    int     `json:"count"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
	Results  []Query `json:"results"`
}

func (c *Client) ListQueries() (p ListQueriesResponse, err error) {

	req, err := c.newRequest("GET", "/api/queries", nil, nil)
	if err != nil {
		return p, err
	}
	_, err = c.do(req, &p)
	if err != nil {
		return p, err
	}
	return p, nil
}

// SearchQueries searches Redash for queries
// q - query to search Form
// drafts - boolean whether or not to include drafts
func (c *Client) SearchQueries(q string, drafts bool) (p []Query, err error) {

	req, err := c.newRequest("GET", "/api/queries/search", nil, &map[string]string{
		"q":              q,
		"include_drafts": strconv.FormatBool(drafts),
	})

	if err != nil {
		return p, err
	}
	_, err = c.do(req, &p)
	if err != nil {
		return p, err
	}
	return p, nil
}

// // DownloadResults will write the latest query results to the filesystem
func (c *Client) DownloadResults(q Query, filelocation, filetype string) (err error) {
	if filetype == "" {
		filetype = "xlsx"
	}
	if filetype != "xlsx" && filetype != "csv" {
		return errors.New(fmt.Sprintf(`Unable to download file of type "%s". Please specify "xlsx" or "csv"`, filetype))
	}
	req, err := c.newRequest("GET", fmt.Sprintf("/api/queries/%d/results/%d.%s", q.ID, q.LatestQueryDataID, filetype), nil, nil)
	if err != nil {
		return err
	}
	_, err = c.download(req, filelocation)
	if err != nil {
		return err
	}
	return nil
}
