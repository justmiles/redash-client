package redash

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	client, err := NewClient(os.Getenv("REDASH_URL"), os.Getenv("REDASH_API_KEY"))
	client.DebugEnabled = true
	assert.Nil(t, err)
	assert.NotNil(t, client)

	_, err = NewClient("http://im a bad url", "")
	assert.NotNil(t, err)

	p, err := client.ListQueriesWithPagination(nil)
	assert.Nil(t, err)
	assert.NotNil(t, p)

	query := Query{
		Name:         "Testing from Golang",
		Query:        "SELECT now();",
		Description:  "This is a test",
		DataSourceID: 1,
	}

	query, err = client.CreateQuery(query)
	assert.Nil(t, err)
	fmt.Println(query)

	err = client.DeleteQuery(query)
	assert.Nil(t, err)
}
