package main

import (
	"fmt"
	. "gin-demo/src"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"regexp"
)

func main() {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err1 := v.RegisterValidation("topicurl", func(fl validator.FieldLevel) bool {

			_, ok1 := fl.Top().Interface().(*Topics)
			_, ok2 := fl.Top().Interface().(*Topic)
			if ok1 || ok2 {

				fmt.Println(fl.FieldName())
				if matched, _ := regexp.MatchString(`^\w{4,10}$`, fl.Field().String()); matched {

					return true
				}

			}
			return false

		})
		if err1 != nil {
			log.Fatal(err1)
		}

		err2 := v.RegisterValidation("topics", func(fl validator.FieldLevel) bool {

			topics, ok := fl.Top().Interface().(*Topics)

			if ok && topics.TopicListSize == len(topics.TopicList) {
				return true

			}
			return false

		})
		if err2 != nil {
			log.Fatal(err2)
		}

	}

	v1 := router.Group("/v1/topics") //单条帖子
	{
		v1.GET("", GetTopicList)
		v1.GET("/:topic_id", GetTopicDetail)

		v1.Use(MustLogin())
		{
			v1.POST("", NewTopic)
			v1.DELETE("/:topic_id", DelTopic)
		}
	}
	v2 := router.Group("/v1/mtopics") //多条帖子
	{
		v2.Use(MustLogin())
		{
			v2.POST("", NewTopics)

		}
	}

	router.Run()

}
