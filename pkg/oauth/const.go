package oauth

const (
	GRANT_TYPE_AUTHORIZATION_CODE = "authorization_code"
	GRANT_TYPE_PASSWORD           = "password"
	GRANT_TYPE_CLIENT_CREDENTIALS = "client_credentials"
	GRANT_TYPE_IMPLICIT           = "implicit"
	GRANT_TYPE_TOKEN              = "token"
	GRANT_TYPE_REFRESH_TOKEN      = "refresh_token"
)

type Api struct {
	Url string
}

var (
	Authorize     = Api{Url: "/oauth2/authorize"}
	VerifyAccount = Api{Url: "/verify/verify_account"}
	Login         = Api{Url: "/user/login"}
	AccessToken   = Api{Url: "/oauth2/token"}
	RefreshToken  = Api{Url: "/oauth2/refresh"}
	UserInfo      = Api{Url: "/oauth2/userinfo"}
)

func (api Api) GetHttp(hostName string) string {
	return "https://" + hostName + api.Url
}
