package pg

import (
	"strconv"

	"github.com/jackc/pgx"
)

func NewPgxConnect(c *Config) (*pgx.ConnPool, error) {
	port, err := strconv.Atoi(c.Port)
	if err != nil {
		return nil, err
	}

	connConfig := pgx.ConnConfig{
		Host:     c.Host,
		Port:     uint16(port),
		Database: c.DatabaseName,
		User:     c.User,
		Password: c.Password,
	}

	pgPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{ConnConfig: connConfig, MaxConnections: c.MaxOpenConns})
	if err != nil {
		return nil, err
	}

	return pgPool, nil
}
