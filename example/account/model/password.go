package model

import validation "github.com/go-ozzo/ozzo-validation"

type Password struct {
	Value string `db:"password" json:"password"`
}

func (password *Password) Validate() error {
	return validation.ValidateStruct(password,
		validation.Field(&password.Value, validation.Required, validation.Length(0, 100)),
	)
}
