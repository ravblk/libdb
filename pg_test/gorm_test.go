package pg_test

import (
	"ravblk/libdb/model"
	"ravblk/libdb/storage/pg"
	"testing"
)

func BenchmarkGORMInsert(b *testing.B) {
	conn, err := pg.NewGORMConnect(c)
	if err != nil {
		b.Fatal(err)
	}

	defer conn.Close()

	us := make([]model.User, 1000)
	for i := 0; i < 1000; i++ {
		us[i] = *u
	}

	for i := 0; i < b.N; i++ {
		if err := conn.Create(&us[i]).Error; err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGORMGet(b *testing.B) {
	conn, err := pg.NewGORMConnect(c)
	if err != nil {
		b.Fatal(err)
	}

	defer conn.Close()

	user := &model.User{}

	for i := 0; i < b.N; i++ {
		if err := conn.Where(&model.User{ID: 10}).First(&user).Error; err != nil {
			b.Fatal(err)
		}
	}
}
