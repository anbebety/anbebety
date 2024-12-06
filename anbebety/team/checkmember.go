package team

import (
	"github.com/gin-gonic/gin"
	db "project/db"
	"project/model"
)

func Checkmember(c *gin.Context) {
	db := db.Dbfrom()
	name := c.PostForm("name")
	member := c.PostForm("member")
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
	var n int64
	var use model.Apply
	err = db.Where("name = ?", member).First(&use).Count(&n).Error
	if err != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	if n == 0 {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "成员不存在",
		})
		return
	}
	var user model.User
	err = db.Where("name = ?", member).First(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":      200,
		"telephone": user.Telephone,
		"message":   user.PersonalInformation,
	})
}
