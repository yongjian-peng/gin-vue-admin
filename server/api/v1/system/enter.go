package system

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DBApi
	JwtApi
	BaseApi
	DepartApi
	SystemApi
	CasbinApi
	AutoCodeApi
	SystemApiApi
	AuthorityApi
	DictionaryApi
	AuthorityMenuApi
	OperationRecordApi
	AutoCodeHistoryApi
	DictionaryDetailApi
	AuthorityBtnApi
}

var (
	apiService              = service.ServiceGroupApp.SystemServiceGroup.ApiService
	jwtService              = service.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService             = service.ServiceGroupApp.SystemServiceGroup.MenuService
	userService             = service.ServiceGroupApp.SystemServiceGroup.UserService
	departService           = service.ServiceGroupApp.SystemServiceGroup.DepartService
	initDBService           = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	casbinService           = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	autoCodeService         = service.ServiceGroupApp.SystemServiceGroup.AutoCodeService
	baseMenuService         = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
	authorityService        = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	dictionaryService       = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
	systemConfigService     = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	operationRecordService  = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
	autoCodeHistoryService  = service.ServiceGroupApp.SystemServiceGroup.AutoCodeHistoryService
	dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
	authorityBtnService     = service.ServiceGroupApp.SystemServiceGroup.AuthorityBtnService
)
