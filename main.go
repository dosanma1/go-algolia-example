// hello_algolia.go
package main

import (
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

type Record struct {
	ObjectID string `json:"objectID"`
	Name     string `json:"name"`
}

func main() {
	// Connect and authenticate with your Algolia app
	client := search.NewClient("JE8KG0V2PC", "YourWriteAPIKey")

	// Create a new index and add a record
	index := client.InitIndex("test_index")
	resSave, err := index.SaveObject(Record{ObjectID: "1", Name: "test_record"})
	if err != nil {
		fmt.Println(err.Error())
		err = nil
	}
	resSave.Wait()

	// Search the index and print the results
	results, err := index.Search("test_record")
	if err != nil {
		fmt.Println(err.Error())
		err = nil
	} else {
		var records []Record
		results.UnmarshalHits(&records)
		fmt.Println(records)
	}
}
