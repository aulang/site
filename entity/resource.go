package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Resource struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Filename      string             `json:"filename" bson:"filename"`           // 文件名
	Bucket        string             `json:"bucket" bson:"bucket"`               // 对象存储桶名
	ContentType   string             `json:"contentType" bson:"contentType"`     // 对象类型
	ContentLength int64              `json:"contentLength" bson:"contentLength"` // 对象类型

	SubjectID string `json:"subjectId" bson:"subjectId"` // 关联对象ID

	OwnerID   string `json:"ownerId" bson:"ownerId"`     // 所有者ID
	OwnerName string `json:"ownerName" bson:"ownerName"` // 所有者名称

	CreationDate time.Time `json:"creationDate" bson:"creationDate"` // 创建日期
}
