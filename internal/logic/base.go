package logic

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"template/internal/model"
)

type BaseLogic struct {
}

func (b *BaseLogic) GetUserInfo(ctx *gin.Context) model.UserTokenInfo {
	userData := ctx.MustGet("user_info")
	if userData == "" {
		return model.UserTokenInfo{}
	}
	userInfo := model.UserTokenInfo{}
	_ = jsoniter.UnmarshalFromString(userData.(string), &userInfo)
	return userInfo
}
