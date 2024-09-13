package site

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ocserv/internal/models"
	"ocserv/internal/repository"
)

type Controller struct {
	siteRepository repository.SiteRepositoryInterface
}

func NewSiteController() *Controller {
	return &Controller{
		siteRepository: repository.NewSiteRepository(),
	}
}

// Get godoc
// @Summary      Get site configuration
// @Description  Get site configuration
// @Tags         site
// @Produce      json
// @Success      200  {object}  models.Site
// @Failure      400  {object}  map[string]string  "error": "Error message"
// @Example 400 {object} {"error": "Detailed error message"}
// @Router       /api/site [get]
func (controller *Controller) Get(c *gin.Context) {
	site, err := controller.siteRepository.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, site)
}

func (controller *Controller) Create(c *gin.Context) {
	type Data struct {
		CaptchaSiteKey   string  `json:"captcha_site_key"  binding:"omitempty"`
		CaptchaSecretKey string  `json:"captcha_secret_key"  binding:"omitempty"`
		DefaultTraffic   float64 `json:"default_traffic" binding:"required"`
	}
	var data Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if data.DefaultTraffic <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid default_traffic"})
		return
	}
	site := models.Site{
		CaptchaSiteKey:   data.CaptchaSiteKey,
		CaptchaSecretKey: data.CaptchaSecretKey,
		DefaultTraffic:   data.DefaultTraffic,
	}
	err := controller.siteRepository.Create(&site)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": site})
}

func (controller *Controller) Update(c *gin.Context) {
	if isStaff, exists := c.Get("isStaff"); !exists || !isStaff.(bool) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "only admin can update site configs",
		})
		return
	}

	type Data struct {
		CaptchaSiteKey   string  `json:"captcha_site_key"  binding:"omitempty"`
		CaptchaSecretKey string  `json:"captcha_secret_key"  binding:"omitempty"`
		DefaultTraffic   float64 `json:"default_traffic"  binding:"omitempty"`
	}

	var data Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var site models.Site
	if data.DefaultTraffic <= 0 {
		site.DefaultTraffic = data.DefaultTraffic
	}
	if data.CaptchaSiteKey != "" {
		site.CaptchaSiteKey = data.CaptchaSiteKey
	}
	if data.CaptchaSecretKey != "" {
		site.CaptchaSecretKey = data.CaptchaSecretKey
	}
	update, err := controller.siteRepository.Update(&site)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, update)
}
