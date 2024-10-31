package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// controller functions
func (u *Router) getlist(c *gin.Context) {
	// TODO: get user list from database
	c.JSON(http.StatusOK, gin.H{
		"message": "user list",
	})
}

func (u *Router) getid(c *gin.Context) {
	id := c.Param("id") // 获取url中的参数
	c.JSON(http.StatusOK, gin.H{
		"message": "user " + id,
	}) // 响应json数据
}

func (u *Router) main(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"Name": "Freshman"}) // 响应html页面
}

func (u *Router) create(c *gin.Context) {
	name := c.DefaultPostForm("name", "unknown")
	password := c.DefaultPostForm("password", "unknown")
	c.JSON(http.StatusOK, gin.H{
		"message":  "user created",
		"name":     name,
		"password": password,
	})
}

func (u *Router) login(c *gin.Context) {
	params := make(map[string]any)
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	name := params["name"]
	password := params["password"]
	c.JSON(http.StatusOK, gin.H{
		"message":  "user login",
		"name":     name,
		"password": password,
	})
}

func (u *Router) deferTest(c *gin.Context) {
	// defer handler.DefaultHandle(c)
	num1 := 1
	num0 := 0
	num := num1 / num0
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"num":     num,
	})
}
