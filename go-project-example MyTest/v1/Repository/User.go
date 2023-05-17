package Repository

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"v1/Common"
)

type User struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Avatar     string    `json:"avatar"`
	CreateTime time.Time `json:"create_time"`
}

type UserDao struct {
}

var (
	userDao  *UserDao
	userOnce sync.Once // 单例模式 ，只执行一次 ， 适合高并发 ， 减少存储的浪费
)

func NewUserDaoInstance() *UserDao {
	userOnce.Do( // 只有第一次调用有效 ， 适合初始化 ， 安全可靠
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) QueryUserById(id int64) (*User, error) {
	if UserIndexMap[id] == nil {
		return nil, errors.New("QueryUserById : User not exit")
	}
	return UserIndexMap[id], nil
}

func (*UserDao) MQueryUserById(ids []int64) (map[int64]*User, error) {
	users := make([]*User, 0)
	for _, id := range ids {
		fmt.Printf("id : ? %d\n", id)
		users = append(users, UserIndexMap[int64(id)])
	}
	if len(users) == 0 {
		return nil, errors.New("MQueryUserById : Users not exit")
	}
	userMap := make(map[int64]*User)
	for _, user := range users {
		userMap[user.Id] = user
	}
	return userMap, nil
}

func (*UserDao) CreateUser(user *User) error {
	// 实现 Id 自增
	uid++
	user.Id = uid
	// 检验是否存在
	if UserIndexMap[user.Id] != nil {
		return errors.New("CreateUser : user has exit")
	}
	// 存储新数据
	if err := StoreOne(Common.UserFile, user); err != nil {
		return err
	}
	// 记录
	UserIndexMap[user.Id] = user
	return nil
}
