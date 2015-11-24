// Definition of the structures and SQL interaction functions
package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// TagSQL contains a Gorm client and the tag and gorm related methods
type TagSQL struct {
	DB *gorm.DB
}

// Save inserts a new tag into the database
func (s *TagSQL) Save(t *Tag, args TagArgs) error {
	if t == nil {
		return errors.New("save: tag is nil")
	}

	var c = &Contact{ID: args.ContactID}

	if t.ID == 0 {
		err := s.DB.Debug().Model(c).Association("Tags").Append(t).Error
		s.DB.Last(t)
		return err
	}

	return s.DB.Model(c).Association("Tags").Replace(t).Error
}

// Delete removes a tag from the database
func (s *TagSQL) Delete(t *Tag, args TagArgs) error {
	return s.DB.Model(&Contact{ID: args.ContactID}).Association("Tags").Delete(t).Error
}

// Find return all the tags containing a given groupID from the database
func (s *TagSQL) Find(args TagArgs) ([]Tag, error) {
	var (
		tags []Tag
		c    = &Contact{ID: args.ContactID}
	)

	err := s.DB.Model(c).Association("Tags").Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}
