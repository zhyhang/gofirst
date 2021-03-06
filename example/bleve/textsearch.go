package main

import (
	"fmt"
	"github.com/blevesearch/bleve"
)

const storepath = "/tmp/bleve.example"

func main() {
	message := struct {
		Id   string
		From string
		Body string
	}{
		Id:   "example",
		From: "marty.schoch@gmail.com",
		Body: "bleve indexing is easy",
	}

	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(storepath, mapping)
	if err != nil {
		index, err = bleve.Open(storepath)
		if err != nil {
			panic(err)
		}

	}
	// index some data
	index.Index(message.Id, message)

	// search for some text
	query := bleve.NewQueryStringQuery("bleve")
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, _ := index.Search(searchRequest)

	fmt.Println(searchResult)
}
