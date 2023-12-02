package consts

type SysUserAuthorityId uint

const (
	SystemUserAuthorityIdAdmin      SysUserAuthorityId = 111
	SystemUserAuthorityIdDefault    SysUserAuthorityId = 222
	SystemUserAuthorityIdSubDefault SysUserAuthorityId = 2221
)

const (
	LoginURL    = "/api/user/login"
	LogoutURL   = "/api/user/logout"
	WebShellURL = "/api/k8s/pod/webshell"
)

const (
	ValidatorContextKey  = "ValidatorContextKey"
	TranslatorContextKey = "TranslatorContextKey"
	ClaimsContextKey     = "ClaimsContextKey"
)

const (
	TokenKey = "token"
)
