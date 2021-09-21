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

	// 未捕获异常恢复
	app.Use(recover.New())
	// 日志级别
	app.Logger().SetLevel("warn")

	// 初始化MVC
	initMvc(mvc.New(app))

	_ = app.Listen(":"+config.Port(), config.Iris())
}

func initMvc(mvcApp *mvc.Application) {
	// 异常处理
	mvcApp.HandleError(errorHandler)

	// Service注册
	mvcApp.Register(
		service.NewWebConfigService(),
		service.NewCategoryService(),
		service.NewResourceService(),
		service.NewArticleService(),
		service.NewCommentService(),
		service.NewStorageService(),
	)

	// ROOT
	mvcApp.Handle(new(controller.IndexController))
	// 配置
	mvcApp.Party("/config").Handle(new(controller.WebConfigController))
	// 类别
	mvcApp.Party("/category").Handle(new(controller.CategoryController))
	// 文章
	mvcApp.Party("/article").Handle(new(controller.ArticleController))
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
