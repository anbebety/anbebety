package private_message

import (
	"github.com/gin-gonic/gin"
	"project/db"
	"project/model"
)

func Send(c *gin.Context) {
	db := db.Dbfrom()
	sender := c.PostForm("sender")
	content := c.PostForm("content")
	receiver := c.PostForm("receiver")
	var n int64
	error := db.Table("users").Where("name=?", sender).Count(&n).Error
	if error != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": error,
		})
		return
	}
	if n == 0 {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "请输入正确的发送者用户名",
		})
		return
	}
	error = db.Table("users").Where("name=?", receiver).Count(&n).Error
	if error != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": error,
		})
		return
	}
	if n == 0 {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "请输入正确的接收者用户名",
		})
		return
	}
	newMessage := model.Message{
		Sender:   sender,
		Content:  content,
		Receiver: receiver,
		State:    0,
	}
	error = db.Create(&newMessage).Error
	if error != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": error,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
	})
}
