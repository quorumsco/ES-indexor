// Definition of the structures and SQL interaction functions
package models

import "github.com/jinzhu/gorm"

// MissionDS implements the MissionSQL methods
type MissionDS interface {
	Save(*Mission, MissionArgs) error
	Delete(*Mission, MissionArgs) error
	First(MissionArgs) (*Mission, error)
	Find(MissionArgs) ([]Mission, error)
}

// Missionstore returns a MissionDS implementing CRUD methods for the missions and containing a gorm client
func MissionStore(db *gorm.DB) MissionDS {
	return &MissionSQL{DB: db}
}
