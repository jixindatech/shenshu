package e

const (
	SUCCESS       = 0
	ERROR         = 500
	InvalidParams = 400

	UserAddFailed     = 10000
	UserGetFailed     = 10001
	UserUpdateFailed  = 10002
	UserDeleteFailed  = 10003
	EmailAddFailed    = 20000
	EmailGetFailed    = 20001
	EmailUpdateFailed = 20002
	LdapAddFailed     = 30000
	LdapGetFailed     = 30001
	LdapUpdateFailed  = 30002
	TxsmsAddFailed    = 40000
	TxsmsGetFailed    = 40001
	TxsmsUpdateFailed = 40002
	MsgAddFailed      = 50000
	MsgGetFailed      = 50001
	MsgUpdateFailed   = 50002
	MsgDeleteFailed   = 50003

	SiteAddFailed          = 60001
	SiteGetFailed          = 60002
	SitePutFailed          = 60003
	SiteDeleteFailed       = 60004
	SiteGetRuleGroupFailed = 60005
	SitePutRuleGroupFailed = 60006
	SiteEnableSiteConfig   = 60007

	SSLAddFailed    = 70001
	SSLGetFailed    = 70002
	SSLPutFailed    = 70003
	SSLDeleteFailed = 70004

	UpstreamAddFailed    = 80001
	UpstreamGetFailed    = 80002
	UpstreamPutFailed    = 80003
	UpstreamDeleteFailed = 80004

	IPAddFailed    = 90001
	IPGetFailed    = 90002
	IPPutFailed    = 90003
	IPDeleteFailed = 90004

	CCAddFailed    = 11001
	CCGetFailed    = 11002
	CCPutFailed    = 11003
	CCDeleteFailed = 11004

	BatchGroupAddFailed    = 12001
	BatchGroupGetFailed    = 12002
	BatchGroupPutFailed    = 12003
	BatchGroupDeleteFailed = 12004

	SpecificGroupAddFailed    = 15001
	SpecificGroupGetFailed    = 15002
	SpecificGroupPutFailed    = 15003
	SpecificGroupDeleteFailed = 15004

	RuleSpecificAddFailed    = 13001
	RuleSpecificGetFailed    = 13002
	RuleSpecificPutFailed    = 13003
	RuleSpecificDeleteFailed = 13004

	RuleBatchAddFailed    = 14001
	RuleBatchGetFailed    = 14002
	RuleBatchPutFailed    = 14003
	RuleBatchDeleteFailed = 14004

	CCEventGetFailed           = 16001
	SpecificRuleEventGetFailed = 16002
	BatchRuleEventGetFailed    = 16003
	SiteRuleEventGetFailed     = 16004
)
