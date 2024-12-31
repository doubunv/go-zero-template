package consts

type IntBool int
type LogType int

const (
	HeaderToken = "AuthorizationJwt"
)

const (
	Trace        = "trace"
	ClientIp     = "client_ip"
	UserAgent    = "user_agent"
	Token        = "token"
	TokenUid     = "token_uid"
	TokenUidRole = "token_uid_role"
	Version      = "version"
	Source       = "source"
	ReqPath      = "req_path"
)
