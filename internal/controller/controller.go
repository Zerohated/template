package controller

import (
	"git.52retail.com/david/tools/pkg/logger"
	conf "gitee.com/evolveZ/project-name/configs"
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
