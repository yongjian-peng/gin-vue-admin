package system

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: err error, userInter model.SysUser

type DepartService struct{}

func (departService *DepartService) CreateDepart(depart system.SysDepart) (err error) {
	var sysDepart system.SysDepart
	if !errors.Is(global.GVA_DB.Where("company_name = ?", depart.CompanyName).First(&sysDepart).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同的代理名称")
	}
	return global.GVA_DB.Create(&depart).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (departService *DepartService) GetDepartInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysDepart{})
	var departList []system.SysDepart
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&departList).Error
	return err, departList, total
}
