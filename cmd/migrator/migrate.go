package main

import (
	"database/sql"
	"fmt"

	"github.com/dgdraganov/super-librarian/pkg/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config, err := config.NewDatabaseConfig()
	if err != nil {
		panic(fmt.Errorf("new db config: %w", err))
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("sql open: %w", err))
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(fmt.Errorf("mysql instance: %w", err))
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///bin/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		panic(fmt.Errorf("new databse instance: %w", err))
	}

	m.Steps(1)
}
