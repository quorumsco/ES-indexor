// Definition of the structures and SQL interaction functions
package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// ContactSQL contains a Gorm client and the mission and gorm related methods
type MissionSQL struct {
	DB *gorm.DB
}

// Save inserts a new mission into the database
func (s *MissionSQL) Save(m *Mission, args MissionArgs) error {
	if m == nil {
		return errors.New("save: mission is nil")
	}

	if m.ID == 0 {
		return s.DB.Save(m).Error
	}

	return s.DB.Update(m).Error
}

// Delete removes a mission from the database
func (s *MissionSQL) Delete(m *Mission, args MissionArgs) error {
	if m == nil {
		return errors.New("save: mission is nil")
	}

	return s.DB.Delete(m).Error
}

// Find returns all the mission with a given groupID from the database
func (s *MissionSQL) Find(args MissionArgs) ([]Mission, error) {
	var missions []Mission

	err := s.DB.Where(args.Mission).Find(&missions).Error

	if err != nil {
		return nil, err
	}

	return missions, nil
}

// First returns a mission from the database using it's ID
func (s *MissionSQL) First(args MissionArgs) (*Mission, error) {
	var m Mission

	if err := s.DB.Where(args.Mission).First(&m).Error; err != nil {
		if err == gorm.RecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &m, nil
}
