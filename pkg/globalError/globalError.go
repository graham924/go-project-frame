package globalError

import "errors"

type GlobalError struct {
	Code             int    `json:"code"`
	Message          string `json:"message"`
	RealErrorMessage string `json:"error_msg"`
}

func (g *GlobalError) Error() string {
	return g.Message
}

// 错误码
const (
	// InternalServerError 内部错误
	InternalServerError = 10101
	// ParamBindError 参数信息有误
	ParamBindError = 10102
	// CallHTTPError 调用第三方HTTP接口有误
	CallHTTPError = 10103
	// ResubmitMsg 请勿重复提交
	ResubmitMsg = 10104

	// GetError 查询失败
	GetError = 20101
	// CreateError 添加失败
	CreateError = 20102
	// DeleteError 删除失败
	DeleteError = 20103
	// UpdateError 更新失败
	UpdateError = 20104

	// LoginError 登陆失败
	LoginError = 30101
	// LogoutError 注销失败
	LogoutError = 30102

	// AuthorizationError 签名信息token有误
	AuthorizationError = 40101
	// AuthorizationExpiredError 登陆过期，请重新登陆
	AuthorizationExpiredError = 40102
	// AuthorizationLackToken 签名信息token缺失
	AuthorizationLackToken = 40103
	// AuthorizationDeniedError 权限不足
	AuthorizationDeniedError = 40104
	// AuthorizationParseError 签名信息token解析失败
	AuthorizationParseError = 40105
)

// 错误码对应文本信息
var codeTag = map[int]string{
	InternalServerError: "Internal Server Error",
	ParamBindError:      "参数信息有误",
	CallHTTPError:       "调用第三方 HTTP 接口失败",
	ResubmitMsg:         "请勿重复提交",

	GetError:    "查询失败",
	CreateError: "添加失败",
	UpdateError: "修改失败",
	DeleteError: "删除失败",

	LoginError:  "登录失败",
	LogoutError: "注销失败",

	AuthorizationError:        "签名信息token有误",
	AuthorizationExpiredError: "登陆过期，请重新登陆",
	AuthorizationLackToken:    "请求未携带token，无权限访问",
	AuthorizationDeniedError:  "权限不足，请联系管理员",
	AuthorizationParseError:   "签名信息token解析失败",
}

func NewGlobalError(code int, err error) error {
	return &GlobalError{
		Code:             code,
		Message:          codeTag[code],
		RealErrorMessage: err.Error(),
	}
}

func GetCodeTag(code int) string {
	return codeTag[code]
}

// GetGlobalError 根据code，构建一个error出来（有些情况下，我们需要自己创建error进行返回）
func GetGlobalError(code int) error {
	return NewGlobalError(code, errors.New(GetCodeTag(code)))
}

func MustNotError(err error) {
	if err != nil {
		panic(err)
	}
}
