package main

import (
	"github.com/y7ut/yio/core/kernel"
	"github.com/y7ut/yio/core/middleware"
	"github.com/y7ut/yio/core/httphandler"
)

const debug = false

func main() {
	
	var BookList = httphandler.Books{"aaa": 50, "bbb": 20}

	var app kernel.YichuApp
	app.Init()
	app.Add(middleware.FirstMiddleware{})
	app.Add(middleware.SecondMiddleware{})
	// 路由
	app.Router("/list", BookList.List)
	app.Router("/price", BookList.Price)
	app.Run("localhost", 8081)
}
