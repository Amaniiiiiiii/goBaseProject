package admin

import (
	"example.com/basebaseProject3/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	// c.String(200, "用户列表--")

	// //查询数据库全部内容

	// albumList := []models.Album{}
	// models.DB.Find(&albumList)
	// c.JSON(200, gin.H{
	// 	"result": albumList,
	// })

	//查询id > 3 的数据
	albumList := []models.Album{}
	models.DB.Where("id > ?", 3).Find(&albumList)
	c.JSON(200, gin.H{
		"result": albumList,
	})

	con.success(c)
}
func (con UserController) Add(c *gin.Context) {

	//增加
	album := models.Album{
		Title:  "测试标题",
		Artist: "测试作者",
		Price:  100.00,
	}

	models.DB.Create(&album)

	c.String(200, "用户列表-add---")
}
func (con UserController) Edit(c *gin.Context) {
	//修改
	album := models.Album{
		Id: 1,
	}

	models.DB.Find(&album)
	album.Title = "测试标题-修改"

	models.DB.Save(&album)

	c.String(200, "用户列表-Edit------")
}
