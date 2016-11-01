package caple

import pg "gopkg.in/pg.v5"

var db *pg.DB

func init() {
	db = pg.Connect(&pg.Options{
		Addr:     config.dbURL,
		User:     config.dbUser,
		Password: config.dbPassword,
		Database: config.dbName,
	})
}
