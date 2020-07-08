package main

import (
	"ravblk/libdb/storage/pg"

	"go.uber.org/zap"
)

var schema = `CREATE TABLE users(
	id           BIGSERIAL PRIMARY KEY,
	email text,
	username text,
	password text,
	status int,
	age int,
	created_at TIMESTAMP
 );`

func main() {
	log, _ := zap.NewProduction()
	c := &pg.Config{
		Host:         "localhost",
		Port:         "5432",
		User:         "postgres",
		Password:     "postgres",
		DatabaseName: "test",
		Schema:       "public",
		SSL:          "disable",
		MaxIdleConns: 50,
		MaxOpenConns: 50,
	}

	conn, err := pg.NewSqlxPQConnect(c)
	if err != nil {
		log.Fatal("", zap.Error(err))
	}

	conn.MustExec(schema)

	log.Info("create table OK")
}
