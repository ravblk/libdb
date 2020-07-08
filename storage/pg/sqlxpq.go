package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewSqlxPQConnect(c *Config) (*sqlx.DB, error) {
	args := fmt.Sprintf(
		"sslmode=%s host=%s port=%s user=%s password='%s' dbname=%s",
		c.SSL,
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DatabaseName,
	)

	conn, err := sqlx.Connect("postgres", args)
	if err != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(c.MaxIdleConns)
	conn.SetMaxOpenConns(c.MaxOpenConns)

	return conn, nil
}
