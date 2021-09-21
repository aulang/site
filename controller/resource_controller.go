package controller

import (
	"github.com/aulang/site/config"
	"github.com/aulang/site/service"
	"github.com/kataras/iris/v12"
	"net/http"
)

type ResourceController struct {
	Ctx             iris.Context
	StorageService  service.StorageService
	ResourceService service.ResourceService
}

// GetBy GET /resource/{id}
func (c *ResourceController) GetBy(id string) {
	resource, err := c.ResourceService.GetByID(id)
	if err != nil {
		c.Ctx.StopWithError(http.StatusNotFound, err)
	}

	reader, err := c.StorageService.Get(config.Bucket, resource.ID.Hex())
	if err != nil {
		c.Ctx.StopWithError(http.StatusInternalServerError, err)
	}
	c.Ctx.ContentType(resource.ContentType)
	c.Ctx.ServeContent(reader, resource.Filename, resource.CreationDate)
}
