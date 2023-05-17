package Service

import (
	"errors"
	"time"
	"unicode/utf8"
	"v1/Common"
	"v1/Repository"
)

type PublishPostFlow struct {
	userId  int64
	content string
	topicId int64

	postId int64
}

func PublishPost(topicId, userId int64, content string) (int64, error) {
	return NewPublishPostFlow(topicId, userId, content).Do()
}

func NewPublishPostFlow(topicId, userId int64, content string) *PublishPostFlow {
	return &PublishPostFlow{
		userId:  userId,
		content: content,
		topicId: topicId,
	}
}

func (f *PublishPostFlow) Do() (int64, error) {
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	if err := f.publish(); err != nil {
		return 0, err
	}
	return f.postId, nil
}

func (f *PublishPostFlow) checkParam() error {
	if Repository.UserIndexMap[f.userId] == nil {
		return errors.New("user has no exit")
	}
	if Repository.TopicIndexMap[f.topicId] == nil {
		return errors.New("topic has no exit")
	}
	if utf8.RuneCountInString(f.content) >= 500 {
		return errors.New("content length must be less than 500")
	}
	return nil
}

func (f *PublishPostFlow) publish() error {
	post := &Repository.Post{
		TopicId:    f.topicId,
		UserId:     f.userId,
		Content:    f.content,
		CreateTime: time.Now().In(Common.ChinaTime),
	}
	if err := Repository.NewPostDaoInstance().CreatePost(post); err != nil {
		return err
	}
	f.postId = post.Id
	return nil
}
