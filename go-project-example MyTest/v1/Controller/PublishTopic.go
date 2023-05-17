package Controller

import (
	"strconv"
	"v1/Common"
	"v1/Service"
	"v1/Util"
)

func PublishTopic(uidStr, title, content string) *Util.PageData {
	uid, _ := strconv.ParseInt(uidStr, 10, 64)

	topicId, err := Service.PublishTopic(uid, title, content)
	if err != nil {
		return Util.NewPageData(Common.CodeError, err.Error(), nil)
	}
	return Util.NewPageData(Common.CodeSuccess, "success", map[string]int64{
		"topic_id": topicId,
	})
}
