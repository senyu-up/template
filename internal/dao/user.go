package dao

import (
	"gorm.io/gorm"
	"template/internal/model"
)

var UserDao = new(user)

type user struct {
}

// UserSearchOneByCond 通用用户信息查询
func (u *user) UserSearchOneByCond(tx *gorm.DB, cond map[string]interface{}, fields ...string) (data *model.User, err error) {
	qs := tx.Model(&model.User{})
	for filed, value := range cond {
		qs = qs.Where(filed, value)
	}
	if len(fields) > 0 {
		qs.Select(fields)
	}
	err = qs.Find(&data).Error
	return
}

// UserUpdateByCond 通用用户信息更新
func (u *user) UserUpdateByCond(tx *gorm.DB, cond map[string]interface{}, data interface{}) error {
	tx = tx.Model(&model.User{})
	for filed, value := range cond {
		tx = tx.Where(filed, value)
	}
	return tx.Updates(data).Error
}
