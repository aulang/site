package controller

import (
	"github.com/kataras/iris/v12/mvc"
	"log"
	"site/repository"
	. "site/service"
)

type IndexController struct {
	beiAnService BeiAnService
}

func (c *IndexController) Get() mvc.Result {
	beiAn, err := c.beiAnService.Get()

	if err != nil {
		log.Printf("获取备案信息失败，%v", err)
	}

	return mvc.View{
		Name: "index.html",
		Data: beiAn,
	}
}

func NewIndexController() *IndexController {
	return &IndexController{
		beiAnService: NewBeiAnService(repository.Collection("beiAn")),
	}
}
