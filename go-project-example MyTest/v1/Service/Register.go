package Service

import (
	"errors"
	"time"
	"unicode/utf8"
	"v1/Common"
	"v1/Repository"
)

type PublishUserFlow struct {
	name   string
	avatar string

	userId int64
}

func PublishUser(name, avatar string) (int64, error) {
	return NewPublishUserFlow(name, avatar).Do()
}

func NewPublishUserFlow(name, avatar string) *PublishUserFlow {
	return &PublishUserFlow{
		name:   name,
		avatar: avatar,
	}
}

func (f *PublishUserFlow) Do() (int64, error) {
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	if err := f.publish(); err != nil {
		return 0, err
	}
	return f.userId, nil
}

func (f *PublishUserFlow) checkParam() error {
	if utf8.RuneCountInString(f.name) >= 20 {
		return errors.New("content length must be less than 20")
	}
	return nil
}

func (f *PublishUserFlow) publish() error {
	user := &Repository.User{
		Name:       f.name,
		Avatar:     f.avatar,
		CreateTime: time.Now().In(Common.ChinaTime),
	}
	if err := Repository.NewUserDaoInstance().CreateUser(user); err != nil {
		return err
	}
	f.userId = user.Id
	return nil
}
