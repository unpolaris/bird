package errcode

var ErrInvalidParams = NewErr(InvalidParams, "请求参数有误")

var (
	NoErr               = NewErr(20000, "OK")
	ErrUserLogin        = NewErr(30000, "登录信息过期，请重新登录")
	ErrRefreshToken     = NewErr(30001, "refresh_token 无效")
	Frequently          = NewErr(40100, "请勿频繁操作")
	RowNotExist         = NewErr(40101, "信息不存在")
	VerificationCodeErr = NewErr(40102, "验证码已失效或者有误")
	UserAlreadyExists   = NewErr(40103, "用户已经存在")
	OcrToIdNumber       = NewErr(40104, "解析身份证失败")
	OcrToIdNumberResult = NewErr(40105, "解析身份证结果失败")
	UserNotExist        = NewErr(40106, "用户信息不存在")
)

var (
	UserIdCannotBeEmpty = NewErr(40200, "用户id不能为空")
	UserNotSubscribe    = NewErr(40201, "用户还没有订阅")
)

var (
	PayNoAlreadyPayed  = NewErr(40300, "订单已支付")
	PayNoAlreadyClosed = NewErr(40301, "订单已关闭")
	PayNoAlreadyRefund = NewErr(40302, "订单已退款")
)
