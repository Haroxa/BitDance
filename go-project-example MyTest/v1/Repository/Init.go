package Repository

import (
	"bufio"
	"encoding/json"
	"os"
	"v1/Common"
	"v1/Util"
)

var UserIndexMap map[int64]*User
var TopicIndexMap map[int64]*Topic
var postIndexMap map[int64][]*Post

var (
	uid int64
	tid int64
	pid int64
)

func InitIndex() error {
	if err := InitUserIndexMap(Common.DataPath); err != nil {
		return err
	}
	if err := InitTopicIndexMap(Common.DataPath); err != nil {
		return err
	}
	if err := InitPostIndexMap(Common.DataPath); err != nil {
		return err
	}
	return nil
}

func InitUserIndexMap(filePath string) error {
	// 打开文件
	open, err := os.Open(filePath + Common.UserFile)
	if err != nil {
		return err
	}
	// 创建读取器
	scanner := bufio.NewScanner(open)
	UserTmpMap := make(map[int64]*User)
	// 读取直至文件尾
	for scanner.Scan() {
		text := scanner.Text()
		var user User
		// string to User struct
		if err = json.Unmarshal([]byte(text), &user); err != nil {
			return err
		}
		// 记录数据
		uid = Util.MaxInt64(uid, user.Id)
		UserTmpMap[user.Id] = &user
	}
	UserIndexMap = UserTmpMap
	open.Close()
	return nil
}

func InitTopicIndexMap(filePath string) error {
	// 打开文件
	open, err := os.Open(filePath + Common.TopicFile)
	if err != nil {
		return err
	}
	// 创建读取器
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	// 读取直至文件尾
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		// string to Topic struct
		if err = json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		// 记录数据
		tid = Util.MaxInt64(tid, topic.Id)
		topicTmpMap[topic.Id] = &topic // 局部变量会一直到其变得无法访问才会被系统回收 //https://blog.csdn.net/fx677588/article/details/78421334
	}
	TopicIndexMap = topicTmpMap
	open.Close()
	return nil
}

func InitPostIndexMap(filePath string) error {
	// 打开文件
	open, err := os.Open(filePath + Common.PostFile)
	if err != nil {
		return err
	}
	// 创建读取器
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)
	// 读取直至文件尾
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		// string to Post struct
		if err = json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		// 记录数据
		pid = Util.MaxInt64(pid, post.Id)
		postTmpMap[post.TopicId] = append(postTmpMap[post.TopicId], &post)
	}
	postIndexMap = postTmpMap
	open.Close()
	return nil
}
