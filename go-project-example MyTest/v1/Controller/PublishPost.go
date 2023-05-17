package Controller

import (
	"strconv"
	"v1/Common"
	"v1/Service"
	"v1/Util"
)

func PublishPost(uidStr, topicIdStr, content string) *Util.PageData {
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	topicId, _ := strconv.ParseInt(topicIdStr, 10, 64)

	postId, err := Service.PublishPost(topicId, uid, content)
	if err != nil {
		return Util.NewPageData(Common.CodeError, err.Error(), nil)
	}
	return Util.NewPageData(Common.CodeSuccess, "success", map[string]int64{
		"post_id": postId,
	})
}
