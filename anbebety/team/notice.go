package team

import (
	"github.com/gin-gonic/gin"
	db "project/db"
	"project/model"
)

func Notice(c *gin.Context) {
	db := db.Dbfrom()
	name := c.PostForm("name")
	content := c.PostForm("content")
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
	err = db.Where("name = ?", name).First(&group).Error
	if err != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	var n int64
	var apply []model.Apply
	err = db.Table("applies").Where("group_title=?", group.Title).
		Where("state=?", 1).Find(&apply).Count(&n).Error
	if err != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	for _, information := range apply {
		sender := name
		receiver := information.Name
		newMessage := model.Message{
			Sender:   sender,
			Content:  content,
			Receiver: receiver,
			State:    0,
		}
		err = db.Create(&newMessage).Error
		if err != nil {
			c.JSON(400, gin.H{
				"code":  400,
				"error": err,
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
	})
	return
}
