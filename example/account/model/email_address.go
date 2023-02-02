package model

import validation "github.com/go-ozzo/ozzo-validation"

type EmailAddress struct {
	EmailPattern string
	Value        string
}

func (emailAddress *EmailAddress) Validate() error {
	return validation.ValidateStruct(emailAddress,
		validation.Field(&emailAddress.Value, validation.Required, validation.Length(0, 100)),
		validation.Field(&emailAddress.EmailPattern, validation.Required, validation.Length(0, 100)),
	)
}
