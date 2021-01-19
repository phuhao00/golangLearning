package main

import (
	"fmt"
	"github.com/blevesearch/bleve/v2"
)

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
	index, err := bleve.New("example.bleve", mapping)
	if err != nil {
		panic(err)
	}
	// index some data
	err = index.Index(message.Id, message)

	// search for some text
	query := bleve.NewMatchQuery("example")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)

	fmt.Println(searchResults, err)
}
