package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BeiAnEntry struct {
	No  string `json:"no"`  //备案号
	Url string `json:"url"` //备案查询地址
}

type BeiAn struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	MIIT BeiAnEntry         `json:"miit"` //工信部备案
	MPS  BeiAnEntry         `json:"mps"`  //公安部备案
}
