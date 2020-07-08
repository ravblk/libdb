package pg_test

import (
	"ravblk/libdb/model"
	"ravblk/libdb/storage/pg"
	"testing"
)

func BenchmarkSQLxPQInsert(b *testing.B) {
	conn, err := pg.NewSqlxPQConnect(c)
	if err != nil {
		b.Fatal(err)
	}

	defer conn.Close()

	for i := 0; i < b.N; i++ {
		if _, err := conn.NamedExec(`
		INSERT INTO
			users (email, username, password, status, age, created_at)
		VALUES 
			(:email, :username, :password, :status, :age, :created_at)`, u); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSQLxPQGet(b *testing.B) {
	u := &model.User{}
	conn, err := pg.NewSqlxPQConnect(c)
	if err != nil {
		b.Fatal(err)
	}

	defer conn.Close()

	for i := 0; i < b.N; i++ {
		if err := conn.Get(u, `
			SELECT 
				id, 
				email, 
				username, 
				password, 
				status, 
				age, 
				created_at 
			FROM 
				users 
			WHERE 
				id=$1;`, 2); err != nil {
			b.Fatal(err)
		}

	}
}
