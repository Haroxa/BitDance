package Common

import (
	"log"
	"time"
)

const (
	CodeError   = -1
	CodeSuccess = 0
)

const (
	DataPath  = "./Data/"
	UserFile  = "user.txt"
	TopicFile = "topic.txt"
	PostFile  = "post.txt"
)

var ChinaTime *time.Location

func init() {
	var err error
	ChinaTime, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Panicln(err)
	}
}
