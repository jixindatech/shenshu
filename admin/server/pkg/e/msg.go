package e

var MsgFlags = map[int]string{
	SUCCESS:       "OK",
	ERROR:         "fail",
	InvalidParams: "请求参数错误",

	UserAddFailed:    "添加用户失败",
	UserGetFailed:    "获取用户失败",
	UserUpdateFailed: "更新用户失败",
	UserDeleteFailed: "删除用户失败",

	EmailAddFailed:    "添加邮箱失败",
	EmailGetFailed:    "获取邮箱失败",
	EmailUpdateFailed: "更新邮箱失败",

	LdapAddFailed:    "添加LDAP失败",
	LdapGetFailed:    "添加LDAP失败",
	LdapUpdateFailed: "添加LDAP失败",

	TxsmsAddFailed:    "添加Txsms失败",
	TxsmsGetFailed:    "获取Txsms失败",
	TxsmsUpdateFailed: "更新Txsms失败",

	MsgAddFailed:    "添加msg失败",
	MsgGetFailed:    "获取msg失败",
	MsgUpdateFailed: "更新msg失败",
	MsgDeleteFailed: "删除msg失败",

	SiteAddFailed:    "添加站点失败",
	SiteGetFailed:    "获取站点失败",
	SitePutFailed:    "更新站点失败",
	SiteDeleteFailed: "删除站点失败",

	SSLAddFailed:    "添加SSL失败",
	SSLGetFailed:    "添加SSL失败",
	SSLPutFailed:    "添加SSL失败",
	SSLDeleteFailed: "添加SSL失败",

	UpstreamAddFailed:    "添加Upstream失败",
	UpstreamGetFailed:    "添加Upstream失败",
	UpstreamPutFailed:    "添加Upstream失败",
	UpstreamDeleteFailed: "添加Upstream失败",

	IPAddFailed:    "添加IP失败",
	IPGetFailed:    "添加IP失败",
	IPPutFailed:    "添加IP失败",
	IPDeleteFailed: "添加IP失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
