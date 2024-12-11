package databases

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/oussamasf/pito/internal/config"
)

func Init(c *config.Config) (*sql.DB, error) {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=verify-full",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName,
	)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Print(err)
		log.Fatal("ff")
	}

	return db, err
}
