package controller

import (
	conf "github.com/Zerohated/project-name/configs"
	"github.com/Zerohated/tools/pkg/logger"
)

var (
	config = conf.Config
	log    = logger.Logger
)

// Controller example
type Controller struct {
	URL string
}

// NewController example
func NewController() *Controller {
	return &Controller{
		URL: config.URL,
	}
}
