package controller

import (
	"net/http"

	"gitee.com/evolveZ/project-name/internal/model"
	"github.com/gin-gonic/gin"
)

func (controller *Controller) EchoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.RespOK{Code: model.CodeOK, Data: nil})
	return
}
