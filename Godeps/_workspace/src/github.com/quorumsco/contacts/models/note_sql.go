// Definition of the structures and SQL interaction functions
package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// NoteSQL contains a Gorm client and the note and gorm related methods
type NoteSQL struct {
	DB *gorm.DB
}

// Save inserts a new note into the database
func (s *NoteSQL) Save(n *Note, args NoteArgs) error {
	if n == nil {
		return errors.New("save: note is nil")
	}

	n.GroupID = args.Note.GroupID
	if n.ID == 0 {
		err := s.DB.Create(n).Error
		s.DB.Last(n)
		return err
	}

	return s.DB.Where("group_id = ?", args.Note.GroupID).Save(n).Error
}

// Delete removes a note from the database
func (s *NoteSQL) Delete(n *Note, args NoteArgs) error {
	if n == nil {
		return errors.New("delete: note is nil")
	}

	return s.DB.Where("group_id = ?", args.Note.GroupID).Delete(n).Error
}

// First returns a note from the database usin it's ID
func (s *NoteSQL) First(args NoteArgs) (*Note, error) {
	var n Note

	if err := s.DB.Where(args.Note).First(&n).Error; err != nil {
		if err == gorm.RecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &n, nil
}

// Find returns all the notes containing a given groupID from the database
func (s *NoteSQL) Find(args NoteArgs) ([]Note, error) {
	var notes []Note

	err := s.DB.Where("group_id = ?", args.Note.GroupID).Where("contact_id = ?", args.ContactID).Find(&notes).Error
	if err != nil {
		return nil, err
	}

	return notes, nil
}
