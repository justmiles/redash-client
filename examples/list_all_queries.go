package main

import "github.com/justmiles/redash"
import "fmt"
import "os"

// set REDASH_URL and REDASH_API_KEY

func main() {
	client, err := redash.NewClient(os.Getenv("REDASH_URL"), os.Getenv("REDASH_API_KEY"))
	if err != nil {
		fmt.Println(err)
		return
	}
	// client.DebugEnabled = true

	res, err := client.ListQueries()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, query := range res.Results {
		fmt.Printf("%s  -  %s\n", query.Name, query.Description)
	}
}
