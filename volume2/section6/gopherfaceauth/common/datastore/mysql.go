package datastore

import (
	"database/sql"
	"log"

	"github.com/EngineerKamesh/gofullstack/volume2/section6/gopherfaceauth/models"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDatastore struct {
	*sql.DB
}

func NewMySQLDatastore(dataSourceName string) (*MySQLDatastore, error) {

	connection, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		return nil, err
	}

	return &MySQLDatastore{
		DB: connection}, nil
}

func (m *MySQLDatastore) CreateUser(user *models.User) error {

	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO user(uuid, username, first_name, last_name, email, password_hash) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.UUID, user.Username, user.FirstName, user.LastName, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (m *MySQLDatastore) GetUser(username string) (*models.User, error) {

	stmt, err := m.Prepare("SELECT uuid, username, first_name, last_name, email, password_hash, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM user WHERE username = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(username)
	u := models.User{}
	err = row.Scan(&u.UUID, &u.Username, &u.FirstName, &u.LastName, &u.Email, &u.PasswordHash, &u.TimestampCreated, &u.TimestampModified)
	return &u, err
}

func (m *MySQLDatastore) Close() {
	m.Close()
}
