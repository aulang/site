package controller

import (
	"log"
	"site/repository"
	. "site/service"
)

type IndexController struct {
	beiAnService BeiAnService
}

func (c *IndexController) Get() interface{} {
	beiAn, err := c.beiAnService.Get()

	if err != nil {
		log.Printf("获取备案信息失败，%v", err)
	}

	return beiAn
}

func NewIndexController() *IndexController {
	return &IndexController{
		beiAnService: NewBeiAnService(repository.Collection("beiAn")),
	}
}
