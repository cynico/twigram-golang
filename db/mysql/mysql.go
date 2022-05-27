package databse

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var Db *sql.DB

func InitDB() error {
	db, err := sql.Open("mysql", "root:dbpass@tcp(localhost)/twigram")
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}
	Db = db
	return nil
}

func Migrate() error {
	if err := Db.Ping(); err != nil {
		return err
	}
	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrate/mysql",
		"mysql",
		driver,
	)

	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil

}
