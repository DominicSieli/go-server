package types

import "time"

type Admin struct {
	ID		  int		`json:"id"`
	FirstName string	`json:"firstName"`
	LastName  string	`json:"lastName"`
	Email	  string	`json:"email"`
	Password  string	`json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateAdminPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email	  string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=128"`
}

type LoginAdminPayload struct {
	Email	 string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
