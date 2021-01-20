package main

import (
	"github.com/aulang/site/config"
	"github.com/aulang/site/controller"
	"github.com/aulang/site/controller/admin"
	"github.com/aulang/site/middleware/oauth"
	"github.com/aulang/site/model"
	"github.com/aulang/site/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()

	app.Use(recover.New())

	app.Logger().SetLevel("warn")

	initMVC(mvc.New(app))

	_ = app.Listen(":"+config.Port(), config.Iris())
}

func initMVC(mvcApp *mvc.Application) {
	mvcApp.HandleError(errorHandler)

	// Service注册
	mvcApp.Register(service.NewWebConfigService())
	mvcApp.Register(service.NewCategoryService())
	mvcApp.Register(service.NewArticleService())
	mvcApp.Register(service.NewCommentService())
	mvcApp.Register(service.NewResourceService())
	mvcApp.Register(service.NewStorageService())

	// ROOT
	mvcApp.Handle(new(controller.IndexController))
	// 配置
	mvcApp.Party("/config").Handle(new(controller.WebConfigController))
	// 类别
	mvcApp.Party("/categories").Handle(new(controller.CategoryController))
	// 文章
	mvcApp.Party("/articles").Handle(new(controller.ArticleController))
	// 评论
	mvcApp.Party("/comment").Handle(new(controller.CommentController))
	// 资源
	mvcApp.Party("/resource").Handle(new(controller.ResourceController))

	// 认证
	auth := oauth.New()
	// Admin
	adminMvc := mvcApp.Party("/admin", auth)

	adminMvc.Party("/config").Handle(new(admin.ConfigController))
	adminMvc.Party("/article").Handle(new(admin.ArticleController))
	adminMvc.Party("/resource").Handle(new(admin.ResourceController))
}

func errorHandler(ctx iris.Context, err error) {
	_, _ = ctx.JSON(model.FailWithError(err))
}
