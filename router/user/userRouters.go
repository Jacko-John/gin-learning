package user

import (
	"github.com/gin-gonic/gin"
)

type Router struct{}

func (u *Router) Register(r *gin.Engine) {
	r.GET("", u.main) // 主页
	r.GET("/test", u.deferTest)
	// router group for user
	user := r.Group("/user")
	{
		// user routes
		user.GET("/list", u.getlist)
		user.GET("/:id", u.getid)      // 可以通过/user/aaa来获取aaa参数 （详情查看getid函数）
		user.POST("/create", u.create) // 获取表单数据
		user.POST("/login", u.login)   // 响应json数据
	}
}
