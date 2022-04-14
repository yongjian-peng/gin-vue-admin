package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SysDepart struct {
	global.GVA_MODEL        // 用户UUID
	Type             string `json:"type" gorm:"comment:代理类型"`                   // 用户登录密码
	CompanyName      string `json:"companyName" gorm:"default:'';comment:代理名称"` // 用户昵称
	Status           int8   `json:"status" gorm:"default:0;comment:代理状态"`       // 用户侧边主题
	CompanyEmail     string `json:"companyEmail" gorm:"default:-;comment:代理邮箱"` // 用户头像
	CreateTime       int64  `json:"createTime" gorm:"default:0;comment:添加时间"`   // 添加时间
	UpdateTime       int64  `json:"updateTime" gorm:"default:0;comment:修改时间"`   // 添加时间
}

func (SysDepart) TableName() string {
	return "sys_depart"
}
