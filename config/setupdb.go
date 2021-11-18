package config

import "database/sql"

func Setup() *sql.DB {

	conf := Getconfig()

	connstr := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ")/" + conf.DB_NAME

	db, err := sql.Open("mysql", connstr)
	if err != nil {
		panic("connection error")
	}
	_ = db.Ping()
	return db
}
