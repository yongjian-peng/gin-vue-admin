package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DepartRouter struct{}

func (s *DepartRouter) InitDepartRouter(Router *gin.RouterGroup) {
	departRouter := Router.Group("depart").Use(middleware.OperationRecord())
	departRouterWithoutRecord := Router.Group("depart")
	departApi := v1.ApiGroupApp.SystemApiGroup.DepartApi
	{
		departRouter.POST("createDepart", departApi.CreateDepart)   // 管理员注册账号
		departRouter.PUT("updateDepart", departApi.UpdateDepart)    // 用户修改密码
		departRouter.DELETE("deleteDepart", departApi.DeleteDepart) // 删除用户
	}
	{
		departRouterWithoutRecord.POST("getDepartList", departApi.GetDepartList) // 分页获取代理列表
		departRouterWithoutRecord.GET("getDepartInfo", departApi.GetDepartInfo)  // 获取代理信息
	}
}
