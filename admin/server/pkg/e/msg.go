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
	SSLGetFailed:    "获取SSL失败",
	SSLPutFailed:    "更新SSL失败",
	SSLDeleteFailed: "删除SSL失败",

	UpstreamAddFailed:    "添加Upstream失败",
	UpstreamGetFailed:    "获取Upstream失败",
	UpstreamPutFailed:    "更新Upstream失败",
	UpstreamDeleteFailed: "删除Upstream失败",

	IPAddFailed:    "添加IP失败",
	IPGetFailed:    "获取IP失败",
	IPPutFailed:    "更新IP失败",
	IPDeleteFailed: "删除IP失败",

	CCAddFailed:    "添加CC失败",
	CCGetFailed:    "获取CC失败",
	CCPutFailed:    "更新CC失败",
	CCDeleteFailed: "删除CC失败",

	RuleGroupAddFailed:    "添加RuleGroup失败",
	RuleGroupGetFailed:    "获取RuleGroup失败",
	RuleGroupPutFailed:    "更新RuleGroup失败",
	RuleGroupDeleteFailed: "删除RuleGroup失败",

	RuleAddFailed:    "添加Rule失败",
	RuleGetFailed:    "获取Rule失败",
	RulePutFailed:    "更新Rule失败",
	RuleDeleteFailed: "删除Rule失败",

	RuleBatchAddFailed:    "添加RuleBatch失败",
	RuleBatchGetFailed:    "获取RuleBatch失败",
	RuleBatchPutFailed:    "更新RuleBatch失败",
	RuleBatchDeleteFailed: "删除RuleBatch失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
