// Definition of the structures and SQL interaction functions
package models

import "github.com/jinzhu/gorm"

// NoteDS implements the NoteSQL methods
type NoteDS interface {
	Save(*Note, NoteArgs) error
	Delete(*Note, NoteArgs) error
	First(NoteArgs) (*Note, error)
	Find(NoteArgs) ([]Note, error)
}

// Notestore returns a NoteDS implementing CRUD methods for the notes and containing a gorm client
func NoteStore(db *gorm.DB) NoteDS {
	return &NoteSQL{DB: db}
}
