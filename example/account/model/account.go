package model

type Account struct {
	ID       uint64 `db:"entity_id" json:"id"`
	Name     Name
	Email    EmailAddress
	Password Password
}
