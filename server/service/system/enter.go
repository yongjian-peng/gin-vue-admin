package system

type ServiceGroup struct {
	JwtService
	ApiService
	MenuService
	UserService
	DepartService
	CasbinService
	InitDBService
	AutoCodeService
	BaseMenuService
	AuthorityService
	DictionaryService
	SystemConfigService
	AutoCodeHistoryService
	OperationRecordService
	DictionaryDetailService
	AuthorityBtnService
}
