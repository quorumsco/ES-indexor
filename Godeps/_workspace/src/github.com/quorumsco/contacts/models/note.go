// Definition of the structures and SQL interaction functions
package models

import "time"

// Note represents the components of a note
type Note struct {
	ID      uint       `db:"id" json:"id"`
	Content string     `db:"content" json:"content"`
	Author  string     `db:"author" json:"author"`
	Date    *time.Time `db:"date" json:"date"`

	GroupID   uint `db:"group_id" json:"group_id"`
	ContactID uint `db:"contact_id" json:"contact_id"`
}

// NoteArgs is used in the RPC communications between the gateway and Contacts
type NoteArgs struct {
	GroupID   uint
	ContactID uint
	Note      *Note
}

// NoteReply is used in the RPC communications between the gateway and Contacts
type NoteReply struct {
	Note  *Note
	Notes []Note
}
