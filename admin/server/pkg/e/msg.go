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
	TxsmsGetFailed:    "添加Txsms失败",
	TxsmsUpdateFailed: "添加Txsms失败",

	MsgAddFailed:    "添加msg失败",
	MsgGetFailed:    "添加msg失败",
	MsgUpdateFailed: "添加msg失败",
	MsgDeleteFailed: "添加msg失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
