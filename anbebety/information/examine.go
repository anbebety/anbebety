package information

import (
	"github.com/gin-gonic/gin"
	"project/db"
	"project/model"
)

func Examine(c *gin.Context) {
	db := db.Dbfrom()
	title := c.PostForm("title")
	name := c.PostForm("name")
	var n int64
	var group model.Group
	error := db.Table("groups").Where("title=?", title).Count(&n).Error
	db.Where("title=?", title).First(&group)
	if error != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": error.Error(),
		})
		return
	}
	if n == 0 {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "队伍不存在",
		})
		return
	}
	cookievalue, _ := c.Cookie(name)
	var session model.Session
	error = db.Table("Sessions").Where("value=?", cookievalue).Count(&n).Error
	db.Where("value=?", cookievalue).First(&session)
	if session.Name != name {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "cookie not exist",
		})
		return
	}
	var user model.User
	error = db.Table("Users").Where("name=?", session.Name).Count(&n).Error
	db.Where("name=?", session.Name).First(&user)
	if error != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": error.Error(),
		})
		return
	}
	if user.Identity != 1 {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "未获得权限",
		})
		return
	}
	IsPass := c.PostForm("IsPass")
	if IsPass == "T" {
		if err := db.Model(&group).Update("state", 1).Error; err != nil {
			c.JSON(400, gin.H{
				"code":  400,
				"error": err.Error(),
			})
			return
		}
	} else if IsPass == "F" {
		error = db.Delete(&group).Error
	}
	if error != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
	})
}
