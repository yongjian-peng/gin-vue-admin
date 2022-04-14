package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type SysDepartResponse struct {
	Depart system.SysDepart `json:"depart"`
}

// type LoginResponse struct {
// 	User      system.SysUser `json:"user"`
// 	Token     string         `json:"token"`
// 	ExpiresAt int64          `json:"expiresAt"`
// }
