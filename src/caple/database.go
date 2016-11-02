package caple

import (
	"log"
	"os"

	pg "gopkg.in/pg.v5"
)

var db *pg.DB

func dbConnect() {
	db = pg.Connect(&pg.Options{
		Addr:     config.dbURL,
		User:     config.dbUser,
		Password: config.dbPassword,
		Database: config.dbName,
	})

	if config.debug {
		pg.SetQueryLogger(log.New(os.Stdout, "", log.LstdFlags))
	}
}
