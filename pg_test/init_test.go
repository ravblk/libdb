package pg_test

import (
	"ravblk/libdb/model"
	"ravblk/libdb/storage/pg"
	"time"
)

var (
	c = &pg.Config{
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
	u = &model.User{
		Email:     "user@mail.ru",
		Username:  "karim",
		Password:  "123456",
		Status:    4,
		Age:       30,
		CreatedAt: time.Now(),
	}
)
