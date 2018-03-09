package db

import (
	"github.com/blevesearch/bleve"
)

// AllImages returns search result matching all images
func AllImages(index bleve.Index) (*bleve.SearchResult, error) {
	query := bleve.NewMatchAllQuery()
	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Fields = []string{"*"}
	totalImages, err := TotalNumberOfImages(index)
	if err != nil {
		return nil, err
	}
	searchRequest.Size = totalImages
	return index.Search(searchRequest)
}

// TotalNumberOfImages queries the index to count the
// total number of images
func TotalNumberOfImages(index bleve.Index) (int, error) {
	query := bleve.NewMatchAllQuery()
	sizeRequest := bleve.NewSearchRequest(query)
	sizeRequest.Size = 0
	results, err := index.Search(sizeRequest)
	if err != nil {
		return 0, err
	}
	return int(results.Total), nil
}
