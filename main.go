package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"site/config"
	"site/controller"
)

func main() {
	app := iris.New()

	app.Use(recover.New())
	app.Logger().SetLevel("warn")

	mvc.New(app).Handle(controller.NewIndexController())

	app.Listen(":"+config.Port(), config.Iris())
}
