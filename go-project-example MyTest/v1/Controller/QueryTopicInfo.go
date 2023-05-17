package Controller

import (
	"strconv"
	"v1/Common"
	"v1/Service"
	"v1/Util"
)

func QueryTopicInfo(topicIdStr string) *Util.PageData {
	topicId, err := strconv.ParseInt(topicIdStr, 10, 64)
	if err != nil {
		return Util.NewPageData(Common.CodeError, err.Error(), nil)
	}
	pageInfo, err := Service.QueryPageInfo(topicId)
	if err != nil {
		return Util.NewPageData(Common.CodeError, err.Error(), nil)
	}
	return Util.NewPageData(Common.CodeSuccess, "success", pageInfo)
}
