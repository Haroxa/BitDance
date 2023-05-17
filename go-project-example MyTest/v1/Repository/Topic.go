package Repository

import (
	"errors"
	"sync"
	"time"
	"v1/Common"
)

type Topic struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"user_id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
}

type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once // 单例模式 ，只执行一次 ， 适合高并发 ， 减少存储的浪费
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do( // 只有第一次调用有效 ， 适合初始化 ， 安全可靠
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

func (*TopicDao) QueryTopicById(id int64) (*Topic, error) {
	if TopicIndexMap[id] == nil {
		return nil, errors.New("QueryTopicById : topic not exit")
	}
	return TopicIndexMap[id], nil
}

func (*TopicDao) CreateTopic(topic *Topic) error {
	// 实现 Id 自增
	tid++
	topic.Id = tid
	// 检验是否存在
	if TopicIndexMap[topic.Id] != nil {
		return errors.New("CreateTopic : topic has exit")
	}
	// 存储新数据
	if err := StoreOne(Common.TopicFile, topic); err != nil {
		return err
	}
	// 记录
	TopicIndexMap[topic.Id] = topic
	return nil
}
