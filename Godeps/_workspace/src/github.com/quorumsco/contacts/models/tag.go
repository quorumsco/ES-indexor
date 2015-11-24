// Definition of the structures and SQL interaction functions
package models

// Tag represents the components of a tag
type Tag struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Name  string `json:"name" sql:"not null;"`
	Color string `json:"color"`
}

// TagArgs is used in the RPC communications between the gateway and Contacts
type TagArgs struct {
	GroupID   uint
	ContactID uint
	Tag       *Tag
}

// TagReply is used in the RPC communications between the gateway and Contacts
type TagReply struct {
	Tag  *Tag
	Tags []Tag
}
