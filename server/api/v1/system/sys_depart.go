package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DepartApi struct{}

// @Tags SysDepart
// @Summary 创建代理
// @Produce  application/json
// @Param data body systemReq.Login true "代理名称，代理邮箱"
// @Success 200
// @Router /depart/createDepart [post]
func (d *DepartApi) CreateDepart(c *gin.Context) {
	var depart system.SysDepart
	_ = c.ShouldBindJSON(&depart)
	if err := utils.Verify(depart, utils.DepartVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := departService.CreateDepart(depart); err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败！", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysDepart
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDepart true "ID, 代理名称，代理邮箱"
// @Success 200
// @Router /depart/getDeartInfo [put]
func (d *DepartApi) UpdateDepart(c *gin.Context) {
	var sysDepart system.SysDepart
	_ = c.ShouldBindJSON(&sysDepart)
	if err := departService.UpdateDepart(&sysDepart); err != nil {
		global.GVA_LOG.Error("更新代理失败！", zap.Error(err))
		response.FailWithMessage("更新代理失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SysDepart
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router /depart/getDepartList [post]
func (d *DepartApi) GetDepartList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := departService.GetDepartInfoList(pageInfo); err != nil {
		// if err, list, total := departService.GetDepartInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags SysDepart
// @Summary 删除代理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "代理ID"
// @Success 200 {object} response.Response{msg=string} "删除用户"
// @Router /depart/deleteDepart [delete]
func (d *DepartApi) DeleteDepart(c *gin.Context) {
	var sysDepart system.SysDepart
	_ = c.ShouldBindJSON(&sysDepart)
	if err := departService.DeleteDepart(sysDepart); err != nil {
		global.GVA_LOG.Error("删除代理失败！", zap.Error(err))
		response.FailWithMessage("删除代理失败！", c)
	} else {
		response.OkWithMessage("删除代理成功！", c)
	}
}

// @Tags SysDepart
// @Summary 获取代理信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取代理信息"
// @Router /depart/getDepartInfo [get]
func (d *DepartApi) GetDepartInfo(c *gin.Context) {

	var depart system.SysDepart
	_ = c.ShouldBindQuery(&depart)
	if err, sysDepart := departService.GetDepartInfo(depart.ID); err != nil {
		global.GVA_LOG.Error("查询失败！", zap.Error(err))
		response.FailWithMessage("查询代理失败", c)
	} else {
		response.OkWithDetailed(gin.H{"depart": sysDepart}, "查询成功", c)
	}
}
