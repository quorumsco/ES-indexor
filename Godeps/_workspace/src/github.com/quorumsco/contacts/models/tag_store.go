// Definition of the structures and SQL interaction functions
package models

import "github.com/jinzhu/gorm"

// TagDS implements the TagSQL methods
type TagDS interface {
	Save(*Tag, TagArgs) error
	Delete(*Tag, TagArgs) error
	Find(TagArgs) ([]Tag, error)
}

// Tagstore returns a TagDS implementing CRUD methods for the tags and containing a gorm client
func TagStore(db *gorm.DB) TagDS {
	return &TagSQL{DB: db}
}
