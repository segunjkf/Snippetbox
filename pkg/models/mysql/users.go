package mysql

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/segunjkf/lets-go/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}


// We'll use the Insert method to add a new record to the users table.
func (m *UserModel) Insert(name, email, password string) error { 
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if HashedPassword != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created) VALUES(?, ?, ?, UTC_TIMESTAMP()`

	// Use The Exec() method to inert the user details and hased password 
	//into the users table

	_, err = m.DB.Exec(stmt, name, email, string(HashedPassword))
	if err != nil {
		var mySQlError *mysql.MySQLError
		if errors.As(err, &mySQlError) {
			if mySQlError.Number == 1062 && strings.Contains(mySQlError.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}

	log.Println("Data insterted")

	return nil
}

// We'll use the Authenticate method to verify whether a user exists with // the provided email address and password. This will return the relevant // user ID if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil 
}

// We'll use the Get method to fetch details for a specific user based // on their user ID.
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil 
}