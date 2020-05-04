package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"site/config"
	"site/controller"
	"site/model"
)

func main() {
	app := iris.New()

	app.Use(recover.New())
	app.Logger().SetLevel("warn")

	initMVC(mvc.New(app))

	app.Listen(":"+config.Port(), config.Iris())
}

func initMVC(mvcApp *mvc.Application) {
	mvcApp.HandleError(errorHandler)

	// ROOT
	mvcApp.Handle(controller.NewIndexController())

	// 配置
	mvcApp.Party("/config").Handle(controller.NewWebConfigController())
	// 菜单
	mvcApp.Party("/menus").Handle(controller.NewMenuController())
}

func errorHandler(ctx iris.Context, err error) {
	ctx.JSON(model.FailWithError(err))
}
