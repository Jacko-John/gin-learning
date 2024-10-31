package router

import (
	"gin-test/router/user"

	"github.com/gin-gonic/gin"
)

type routerFunc interface {
	Register(r *gin.Engine)
}

var routerList = make([]routerFunc, 0)

func GetRouter() *gin.Engine {
	// construct router list (只需在这里手动append)
	routerList = append(routerList, &user.Router{})
	// ...

	// create gin engine
	r := newRouter()
	// register all routers
	for _, router := range routerList {
		router.Register(r)
	}
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", nil)
	})
	return r
}

// TODO: 实现common controller避免在不同的router中重复实现 （可以后续优化代码结构的时候再考虑）
