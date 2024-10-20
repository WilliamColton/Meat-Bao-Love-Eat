package response

// 通用码
var (
	ERROR   = NewError(0, "失败")
	Success = NewError(1, "成功")
)

// 资源码
var (
	NotFound = NewError(10000, "找不到")
)

// 鉴权码
var (
	UnauthorizedTokenError    = NewError(20000, "鉴权失败，Token错误")
	UnauthorizedTokenGenerate = NewError(20001, "鉴权失败，Token生成失败")
)
