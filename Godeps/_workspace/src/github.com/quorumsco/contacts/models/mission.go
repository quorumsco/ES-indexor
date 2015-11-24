// Definition of the structures and SQL interaction functions
package models

import "time"

// Mission represents the components of a Mission
type Mission struct {
	ID   uint       `gorm:"primary_key" json:"id"`
	Date *time.Time `json:"date,omitempty"`

	GroupID uint `sql:"not null" db:"group_id" json:"-"`

	Contacts []Contact `json:"contacts,omitempty" gorm:"many2many:mission_contacts;"`
}

// MissionArgs is used in the RPC communications between the gateway and Contacts
type MissionArgs struct {
	Mission *Mission
}

// MissonReply is used in the RPC communications between the gateway and Contacts
type MissionReply struct {
	Mission  *Mission
	Missions []Mission
}
