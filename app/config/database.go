package config

import "github.com/jmoiron/sqlx"

func ConnectMysql() (*sqlx.DB, error) {
	return sqlx.Connect("mysql", "user1:pwd1@(localhost:3306)/db1")
	
}
