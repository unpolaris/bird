package errcode

var ErrInvalidParams = NewErr(InvalidParams, "请求参数有误")

var (
	NoErr = NewErr(20000, "OK")
)

var (
	ErrBirdCreate = NewErr(40200, "添加鸟类失败")
	ErrBirdUpdate = NewErr(40201, "更新鸟类失败")
	ErrBirdList   = NewErr(40202, "获取鸟类列表失败")
	ErrBirdInfo   = NewErr(40203, "获取鸟类详情失败")
	ErrBirdDel    = NewErr(40204, "删除鸟类失败")
	ErrBirdSearch = NewErr(40205, "搜索鸟类失败")
)

var (
	ErrBirdRedisGet      = NewErr(40300, "获取鸟类列表失败")
	ErrBirdMarshalString = NewErr(40301, "序列化鸟类列表失败")
)
