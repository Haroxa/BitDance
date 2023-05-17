package Service

import (
	"errors"
	"time"
	"unicode/utf8"
	"v1/Common"
	"v1/Repository"
)

type PublishTopicFlow struct {
	userId  int64
	title   string
	content string

	topicId int64
}

func PublishTopic(userId int64, title, content string) (int64, error) {
	return NewPublishTopicFlow(userId, title, content).Do()
}

func NewPublishTopicFlow(userId int64, title, content string) *PublishTopicFlow {
	return &PublishTopicFlow{
		userId:  userId,
		title:   title,
		content: content,
	}
}

func (f *PublishTopicFlow) Do() (int64, error) {
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	if err := f.publish(); err != nil {
		return 0, err
	}
	return f.topicId, nil
}

func (f *PublishTopicFlow) checkParam() error {
	if Repository.UserIndexMap[f.userId] == nil {
		return errors.New("user has no exit")
	}
	if utf8.RuneCountInString(f.title) >= 50 {
		return errors.New("title length must be less than 50")
	}
	if utf8.RuneCountInString(f.content) >= 500 {
		return errors.New("content length must be less than 500")
	}
	return nil
}

func (f *PublishTopicFlow) publish() error {
	topic := &Repository.Topic{
		UserId:     f.userId,
		Title:      f.title,
		Content:    f.content,
		CreateTime: time.Now().In(Common.ChinaTime),
	}
	if err := Repository.NewTopicDaoInstance().CreateTopic(topic); err != nil {
		return err
	}
	f.topicId = topic.Id
	return nil
}
