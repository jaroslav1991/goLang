//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=interfaces_mock.go
package users

import "goLang/pkg/user"

type Storage interface {
	Create(user *user.User) error
	GetUser(id int) (*user.User, error)
	FindUsersEmail(email string) ([]user.User, error)
	FindUsersNameEmail(name, email string) ([]user.User, error)
}
