package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/senyu-up/toolbox/tool/encrypt"
	"github.com/senyu-up/toolbox/tool/struct_tool"
	"github.com/senyu-up/toolbox/tool/su_logger"
	"template/global"
	"template/internal/dao"
	"template/internal/model"
	"time"
)

var UserLogic = new(user)

type user struct {
	BaseLogic
}

func (u *user) UserLogin(ctx *gin.Context, params *model.UserLoginParams) (data interface{}, err error) {
	tx := global.GetFacade().GetMysqlClient().WithContext(ctx)
	go func() {
		su_logger.Info(ctx, "111")
	}()
	su_logger.Info(ctx, "222")
	panic("1111")
	//查找唯一标识符是否存在
	userInfo, err := dao.UserDao.UserSearchOneByCond(tx, map[string]interface{}{"identifier": params.Identifier})
	userData := model.UserTokenInfo{
		UserId:     userInfo.UserId,
		Username:   userInfo.Username,
		Phone:      userInfo.Phone,
		Identifier: userInfo.Identifier,
	}
	//不存在则创建后登录
	if userInfo.Id == 0 || userInfo.Identifier == "" {
		nowTime := time.Now().Unix()
		userCreateData := &model.User{
			Identifier: params.Identifier,
			InviterId:  params.InvitationCode,
			UpdatedAt:  nowTime,
			CreatedAt:  nowTime,
		}
		if err = tx.Create(userCreateData).Error; err != nil {
			tx.Rollback()
			return
		}
		uid := fmt.Sprintf("xhs_1%06d", userCreateData.Id)
		err = dao.UserDao.UserUpdateByCond(tx, map[string]interface{}{"id": userCreateData.Id}, map[string]interface{}{"user_id": uid, "username": uid})
		if err != nil {
			tx.Rollback()
			return
		}
		userData.UserId = uid
		userData.Username = uid
		userData.Identifier = params.Identifier
	}

	sUserData, _ := jsoniter.MarshalToString(userData)
	token, err := encrypt.CreateToken(sUserData, 7, global.GetConfig().Jwt.TokenSecret)
	if err != nil {
		return
	}
	resp := &model.UserLoginResp{UserId: userData.UserId, Username: userData.Username, Token: token}
	return resp, nil
}

func (u *user) UserSearch(ctx *gin.Context, params *model.UserSearchParams) (data interface{}, err error) {
	//当用户id和电话号码为空时，则使用token用户信息
	if params.UserId == "" && params.Phone == "" && params.Identifier == "" {
		params.UserId = u.GetUserInfo(ctx).UserId
	}
	tx := global.GetFacade().GetMysqlClient()
	userInfoFromDB := new(model.User)
	if params.UserId != "" {
		userInfoFromDB, err = dao.UserDao.UserSearchOneByCond(tx, map[string]interface{}{"user_id": params.UserId})
		if err != nil {
			return nil, err
		}
		if userInfoFromDB.Id == 0 || userInfoFromDB.UserId == "" {
			return nil, fmt.Errorf("用户不存在")
		}
	}
	if params.Phone != "" {
		userInfoFromDB, err = dao.UserDao.UserSearchOneByCond(tx, map[string]interface{}{"phone": params.Phone})
		if err != nil {
			return nil, err
		}
		if userInfoFromDB.Id == 0 || userInfoFromDB.UserId == "" {
			return nil, fmt.Errorf("用户不存在")
		}
	}
	if params.Identifier != "" {
		userInfoFromDB, err = dao.UserDao.UserSearchOneByCond(tx, map[string]interface{}{"identifier": params.Identifier})
		if err != nil {
			return nil, err
		}
		if userInfoFromDB.Id == 0 || userInfoFromDB.UserId == "" {
			return nil, fmt.Errorf("用户不存在")
		}
	}
	respData := new(model.UserSearchResp)
	_ = struct_tool.DeepCopy(respData, userInfoFromDB)

	return respData, nil
}
