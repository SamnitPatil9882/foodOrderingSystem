package boltdb

import "database/sql"

type BaseRepsitory struct {
	DB *sql.DB
}
