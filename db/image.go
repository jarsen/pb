package db

import (
	"encoding/json"
	"log"
	"net/url"
	"time"

	"github.com/blevesearch/bleve"
)

type Image struct {
	ID          string
	Url         *url.URL
	Description string
	Date        time.Time
}

// Index adds the image to the bleve index and to the
func (image *Image) AddTo(index bleve.Index) error {
	buf, jsonErr := json.Marshal(image)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	if err := index.Index(image.ID, image); err != nil {
		return err
	}
	return index.SetInternal([]byte(image.ID), buf)
}
