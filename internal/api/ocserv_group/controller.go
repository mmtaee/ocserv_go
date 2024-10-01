package ocserv_group

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"ocserv/internal/repository"
	customErrors "ocserv/pkg/errors"
	"reflect"
)

type Controller struct {
	ocservGroupRepository repository.OcservGroupRepositoryInterface
}

func NewOcservGroupController() *Controller {
	return &Controller{
		ocservGroupRepository: repository.NewOcservGroupRepository(),
	}
}

// List godoc
// @Summary      List Ocserv Group
// @Description  List Ocserv Group
// @Tags         ocserv group
// @Produce      json
// @Param        Authorization header string true "Bearer token"
// @Success      200  {object}  []string
// @Router       /api/v1/ocserv/groups/ [get]
func (controller *Controller) List(c *gin.Context) {
	c.JSON(http.StatusOK, controller.ocservGroupRepository.GroupList())
}

// Create godoc
// @Summary      Create Ocserv Group
// @Description  Create Ocserv Group
// @Tags         ocserv group
// @Produce      json
// @Param        Authorization header string true "Bearer token"
// @Param        site  body     CreateOcservGroupData  true  "Request Body"
// @Success      201  {object}  nil
// @Failure      400  {object}  nil
// @Router       /api/v1/ocserv/groups/ [post]
func (controller *Controller) Create(c *gin.Context) {
	var data CreateOcservGroupData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, customErrors.InvalidBodyError(err))
		return
	}
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "config", reflect.ValueOf(data.Config))
	ctx = context.WithValue(ctx, "name", data.GroupName)
	err := controller.ocservGroupRepository.GroupCreateOrUpdate(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

// Update godoc
// @Summary      Update Ocserv Group
// @Description  Update Ocserv Group
// @Tags         ocserv group
// @Produce      json
// @Param        Authorization header string true "Bearer token"
// @Param        site  body     GroupConfig  true  "Request Body"
// @Success      200  {object}  nil
// @Failure      400  {object}  nil
// @Failure      404  {object}  nil
// @Router       /api/v1/ocserv/groups/:name/ [patch]
func (controller *Controller) Update(c *gin.Context) {
	var data GroupConfig
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, customErrors.InvalidBodyError(err))
		return
	}
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "config", reflect.ValueOf(data))
	ctx = context.WithValue(ctx, "name", c.Param("name"))
	err := controller.ocservGroupRepository.GroupCreateOrUpdate(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

// Delete godoc
// @Summary      Delete Ocserv Group
// @Description  Delete Ocserv Group
// @Tags         ocserv group
// @Produce      json
// @Param        Authorization header string true "Bearer token"
// @Success      204  {object}  nil
// @Failure      400  {object}  nil
// @Router       /api/v1/ocserv/groups/:name/ [delete]
func (controller *Controller) Delete(c *gin.Context) {
	c.Param("name")
	err := controller.ocservGroupRepository.GroupDelete(c.Param("name"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
