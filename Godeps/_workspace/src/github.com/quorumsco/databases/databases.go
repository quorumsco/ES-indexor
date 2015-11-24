// Connects to databases and return the client
package databases

import (
	"time"

	"github.com/iogo-framework/logs"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	//TIMEOUT time between each try
	TIMEOUT = 5 * time.Second
	//RETRY number of tries
	RETRY = 3
)

// Return a sqlx client
func InitSQLX(dialect, args string) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	var i int
retry:
	for {
		db, err = sqlx.Connect(dialect, args)
		switch {
		case err == nil:
			break retry
		case i >= RETRY:
			return nil, err
		default:
			logs.Error(err)
			i++
		}
		time.Sleep(TIMEOUT)
	}

	if db.Ping() != nil {
		return db, err
	}

	return db, nil
}

// Return a gorm client
func InitGORM(dialect, args string) (*gorm.DB, error) {
	var db gorm.DB
	var err error

	var i int
retry:
	for {
		db, err = gorm.Open(dialect, args)
		switch {
		case err == nil:
			break retry
		case i >= RETRY:
			return nil, err
		default:
			logs.Error(err)
			i++
		}
		time.Sleep(TIMEOUT)
	}

	if db.DB().Ping() != nil {
		return &db, err
	}

	return &db, nil
}
