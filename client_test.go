package redash

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	client, err := NewClient(os.Getenv("REDASH_URL"), os.Getenv("REDASH_API_KEY"))
	// client.DebugEnabled = true
	assert.Nil(t, err)
	assert.NotNil(t, client)
	_, err = NewClient("http://im a bad url", "")
	assert.NotNil(t, err)

	p, err := client.ListQueries()
	assert.Nil(t, err)
	for _, query := range p.Results {
		fmt.Println(query.Name)
	}
}
