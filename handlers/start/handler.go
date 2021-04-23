package start

import (
	"umock/logger"
	"umock/models"
)

type StartHandler struct {
	conf models.Config
	log  *logger.Logger
}

func New(conf models.Config, log *logger.Logger) (*StartHandler, error) {
	return &StartHandler{
		conf,
		log,
	}, nil
}
