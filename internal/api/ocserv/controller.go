package ocserv

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ocserv/pkg/errors"
)

type Controller struct {
}

func NewOcservController() *Controller {
	return &Controller{}
}

func (controller *Controller) Create(c *gin.Context) {
	var data CreateData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, errors.InvalidBodyError(err))
		return
	}
}
