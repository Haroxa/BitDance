package Repository

import (
	"errors"
	"sync"
	"time"
	"v1/Common"
)

type Post struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"user_id"`
	TopicId    int64     `json:"topic_id"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
}

type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once // 单例模式 ，只执行一次 ， 适合高并发 ， 减少存储的浪费
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do( // 只有第一次调用有效 ， 适合初始化 ， 安全可靠
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

func (*PostDao) QueryPostById(id int64) ([]*Post, error) {
	if postIndexMap[id] == nil {
		return nil, errors.New("QueryPostById : post not exit")
	}
	return postIndexMap[id], nil
}

func (*PostDao) QueryPostByTopicId(topicId int64) ([]*Post, error) {
	var posts []*Post
	posts = postIndexMap[topicId]
	if posts == nil {
		return nil, errors.New("QueryPostByTopicId : posts not exit")
	}
	return posts, nil
}

func (*PostDao) CreatePost(post *Post) error {
	// 实现 Id 自增
	pid++
	post.Id = pid
	// 存储新数据
	if err := StoreOne(Common.PostFile, post); err != nil {
		return err
	}
	// 记录
	postIndexMap[post.TopicId] = append(postIndexMap[post.TopicId], post)
	return nil
}
