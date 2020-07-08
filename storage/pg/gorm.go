package pg

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewGORMConnect(c *Config) (*gorm.DB, error) {
	args := fmt.Sprintf(
		"sslmode=%s host=%s port=%s user=%s password='%s' dbname=%s",
		c.SSL,
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DatabaseName,
	)

	db, err := gorm.Open("postgres", args)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxOpenConns(c.MaxOpenConns)
	db.DB().SetMaxIdleConns(c.MaxIdleConns)

	return db, nil
}
