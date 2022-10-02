package storage

import (
	"database/sql"
	"goLang/pkg/user"
	"log"
)

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *storage {
	return &storage{db: db}
}

const (
	createUserQuery  = `insert into users (email, name, password_hash) values ($1, $2, '') returning id`
	getUserByIdQuery = `select name, email from users  where id=$1`
	getUsersQuery    = `select * from users`
)

func (s *storage) GetUsers() {}

func (s *storage) GetUser(id int) (*user.User, error) {
	rows, err := s.db.Query(getUserByIdQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user user.User

		if err := rows.Scan(&user.Name, &user.Email); err != nil {
			return nil, err
		}
		log.Printf("username: %s, email: %s", user.Name, user.Email)

		return &user, nil
	}

	return nil, nil
}

func (s *storage) Create(user *user.User) error {
	rows, err := s.db.Query(createUserQuery, user.Email, user.Name)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.UserId); err != nil {
			return err
		}
	}

	log.Println("new user id", user.UserId)

	return nil
}
