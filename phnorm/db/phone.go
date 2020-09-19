package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// DB makes things organised by its useful methods
type DB struct {
	db *sql.DB
}

// Phone repserents the phone_numbers table in the DB
type Phone struct {
	ID     int
	Number string
}

// Open is abstraction around sql.Open. Returns a pointer to DB
// if successful.
func Open(driverName, dataSource string) (*DB, error) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// Close closes the db instance Open'ed by Open
func (db *DB) Close() error {
	return db.db.Close()
}

// Seed populates the table with base data
func (db *DB) Seed() error {
	data := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}

	for _, number := range data {
		if _, err := insertPhone(db.db, number); err != nil {
			return err
		}
	}

	return nil
}

// FindPhone takes a number and returns a pointer to a Phone
func (db *DB) FindPhone(number string) (*Phone, error) {
	var p Phone
	row := db.db.QueryRow("SELECT * FROM phone_numbers WHERE value=$1", number)
	err := row.Scan(&p.ID, &p.Number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &p, nil
}

// UpdatePhone takes a pointer to a Phone and upadates that to the database.
func (db *DB) UpdatePhone(p *Phone) error {
	statement := `UPDATE phone_numbers SET value=$2 WHERE id=$1`
	_, err := db.db.Exec(statement, p.ID, p.Number)
	return err
}

// DeletePhone removes a phone number by it's ID.
func (db *DB) DeletePhone(id int) error {
	statement := `DELETE FROM phone_numbers WHERE id=$1`
	_, err := db.db.Exec(statement, id)
	return err
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// AllPhones returns a slice of all entries in phone_numbers table.
func (db *DB) AllPhones() ([]Phone, error) {
	return getAllPhones(db.db)
}

func getAllPhones(db *sql.DB) ([]Phone, error) {
	rows, err := db.Query("SELECT id, value FROM phone_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var numbers []Phone

	for rows.Next() {
		var p Phone
		if err := rows.Scan(&p.ID, &p.Number); err != nil {
			return nil, err
		}
		numbers = append(numbers, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

func getPhone(db *sql.DB, id int) (string, error) {
	var number string
	row := db.QueryRow("SELECT value FROM phone_numbers WHERE id=$1", id)
	err := row.Scan(&number)
	if err != nil {
		return "", err
	}

	return number, nil
}

// Migrate
func Migrate(driverName, dataSource string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}

	err = createPhoneNumbersTable(db)
	if err != nil {
		return nil
	}
	return db.Close()
}

func createPhoneNumbersTable(db *sql.DB) error {
	statement := `
		CREATE TABLE IF NOT EXISTS phone_numbers (
			id SERIAL,
			value VARCHAR(255)
		)`

	_, err := db.Exec(statement)
	return err
}

// Reset drops the database and creates a new one
func Reset(driverName, dataStore, dbName string) error {
	db, err := sql.Open(driverName, dataStore)
	if err != nil {
		return err
	}
	err = resetDB(db, dbName)
	if err != nil {
		return err
	}
	return db.Close()
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}

	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}

	return nil
}
