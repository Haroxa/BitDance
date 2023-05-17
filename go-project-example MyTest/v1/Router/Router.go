package Router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"v1/Controller"
)

func Start() {

	// 开启路由
	r := gin.Default()
	// 使用中间件，记录日志
	r.Use(gin.Logger())
	// 测试请求
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// 获取话题页面
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := Controller.QueryTopicInfo(topicId)
		c.JSON(http.StatusOK, data)
	})
	// 发帖评论
	r.POST("/community/post/do", func(c *gin.Context) {
		uid, _ := c.GetPostForm("uid")
		topicId, _ := c.GetPostForm("topic_id")
		content, _ := c.GetPostForm("content")
		data := Controller.PublishPost(uid, topicId, content)
		c.JSON(http.StatusOK, data)
	})
	// 创建话题
	r.POST("/community/topic/do", func(c *gin.Context) {
		uid, _ := c.GetPostForm("uid")
		title, _ := c.GetPostForm("title")
		content, _ := c.GetPostForm("content")
		data := Controller.PublishTopic(uid, title, content)
		c.JSON(http.StatusOK, data)
	})
	// 注册用户
	r.POST("/community/user/do", func(c *gin.Context) {
		name, _ := c.GetPostForm("name")
		avatar, _ := c.GetPostForm("avatar")
		data := Controller.Register(name, avatar)
		c.JSON(http.StatusOK, data)
	})

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
