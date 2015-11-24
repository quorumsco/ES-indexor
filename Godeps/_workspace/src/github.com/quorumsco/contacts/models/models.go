// Definition of the structures and SQL interaction functions
package models

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Models return one of every model the database must create..
func Models() []interface{} {
	return []interface{}{
		&Contact{}, &Note{}, &Tag{}, &Mission{}, &Address{},
	}
}
