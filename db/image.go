package db

import (
	"time"

	"github.com/blevesearch/bleve"
)

type Image struct {
	ID          string
	URL         string
	Description string
	Date        time.Time
}

// Index adds the image to the bleve index and to the
func (image *Image) AddTo(index bleve.Index) error {
	return index.Index(image.ID, image)
}
