package team

import (
	"github.com/gin-gonic/gin"
	db "project/db"
	"project/model"
)

func Select(c *gin.Context) {
	db := db.Dbfrom()
	name := c.PostForm("name")
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
	var group model.Group
	var n int64
	err = db.Where("name=?", name).First(&group).Count(&n).Error
	if n == 0 {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "你没有队伍",
		})
		return
	}
	if err != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	var apply model.Apply
	err = db.Where("group_title=?", group.Title).Where("state=?", 0).First(&apply).Error
	if err != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	IsPass := c.PostForm("IsPass")
	if IsPass == "T" {
		if err := db.Model(&apply).Update("state", 1).Error; err != nil {
			c.JSON(400, gin.H{
				"code":  400,
				"error": err.Error(),
			})
			return
		}
	} else if IsPass == "F" {
		err = db.Delete(&apply).Error
	}
	if err != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
	})
}
