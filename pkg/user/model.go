package user

type User struct {
	UserId int64  `db:"id" json:"userId"`
	Name   string `db:"name" json:"name"`
	Email  string `db:"email" json:"email"`
}
