package main

import (
	"fmt"
	"html/template"
	"time"

	"example.com/basebaseProject3/routers"

	"github.com/gin-gonic/gin"
)

func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()

	fmt.Println("initMiddleware start")
	c.Next()
	fmt.Println("initMiddleware end")
	end := time.Now().UnixNano()
	fmt.Println("initMiddleware cost:", (end-start)/1e6, "ms")
}

func initMiddlewareTwo(c *gin.Context) {

	fmt.Println("1-我是一个中间件-initMiddlewareTwo")
	//调用该请求的剩余处理程序
	c.Next()

	fmt.Println("2-我是一个中间件-initMiddlewareTwo")

}

func main() {

	// 创建一个默认的路由引擎
	r := gin.Default()
	//自定义模板函数  注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
	//加载模板 放在配置路由前面
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")

	//r.Use(initMiddleware, initMiddlewareTwo)

	r.GET("/testForMiddleware", func(c *gin.Context) {
		c.String(200, "testForMiddleware")
	})

	routers.AdminRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.DefaultRoutersInit(r)

	r.Run()
}
