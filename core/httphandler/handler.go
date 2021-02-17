package httphandler

import (
	"github.com/y7ut/yio/core/middleware"
	"net/http"
)

// MyHandler 我的处理器
type MyHandler struct {
	ml middleware.MiddlewareList
	Callable http.HandlerFunc
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.HandlerIt(req, w)
}

// HandlerIt 处理逻辑
func (h *MyHandler) HandlerIt(r *http.Request, w http.ResponseWriter) {
	middleware := h.ml.Current()
	h.ml.Next()
	if middleware == nil {
		h.Callable(w, r)
		return
	}
	middleware.Process(r, w, h)
}


// AddMiddleware 添加中间件
func (h *MyHandler) AddMiddleware(m middleware.Middleware) {
	h.ml.List = append(h.ml.List, m)
}
