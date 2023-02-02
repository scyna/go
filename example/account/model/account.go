package model

type Account struct {
	ID       uint64       `db:"entity_id" json:"id"`
	Name     Name         `db:"name" json:"name"`
	Email    EmailAddress `db:"email" json:"email"`
	Password Password
}
