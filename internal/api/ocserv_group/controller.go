package ocserv_group

import "github.com/gin-gonic/gin"

type Controller struct{}

func NewOcservGroupController() *Controller {
	return &Controller{}
}

func (controller *Controller) List(c *gin.Context) {}

func (controller *Controller) Create(c *gin.Context) {}

func (controller *Controller) Update(c *gin.Context) {}

func (controller *Controller) Delete(c *gin.Context) {}
