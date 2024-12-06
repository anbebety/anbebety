package logout

import (
	"github.com/gin-gonic/gin"
	"project/db"
	"project/model"
)

func Logout(c *gin.Context) {
	name := c.PostForm("name")
	db := db.Dbfrom()
	var session model.Session
	cookievalue, err := c.Cookie(name)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}
	err = db.Where("value = ?", cookievalue).First(&session).Error
	if err != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	if session.Name != name {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "cookie not exist",
		})
		return
	}
	db.Delete(&session)
	c.SetCookie(name, cookievalue, -1, "/", "", false, false)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "退出成功",
	})
}
