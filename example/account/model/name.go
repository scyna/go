package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Name struct {
	Value string `db:"name" json:"name"`
}

func (name *Name) Validate() error {
	return validation.ValidateStruct(name,
		validation.Field(&name.Value, validation.Required, validation.Length(0, 100)),
	)
}
