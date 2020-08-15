package main

import (
	"github.com/aulang/site/config"
	"github.com/aulang/site/controller"
	"github.com/aulang/site/model"
	"github.com/aulang/site/service"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
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

	// Service注册
	mvcApp.Register(service.NewWebConfigService())
	mvcApp.Register(service.NewMenuService())
	mvcApp.Register(service.NewCategoryService())
	mvcApp.Register(service.NewArticleService())
	mvcApp.Register(service.NewCommentService())

	// ROOT
	mvcApp.Handle(new(controller.IndexController))
	// 配置
	mvcApp.Party("/config", crs).Handle(new(controller.WebConfigController))
	// 菜单
	mvcApp.Party("/menus", crs).Handle(new(controller.MenuController))
	// 类别
	mvcApp.Party("/categories", crs).Handle(new(controller.CategoryController))
	// 文章
	mvcApp.Party("/articles", crs).Handle(new(controller.ArticleController))
	// 评论
	mvcApp.Party("/comment", crs).Handle(new(controller.CommentController))
}

func errorHandler(ctx iris.Context, err error) {
	ctx.JSON(model.FailWithError(err))
}
