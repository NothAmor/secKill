package secgin

type SecError struct {
	Code int
	Msg  string
}

var (
	//参数校验
	PARAM_MISSING     SecError = SecError{10001, "参数校验缺失"}
	PARAM_ERROR       SecError = SecError{10002, "参数校验错误"}
	PARAM_USER_MISSIG SecError = SecError{10101, "用户名缺失"}
	PARAM_USER_ERROR  SecError = SecError{10102, "用户名错误"}
	PARAM_MOBILEPHONE SecError = SecError{10200, "手机号错误"}
	PARAM_TELEPHONE   SecError = SecError{10300, "电话错误"}
	PARAM_EMAIL       SecError = SecError{10400, "邮箱错误"}

	SYS_DEFAULT SecError = SecError{50000, "系统错误"}
)
