package db

import (
	"path/filepath"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/lang/en"
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
	indexMapping := bleve.NewIndexMapping()
	imageMapping := bleve.NewDocumentMapping()
	descriptionMapping := bleve.NewTextFieldMapping()
	descriptionMapping.Analyzer = en.AnalyzerName
	imageMapping.AddFieldMappingsAt("description", descriptionMapping)
	indexMapping.AddDocumentMapping("image", imageMapping)
	return indexMapping
}
