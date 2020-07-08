package pg_test

import (
	"ravblk/libdb/model"
	"ravblk/libdb/storage/pg"
	"testing"
)

func BenchmarkPgxInsert(b *testing.B) {
	conn, err := pg.NewPgxConnect(c)
	if err != nil {
		b.Fatal(err)
	}

	defer conn.Close()

	for i := 0; i < b.N; i++ {
		if _, err := conn.Exec(`
			INSERT INTO 
				users 
				(email, username, password, status, age, created_at) 
			VALUES ($1,$2,$3,$4,$5,$6)`,
			u.Email,
			u.Username,
			u.Password,
			u.Status,
			u.Age,
			u.CreatedAt); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPgxGet(b *testing.B) {
	u := &model.User{}
	conn, err := pg.NewPgxConnect(c)
	if err != nil {
		b.Fatal(err)
	}

	defer conn.Close()

	for i := 0; i < b.N; i++ {
		if err := conn.QueryRow(`
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
			id=$1;`, 101).Scan(&u.ID,
			&u.Email,
			&u.Username,
			&u.Password,
			&u.Status,
			&u.Age,
			&u.CreatedAt); err != nil {
			b.Fatal(err)
		}
	}
}
