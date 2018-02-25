package db

import (
	"path/filepath"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
	homedir "github.com/mitchellh/go-homedir"
)

func Init() (bleve.Index, error) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, ".pd.db")
	index, err := bleve.Open(dbPath)
	if err != nil {
		indexMapping := buildIndexMapping()
		return bleve.New(dbPath, indexMapping)
	}
	return index, err
}

func buildIndexMapping() mapping.IndexMapping {
	enFieldMapping := bleve.NewTextFieldMapping()
	enFieldMapping.Analyzer = "en"

	imageMapping := bleve.NewDocumentMapping()
	imageMapping.AddFieldMappingsAt("Description", enFieldMapping)

	indexMapping := bleve.NewIndexMapping()
	indexMapping.DefaultMapping = imageMapping
	return indexMapping
}
