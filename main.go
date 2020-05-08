package main

import (
	"github.com/iris-contrib/middleware/cors"
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

	// CORS
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	// ROOT
	mvcApp.Handle(controller.NewIndexController())

	// 配置
	mvcApp.Party("/config", crs).Handle(controller.NewWebConfigController())
	// 菜单
	mvcApp.Party("/menus", crs).Handle(controller.NewMenuController())
	// 文章
	mvcApp.Party("/articles", crs).Handle(controller.NewArticleController())
	// 评论
	mvcApp.Party("/comment", crs).Handle(controller.NewCommentController())
	// 类别
	mvcApp.Party("/categories", crs).Handle(controller.NewCategoryController())
}

func errorHandler(ctx iris.Context, err error) {
	ctx.JSON(model.FailWithError(err))
}
