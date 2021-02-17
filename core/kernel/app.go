package kernel

import (
	"log"
	"fmt"
	"net/http"
	"github.com/y7ut/yio/core/middleware"
	"github.com/y7ut/yio/core/httphandler"
)


// YichuApp 我的应用
type YichuApp struct {
	servermux         *http.ServeMux
	globalMiddlewares []middleware.Middleware
}

// Init 初始化
func (app *YichuApp) Init() {
	app.servermux = http.NewServeMux()
}

// Add 添加中间件
func (app *YichuApp) Add(middlewares middleware.Middleware) {
	app.globalMiddlewares = append(app.globalMiddlewares, middlewares)
}

// Router 路由
func (app *YichuApp) Router(urlstring string, handler func(res http.ResponseWriter, r *http.Request)) {
	var currentHandler httphandler.MyHandler
	for _, v := range app.globalMiddlewares {
		currentHandler.AddMiddleware(v)
	}
	currentHandler.Callable = http.HandlerFunc(handler)
	app.servermux.Handle(urlstring, currentHandler)
}

// Run 启动应用服务
func (app *YichuApp) Run(host string, port int) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), app.servermux))
}