package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MustLogin() gin.HandlerFunc {  //必须登录
	return func(c *gin.Context) {
		if _,status:=c.GetQuery("token");!status{
			c.String(http.StatusUnauthorized,"缺少token参数")
			c.Abort()
		}else{
			c.Next()
		}
	}
}

func GetTopicDetail(c *gin.Context)  {
   c.JSON(200,CreateTopic(101,"帖子标题"))
}
func NewTopic(c *gin.Context)  { //单帖子新增
	topic:=Topic{}
	err:=c.BindJSON(&topic)
	if err!=nil{
		c.String(400,"参数错误:%s",err.Error())
	}else{
		c.JSON(200,topic)
	}
}
func NewTopics(c *gin.Context)  { //多帖子批量新增
	topics:=Topics{}
	err:=c.BindJSON(&topics)
	if err!=nil{
		c.String(400,"参数错误:%s",err.Error())
	}else{
		c.JSON(200,topics)
	}
}

func DelTopic(c *gin.Context)  {
	//判断登录
	c.String(200,"删除帖子")
}
func GetTopicList(c *gin.Context)  {
	 query:=TopicQuery{}
	 err:=c.BindQuery(&query)
	 if err!=nil{
	 	c.String(400,"参数错误:%s",err.Error())
	 }else{
	 	c.JSON(200,query)
	 }
}