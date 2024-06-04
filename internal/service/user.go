package service

import (
	"github.com/gin-gonic/gin"
	"github.com/senyu-up/toolbox/tool/http/gin_server/controller"
	"template/internal/logic"
	"template/internal/model"
)

var UserController = new(user)

type user struct {
	controller.BaseController
}

//// UserRegister
//// @tags 用户管理
//// @summary 用户注册
//// @param req body model.UserRegisterParams true "json入参"
//// @success 200 {object} CommonResp
//// @router /xhs/user/register [post]
//func (u *user) UserRegister(ctx *gin.Context) {
//	params := new(model.UserRegisterParams)
//	u.call_(ctx, params, logic.UserLogic.UserRegister)
//}

// UserLogin
// @tags 用户管理
// @summary 用户登录
// @param req body model.UserLoginParams true "json入参"
// @success 200 {object} model.UserLoginResp
// @router /user/login [post]
func (u *user) UserLogin(ctx *gin.Context) {
	params := new(model.UserLoginParams)
	u.Call_(ctx, params, logic.UserLogic.UserLogin)
}

// UserSearch
// @tags 用户管理
// @summary 用户查询
// @param req body model.UserSearchParams true "json入参"
// @success 200 {object} model.User
// @router /user/search [post]
func (u *user) UserSearch(ctx *gin.Context) {
	params := new(model.UserSearchParams)
	u.Call_(ctx, params, logic.UserLogic.UserSearch)
}
