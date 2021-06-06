package databases

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresDB manages Postgres connections
type PostgresDB struct {
	DBConn *gorm.DB
}

// Init Postgres database connection
func (db *PostgresDB) Init() error {
	var err error

	dsn := os.Getenv("POSTGRES_DB_CONN_STR")
	db.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("[INFO]: Database connection successfully opened")

	return err
}
