package admin

import (
	"github.com/aulang/site/config"
	"github.com/aulang/site/entity"
	"github.com/aulang/site/middleware/oauth"
	. "github.com/aulang/site/model"
	"github.com/aulang/site/service"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mime"
	"strings"
	"time"
)

type ResourceController struct {
	Ctx             iris.Context
	StorageService  service.StorageService
	ResourceService service.ResourceService
}

// GET /admin/resource/subject/{subjectId}
func (c *ResourceController) GetSubjectBy(subjectId string) Response {
	resources, err := c.ResourceService.GetBySubjectID(subjectId)
	if err != nil {
		return FailWithError(err)
	}
	return SuccessWithData(resources)
}

// POST /admin/resource/subject/{subjectId}
func (c *ResourceController) PostSubjectBy(subjectId string) Response {
	user := c.Ctx.User().(*oauth.SimpleUser)

	file, header, err := c.Ctx.FormFile("file")

	if err != nil {
		return FailWithError(err)
	}

	defer file.Close()

	filename := header.Filename
	contentLength := header.Size
	contentType := ContentTypeByFilename(filename)

	resource := entity.Resource{
		ID:            primitive.NewObjectID(),
		Filename:      filename,
		Bucket:        config.Bucket,
		ContentType:   contentType,
		ContentLength: contentLength,
		SubjectID:     subjectId,
		OwnerID:       user.ID,
		OwnerName:     user.Nickname,
		CreationDate:  time.Now(),
	}

	err = c.StorageService.Put(config.Bucket, resource.ID.Hex(), contentType, file, header.Size)
	if err != nil {
		return FailWithError(err)
	}

	c.ResourceService.Save(&resource)
	return SuccessWithData(resource)
}

func ContentTypeByFilename(filename string) string {
	index := strings.LastIndex(filename, ".")
	if index < 0 {
		return "application/octet-stream"
	}

	ext := strings.ToLower(filename[index:])
	typ := mime.TypeByExtension(ext)

	if typ == "" {
		return "application/octet-stream"
	}

	return typ
}

// DELETE /admin/resource/{id}
func (c *ResourceController) DeleteBy(id string) Response {
	if err := c.ResourceService.Delete(id); err != nil {
		return FailWithError(err)
	} else {
		_ = c.StorageService.Remove(config.Bucket, id)
		return Success()
	}
}
