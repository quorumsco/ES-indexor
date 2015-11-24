// Definition of the structures and SQL interaction functions
package models

import "github.com/jinzhu/gorm"

// ContactDS implements the ContactSQL methods
type ContactDS interface {
	Save(*Contact, ContactArgs) error
	Delete(*Contact, ContactArgs) error
	First(ContactArgs) (*Contact, error)
	Find(ContactArgs) ([]Contact, error)
	FindByMission(*Mission, ContactArgs) ([]Contact, error)

	// FindNotes(*Contact, *ContactArgs) error
	// FindTags(*Contact) error
}

// Contactstore returns a ContactDS implementing CRUD methods for the contacts and containing a gorm client
func ContactStore(db *gorm.DB) ContactDS {
	return &ContactSQL{DB: db}
}
