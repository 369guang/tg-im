package errno

// 定义错误码
// 1000 - 1999 通用错误
// 2000 - 2999 JSON错误
// 3000 - 3999 RPC错误

// JSON和RPC错误码
const (
	Success        = 0    // 成功
	ErrCodeUnknown = 1000 // 未知错误
	ErrCodeParam   = 1001 // 参数错误
	ErrCodeDB      = 1002 // 数据库错误
	ErrCodeRedis   = 1003 // Redis错误
	ErrCodeRPC     = 1004 // RPC错误
	ErrCodeJSON    = 1005 // JSON错误

	ErrCodeAuth = 2000 // 认证错误
	ErrCodeUser = 2001 // 用户错误
	ErrCodeRole = 2002 // 角色错误
	ErrCodePerm = 2003 // 权限错误
	ErrCodeMenu = 2004 // 菜单错误
	ErrCodeFile = 2005 // 文件错误
	ErrCodeLog  = 2006 // 日志错误
	ErrCodeTask = 2007 // 任务错误

	ErrorRPCAuth = 3000 // RPC认证错误
)

// ErrMsg 错误码对应的错误信息
var ErrMsg = map[int]string{
	Success:        "成功",
	ErrCodeUnknown: "未知错误",
	ErrCodeParam:   "参数错误",
	ErrCodeDB:      "数据库错误",
	ErrCodeRedis:   "Redis错误",
	ErrCodeRPC:     "RPC错误",
	ErrCodeJSON:    "JSON错误",

	ErrCodeAuth: "认证错误",
	ErrCodeUser: "用户错误",
	ErrCodeRole: "角色错误",
	ErrCodePerm: "权限错误",
	ErrCodeMenu: "菜单错误",
	ErrCodeFile: "文件错误",
	ErrCodeLog:  "日志错误",
	ErrCodeTask: "任务错误",

	ErrorRPCAuth: "RPC认证错误",
}
