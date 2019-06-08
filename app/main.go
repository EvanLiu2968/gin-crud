package main

import "github.com/kataras/iris"

func main() {
	app := iris.Default()
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"code": "00",
			"message": "请求成功",
			"response": "",
		})
	})
	
	app.Run(iris.Addr(":8080"))
}
