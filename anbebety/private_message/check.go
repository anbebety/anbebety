package private_message

import (
	"github.com/gin-gonic/gin"
	"project/db"
	"project/dto"
	"project/model"
)

func Check(c *gin.Context) {
	db := db.Dbfrom()
	name := c.PostForm("name")
	cookievalue, _ := c.Cookie(name)
	var n int64
	var session model.Session
	err := db.Where("name=?", name).Find(&session).Error
	if err != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": err,
		})
		return
	}
	if session.Value != cookievalue {
		c.JSON(422, gin.H{
			"code":    422,
			"message": "cookie not exist",
		})
		return
	}
	var message []model.Message
	err = db.Table("messages").Where("receiver=?", name).
		Where("state=?", 1).Find(&message).Count(&n).Error
	if n == 0 {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "没有新邮件",
		})
		return
	}
	if err != nil {
		c.JSON(400, gin.H{
			"code":  400,
			"error": err,
		})
		return
	}
	var ContentGroup dto.CheckDto
	var Content dto.CheckResponse
	for _, information := range message {
		Content.SenderName = information.Sender
		Content.Content = information.Content
		ContentGroup.Records = append(ContentGroup.Records, Content)
		err = db.Model(&message).Update("state", 1).Error
		if err != nil {
			c.JSON(400, gin.H{
				"code":  400,
				"error": err,
			})
		}
	}
	ContentGroup.Total = int(n)
	c.JSON(200, gin.H{
		"code":    200,
		"message": ContentGroup,
	})
}
