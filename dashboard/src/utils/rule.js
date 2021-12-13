const ACTION_TYPES = [
  { label: '允许', value: 1 },
  { label: '阻断', value: 2 },
  { label: '日志', value: 4 }
]
const VARIABLES_TEXT = {
  URI: 'URI',
  METHOD: '请求方法',
  REQ_HEADER: '请求头',
  IP: 'IP',
  FILE: '上传文件',
  FILE_NAMES: '上传文件名称'
}

const OPERATORS_TEXT = {
  not_exist: '不存在',
  ip_contains: '属于',
  not_ip_contains: '不属于',
  str_equal: '等于',
  not_str_equal: '不等于',
  str_contains: '包含',
  not_str_contains: '不包含',
  prefix_equal: '前缀匹配',
  suffix_equal: '后缀匹配',
  num_gt: '数值大于',
  num_le: '数值小于',
  num_equal: '数值等于',
  re_equal: '正则匹配'
}
const IP_OPERATORS = [
  { value: 'ip_contains', label: '属于' },
  { value: 'not_ip_contains', label: '不属于' }
]
const METHOD_OPERATORS = [
  { value: 'str_equal', label: '等于' },
  { value: 'not_str_equal', label: '不等于' }
]
const URI_OPERATORS = [
  { value: 'str_equal', label: '等于' },
  { value: 'prefix_equal', label: '前缀匹配' },
  { value: 'suffix_equal', label: '后缀匹配' },
  { value: 'str_contains', label: '包含' },
  { value: 'not_str_contains', label: '不包含' },
  { value: 're_equal', label: '正则匹配' }
]
const QUERY_OPERATORS = [
  { value: 'str_equal', label: '等于' },
  { value: 'str_contains', label: '包含' },
  { value: 'not_str_contains', label: '不包含' },
  { value: 're_equal', label: '正则匹配' }
]

const FILE_OPERATORS = [
  { value: 'str_equal', label: '等于' },
  { value: 'str_contains', label: '包含' },
  { value: 'not_str_contains', label: '不包含' },
  { value: 're_equal', label: '正则匹配' }
]

const FILE_NAMES_OPERATORS = [
  { value: 'str_equal', label: '等于' },
  { value: 'str_contains', label: '包含' },
  { value: 'not_str_contains', label: '不包含' },
  { value: 're_equal', label: '正则匹配' }
]

const POST_BODY_OPERATORS = [
  { value: 'str_contains', label: '包含' },
  { value: 'not_str_contains', label: '不包含' },
  { value: 're_equal', label: '正则匹配' }
]

const REQ_HEADER_OPERATORS = [
  { value: 'str_equal', label: '等于' },
  { value: 'str_contains', label: '包含' },
  { value: 're_equal', label: '正则匹配' },
  { value: 'not_exist', label: '不存在' },
  { value: 'num_gt', label: '数值大于' },
  { value: 'num_le', label: '数值小于' },
  { value: 'num_equal', label: '数值等于' }
]

const CC_ACTIONS = [
  { label: '验证码', value: 'captcha' },
  { label: '返回400', value: 'http400' }
]
const CC_ACTION_TEXT = {
  captcha: '验证码',
  http400: '400响应'
}
const ACTION_TEXT = {
  1: '允许',
  2: '阻断',
  4: '日志'
}
const ARGS_OPTIONS = [
  { name: 'query', value: '请求路径参数' },
  { name: 'post', value: 'post参数' },
  { name: 'cookie', value: 'cookie参数' }
]
const ARGS_TEXT = {
  query: '请求行参数',
  post: '请求体参数',
  cookie: 'Cookie参数'
}
export {
  ACTION_TYPES,
  REQ_HEADER_OPERATORS,
  IP_OPERATORS,
  METHOD_OPERATORS,
  URI_OPERATORS,
  QUERY_OPERATORS,
  POST_BODY_OPERATORS,
  FILE_OPERATORS,
  FILE_NAMES_OPERATORS,
  OPERATORS_TEXT,
  VARIABLES_TEXT,
  CC_ACTIONS,
  CC_ACTION_TEXT,
  ACTION_TEXT,
  ARGS_OPTIONS,
  ARGS_TEXT
}

