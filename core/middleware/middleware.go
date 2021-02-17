package middleware

import (
	"fmt"
	"log"
	"net/http"
)

// Middleware 中间件
type Middleware interface {
	Process(r *http.Request, res http.ResponseWriter, h HandlerWithMiddleware)
}

// MiddlewareList  中間件清單
type MiddlewareList struct {
	List  []Middleware
	index int
}

// Next 下一個
func (ml *MiddlewareList) Next() {
	ml.index++
}

// Current 當前
func (ml *MiddlewareList) Current() Middleware {
	if ml.index >= len(ml.List) {
		return nil
	}
	return ml.List[ml.index]
}

// FirstMiddleware 测试中间件1号
type FirstMiddleware struct {

}

// Process 实现了中间件接口
func (m FirstMiddleware) Process(r *http.Request, w http.ResponseWriter, h HandlerWithMiddleware) {
	log.Println("go to First")
	h.HandlerIt(r, w)
	log.Println("bye  First!")
	fmt.Fprintf(w, "edit by first")
}


// SecondMiddleware 测试中间件2号
type SecondMiddleware struct {

}

// Process 实现了中间件接口
func (m SecondMiddleware) Process(r *http.Request, w http.ResponseWriter, h HandlerWithMiddleware) {
	log.Println("go to Second")
	fmt.Fprintf(w, "edit by second")
	h.HandlerIt(r, w)
	log.Println("bye  Second!")
}


// HandlerWithMiddleware 使用中间件的handle
type HandlerWithMiddleware interface {
	HandlerIt(r *http.Request, res http.ResponseWriter)
	AddMiddleware(m Middleware)
}

