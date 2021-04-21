package start

import "github.com/chonla/umock/models"

type StartHandler struct {
	conf models.Config
}

func New(conf models.Config) (*StartHandler, error) {
	return &StartHandler{
		conf,
	}, nil
}
