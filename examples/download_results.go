package main

import (
	"fmt"
	"os"

	"github.com/justmiles/redash-client"
)

// set REDASH_URL and REDASH_API_KEY

func main() {
	client, err := redash.NewClient(os.Getenv("REDASH_URL"), os.Getenv("REDASH_API_KEY"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Search for a query
	queries, err := client.SearchQueries("Traffic to HBase", false, 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Download the latest results for matching querues
	for _, query := range queries {

		// Download a query as xlsx
		err := client.DownloadResults(query, query.Name+".xlsx", "xlsx")
		if err != nil {
			fmt.Println(err)
			return
		}

		// Download a query as csv
		err = client.DownloadResults(query, query.Name+".csv", "csv")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
