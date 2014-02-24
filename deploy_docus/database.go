package deploy_docus

import (
	"database/sql"
	"github.com/lib/pq"
	"os"
)

type Database struct {
	*sql.DB
}

func (d *Database) createTables() (sql.Result, error) {
	command := `CREATE TABLE IF NOT EXISTS repositories (id bigserial PRIMARY KEY, origin text, destination text, rsa_key text);`
	return d.Exec(command)
}

func GetDbConnection() (*Database, error) {
	url := os.Getenv("DATABASE_URL")
	conn, err := pq.ParseURL(url)

	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	database := &Database{db}
	_, err = database.createTables()
	if err != nil {
		return nil, err
	}

	return database, nil
}

func QueryRow(query string, args ...interface{}) (*sql.Row, error) {
	db, err := GetDbConnection()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	row := db.QueryRow(query, args...)
	return row, nil
}
