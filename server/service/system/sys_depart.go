package system

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type DepartService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDepart
//@description: 创建代理
//@param: depart model.SysDepart
//@return: err error

func (departService *DepartService) CreateDepart(depart system.SysDepart) (err error) {
	var sysDepart system.SysDepart
	if !errors.Is(global.GVA_DB.Where("company_name = ?", depart.CompanyName).First(&sysDepart).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同的代理名称")
	}
	return global.GVA_DB.Create(&depart).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDepartInfoList
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

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDepart
//@description: 更新代理
//@param: depart *model.SysDepart
//@return: err error

func (departService *DepartService) UpdateDepart(depart *system.SysDepart) (err error) {
	var sysDepart system.SysDepart
	sysDepartMap := map[string]interface{}{
		"CompanyName":  depart.CompanyName,
		"CompanyEmail": depart.CompanyEmail,
		"Type":         depart.Type,
	}

	db := global.GVA_DB.Where("id = ?", depart.ID).First(&sysDepart)

	return db.Updates(sysDepartMap).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDepart
//@description: 根据id或者type获取字典单条数据
//@param: Type string, Id uint
//@return: err error, sysDepart model.SysDepart

func (departService *DepartService) GetDepartInfo(Id uint) (err error, depart system.SysDepart) {
	var sysDepart system.SysDepart
	err = global.GVA_DB.Where("id = ?", Id).First(&sysDepart).Error
	return err, sysDepart
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDepart
//@description: 删除代理
//@param: sysDepart model.SysDepart
//@return: err error
func (departService *DepartService) DeleteDepart(depart system.SysDepart) (err error) {
	var sysDepart system.SysDepart
	err = global.GVA_DB.Where("id = ?", depart.ID).First(&sysDepart).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("请不要搞事情")
	}
	if err != nil {
		return err
	}

	err = global.GVA_DB.Delete(&sysDepart).Error
	if err != nil {
		return err
	}
	return nil
}
