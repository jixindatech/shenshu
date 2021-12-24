const ACTION_TYPES = [
  { label: '允许', value: 1 },
  { label: '阻断', value: 2 },
  { label: '日志', value: 4 }
]
const VARIABLES = [
  'IP',
  'HTTP_VERSION',
  'METHOD',
  'URI',
  'URI_ARGS',
  'QUERY_STRING',
  'REQUEST_HEADERS',
  'BODY_ARGS',
  'REQUEST_BODY',
  'FILES',
  'FILES_NAMES',
  'FILES_CONTENT',
  'FILES_SIZE',
  'FILES_SIZES'
]

const VARIABLES_TEXT = {
  IP: '请求IP',
  HTTP_VERSION: 'Http版本',
  METHOD: '请求方法',
  URI: 'URI',
  URI_ARGS: '请求参数值',
  QUERY_STRING: '请求参数行',
  REQUEST_HEADERS: '请求头',
  BODY_ARGS: '请求体参数',
  REQUEST_BODY: '请求体内容',
  FILES: '上传文件的名称',
  FILES_NAMES: '上传文件名称',
  FILES_CONTENT: '上传文件内容',
  FILES_SIZE: '上传文件体积',
  FILES_SIZES: '上传文件总体积'
}

const OPERATORS_TEXT = {
  EXISTS: '存在',
  CONTAINS: '属于',
  EQUALS: '等于',
  STR_CONTAINS: '包含',
  GREATER: '大于',
  LESS: '小于',
  REGEX: '正则匹配'
}

const EQUAL_OPERATORS = [
  { value: 'EQUALS', label: '等于' }
]

const STR_OPERATORS = [
  { value: 'EQUALS', label: '等于' },
  { value: 'STR_CONTAINS', label: '包含' },
  { value: 'REGEX', label: '正则匹配' }
]

const NUM_OPERATORS = [
  { value: 'GREATER', label: '数值大于' },
  { value: 'LESS', label: '数值小于' }
]

const REQUEST_HEADER_OPERATORS = [
  { value: 'EQUALS', label: '等于' },
  { value: 'STR_CONTAINS', label: '包含' },
  { value: 'REGEX', label: '正则匹配' },
  { value: 'GREATER', label: '数值大于' },
  { value: 'LESS', label: '数值小于' }
]

const CC_ACTIONS = [
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
const RULE_TYPES = [
  { value: 1, label: '批规则' },
  { value: 2, label: '复杂规则' }
]
const RULE_TYPES_TEXT = {
  1: '批规则',
  2: '复杂规则'
}

export {
  ACTION_TYPES,
  VARIABLES,
  REQUEST_HEADER_OPERATORS,
  NUM_OPERATORS,
  STR_OPERATORS,
  EQUAL_OPERATORS,
  OPERATORS_TEXT,
  VARIABLES_TEXT,
  CC_ACTIONS,
  CC_ACTION_TEXT,
  ACTION_TEXT,
  RULE_TYPES,
  RULE_TYPES_TEXT
}

