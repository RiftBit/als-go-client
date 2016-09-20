package alsgoclient

import (
	"gopkg.in/go-playground/validator.v9"
)

type Config struct {
	Uri		string	`validate:"required"`
	Login		string	`validate:"required"`
	Password	string	`validate:"required"`
	IsAsync		bool
	Timeout		int	`validate:"required,gte=50,lte=60000"`
}

func (s *Config) validate() (bool, error) {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		return false, err
	}
	return true, nil
}

