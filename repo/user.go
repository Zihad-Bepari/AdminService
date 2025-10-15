package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID            int    `json:"id" db:"id"`
	ExternalId    string `json:"externalid" db:"external_id"`
	Email         string `json:"email" db:"email"`
	Name          string `json:"name" db:"name"`
	Password_Hash string `json:"password_hash" db:"password_hash"`
	Role          string `json:"role" db:"role"`
	Immutable     bool   `json:"immutable" db:"immutable"`
}

type userrepo struct {
	dbcon *sqlx.DB
}

type Userrepo interface {
	Create(user User) (*User, error)
	UpdateRole(id int, role string) error
}

func Newuserrepo(dbcon *sqlx.DB) Userrepo {
	return &userrepo{
		dbcon: dbcon,
	}
}

func (r userrepo) Create(user User) (*User, error) {
	query := `
	   INSERT INTO users(
	       external_id,
		   email,
		   name,
		   password_hash,
		   role,
		   immutable  
	   )
		   VALUES(
		       :external_id,
			   :email,
			   :name,
			   :password_hash,
			   :role,
			   :immutable
		   )
	    RETURNING id
	`

	var userID int
	rows, err := r.dbcon.NamedQuery(query, user)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&userID)
	}

	user.ID = userID
	return &user, nil
}
func (r *userrepo) UpdateRole(id int, role string) error {
	query := `
		UPDATE users 
		SET role = $1, updated_at = NOW() 
		WHERE id = $2
	`
	if _, err := r.dbcon.Exec(query, role, id); err != nil {
		return fmt.Errorf("failed to update role for user %d: %w", id, err)
	}
	return nil
}
