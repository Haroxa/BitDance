package Controller

import (
	"v1/Common"
	"v1/Service"
	"v1/Util"
)

func Register(name, avatar string) *Util.PageData {
	userId, err := Service.PublishUser(name, avatar)
	if err != nil {
		return Util.NewPageData(Common.CodeError, err.Error(), nil)
	}
	return Util.NewPageData(Common.CodeSuccess, "success", map[string]int64{
		"user_id": userId,
	})
}
