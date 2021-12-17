<template>
  <el-dialog
    :title="title"
    :visible.sync="visible"
    width="690px"
    :before-close="handleClose"
  >
    <el-form
      ref="formData"
      :rules="rules"
      :model="formData"
      label-position="top"
      label-width="80px"
      status-icon
    >
      <el-form-item label="规则名称：" prop="name" style="width: 300px">
        <el-input v-model="formData.name" placeholder="请输入规则类别" />
      </el-form-item>
      <p style="line-height: 40px;color: #606266; font-weight: bold; margin-bottom: 0; margin-top: 0; display: inline-block"><span style="color: red;display: inline-block;">*</span>匹配条件</p>
      <el-tooltip content="支持CIDR类型IP,多个IP以英文逗号(,)分割" placement="right" effect="light">
        <i class="el-icon-question" />
      </el-tooltip>
      <el-table
        :header-cell-style="{background:'#646564',color:'#fff'}"
        :row-style="{height:'20px'}"
        :cell-style="{padding:'1px'}"
        style="font-size: 15px; margin-top: 0px;"
        size="mini"
        show-header
        :data="formData.rules"
      >
        <el-table-column align="center" label="匹配字段" width="165px">
          <template slot-scope="scope">
            <el-form-item :prop="'rules.' + scope.$index + '.variable'" :rules="rules.variable">
              <el-select v-model="scope.row.variable" size="mini" placeholder="请选择匹配字段" @change="variableChange(scope.row)">
                <el-option v-for="(item,index) in VARIABLES " :key="index" :label="item" :value="item" />
              </el-select>
            </el-form-item>
          </template>
        </el-table-column>
        <el-table-column align="center" label="运算符" width="150px">
          <template slot-scope="scope">
            <el-form-item v-if="scope.row.variable === ''" :prop="'rules.' + scope.$index + '.operator'" :rules="rules.operator">
              <el-select v-model="scope.row.operator" size="mini" placeholder="请选择匹配字段">
                <el-option v-for="(item,index) in OPERATORS " :key="index" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item v-if="scope.row.variable === 'REQ_HEADER'" :prop="'rules.' + scope.$index + '.operator'" :rules="rules.operator">
              <el-select v-model="scope.row.operator" size="mini" placeholder="请选择匹配字段">
                <el-option v-for="(item,index) in REQ_HEADER_OPERATORS " :key="index" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item v-if="scope.row.variable === 'IP'" :prop="'rules.' + scope.$index + '.operator'" :rules="rules.operator">
              <el-select v-model="scope.row.operator" size="mini" placeholder="请选择匹配字段">
                <el-option v-for="(item,index) in IP_OPERATORS " :key="index" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item v-if="scope.row.variable === 'METHOD'" :prop="'rules.' + scope.$index + '.operator'" :rules="rules.operator">
              <el-select v-model="scope.row.operator" size="mini" placeholder="请选择匹配字段">
                <el-option v-for="(item,index) in METHOD_OPERATORS " :key="index" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item v-if="scope.row.variable === 'URI'" :prop="'rules.' + scope.$index + '.operator'" :rules="rules.operator">
              <el-select v-model="scope.row.operator" size="mini" placeholder="请选择匹配字段">
                <el-option v-for="(item,index) in URI_OPERATORS " :key="index" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item v-if="scope.row.variable === 'QUERY'" :prop="'rules.' + scope.$index + '.operator'" :rules="rules.operator">
              <el-select v-model="scope.row.operator" size="mini" placeholder="请选择匹配字段">
                <el-option v-for="(item,index) in QUERY_OPERATORS " :key="index" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item v-if="scope.row.variable === 'POST_BODY'" :prop="'rules.' + scope.$index + '.operator'" :rules="rules.operator">
              <el-select v-model="scope.row.operator" size="mini" placeholder="请选择匹配字段">
                <el-option v-for="(item,index) in POST_BODY_OPERATORS " :key="index" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item v-if="scope.row.variable === 'FILE'" :prop="'rules.' + scope.$index + '.operator'" :rules="rules.file">
              <el-select v-model="scope.row.operator" size="mini" placeholder="请选择匹配字段">
                <el-option v-for="(item,index) in FILE_OPERATORS " :key="index" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item v-if="scope.row.variable === 'FILE_NAMES'" :prop="'rules.' + scope.$index + '.operator'" :rules="rules.fileName">
              <el-select v-model="scope.row.operator" size="mini" placeholder="请选择匹配字段">
                <el-option v-for="(item,index) in FILE_NAMES_OPERATORS " :key="index" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
          </template>
        </el-table-column>
        <el-table-column align="center" label="匹配内容" width="300px">
          <template slot-scope="scope">
            <el-form-item v-if="scope.row.variable === 'REQ_HEADER'" :prop="'rules.' + scope.$index + '.header'" :rules="rules.header">
              <el-input v-model="scope.row.header" size="mini" placeholder="请输入请求头" />
            </el-form-item>
            <el-form-item v-if="scope.row.operator !== 'not_exist'" :prop="'rules.' + scope.$index + '.pattern'" :rules="rules.pattern">
              <el-input v-model="scope.row.pattern" size="mini" placeholder="请输入匹配内容" />
            </el-form-item>
          </template>
        </el-table-column>
        <el-table-column align="center" width="30px">
          <template slot-scope="scope">
            <el-button type="text" icon="el-icon-delete" size="medium" @click="deleteRule(scope.row, scope.$index)" />
          </template>
        </el-table-column>
      </el-table>
      <el-button type="text" icon="el-icon-plus" size="mini" style="margin-bottom: 20px;" @click="addRule()">新增条件</el-button><p style="display: inline;">最多添加5条</p>
      <el-form-item label="匹配动作：" prop="action" style="width: 300px">
        <el-select v-model="formData.action" placeholder="请选择匹配动作">
          <el-option v-for="(item,index) in ACTION_TYPES" :key="index" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item label="优先级：" prop="priority" style="width: 300px">
        <el-input v-model.number="formData.priority" type="number" />
      </el-form-item>
      <el-form-item label="状态：" prop="status">
        <el-switch
          v-model="formData.status"
          active-text="开"
          :active-value="1"
          inactive-text="关"
          :inactive-value="2"
        />
      </el-form-item>
      <el-form-item label="规则备注：" prop="remark">
        <el-input v-model="formData.remark" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          size="mini"
          @click="submitForm('formData')"
        >确定</el-button>
        <el-button size="mini" @click="handleClose">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script>
import * as api from '@/api/rulespecific'
import { ACTION_TYPES, REQ_HEADER_OPERATORS, IP_OPERATORS, METHOD_OPERATORS, URI_OPERATORS, QUERY_OPERATORS, POST_BODY_OPERATORS, FILE_OPERATORS, FILE_NAMES_OPERATORS } from '@/utils/rule'
import { isInteger } from '@/utils/validate'

export default {
  props: {
    title: {
      type: String,
      default: ''
    },
    visible: {
      type: Boolean,
      default: false
    },
    id: {
      type: Number,
      default: 0
    },
    formData: {
      type: Object,
      default: function() { return {} }
    },
    remoteClose: {
      type: Function,
      default: function() { }
    }
  },

  data() {
    return {
      VARIABLES: ['IP', 'METHOD', 'URI', 'REQ_HEADER', 'QUERY', 'POST_BODY', 'FILE', 'FILE_NAMES'],
      OPERATORS: [
        { value: '', label: '' }
      ],
      ACTION_TYPES,
      IP_OPERATORS,
      METHOD_OPERATORS,
      URI_OPERATORS,
      REQ_HEADER_OPERATORS,
      QUERY_OPERATORS,
      POST_BODY_OPERATORS,
      FILE_OPERATORS,
      FILE_NAMES_OPERATORS,
      rules: {
        type: [
          { required: true, message: '请输入类型', trigger: 'change' }
        ],
        variable: [
          { required: true, message: '请选择匹配字段', trigger: 'change' }
        ],
        operator: [
          { required: true, message: '请选择匹配方式', trigger: 'change' }
        ],
        pattern: [
          { required: true, message: '请输入匹配内容', trigger: 'change' }
        ],
        header: [
          { required: true, message: '请输入请求头', trigger: 'change' }
        ],
        file: [
          { required: true, message: '请输入文件匹配内容', trigger: 'change' }
        ],
        fileName: [
          { required: true, message: '请输入文件名称', trigger: 'change' }
        ],
        action: [
          { required: true, message: '请选择规则动作', trigger: 'change' }
        ],
        priority: [
          { required: true, message: '请输入权重', trigger: 'blur', validator: isInteger }
        ]
      }
    }
  },
  watch: {
    visible(newVal, oldVal) {
      if (newVal) {
        if (this.formData.rules === undefined) {
          this.formData.rules = []
          const item = { variable: '', operator: '', pattern: '', header: '' }
          this.formData.rules.push(item)
        }
      }
    }
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.submitData()
        } else {
          // console.log('error submit!!');
          return false
        }
      })
    },

    async submitData() {
      let response = null
      if (this.formData.id) {
        response = await api.update(this.formData.id, this.formData)
      } else {
        response = await api.add(this.id, this.formData)
      }

      if (response.code === 0) {
        this.$message({ message: '保存成功', type: 'success' })
        this.handleClose()
      }
    },
    handleClose() {
      this.$refs['formData'].resetFields()
      this.remoteClose()
    },
    deleteRule(row, index) {
      if (this.formData.rules.length === 1) {
        this.$message({ message: '至少包含一个条件', type: 'error' })
        return
      }
      this.formData.rules.splice(index, 1)
    },
    addRule() {
      console.log('add rule')
      if (this.formData.rules.length === 5) {
        this.$message({ message: '至多包含五个条件', type: 'error' })
        return
      }
      const item = { variable: '', operator: '', pattern: '' }
      this.formData.rules.push(item)
    },
    variableChange(row) {
      if (row.variable === 'IP') {
        row.operator = this.IP_OPERATORS[0].value
      } else if (row.variable === 'METHOD') {
        row.operator = this.METHOD_OPERATORS[0].value
      } else if (row.variable === 'URI') {
        row.operator = this.URI_OPERATORS[0].value
      } else if (row.variable === 'QUERY') {
        row.operator = this.QUERY_OPERATORS[0].value
      } else if (row.variable === 'REQ_HEADER') {
        row.operator = this.REQ_HEADER_OPERATORS[0].value
      } else if (row.variable === 'POST_BODY') {
        row.operator = this.POST_BODY_OPERATORS[0].value
      } else if (row.variable === 'FILE') {
        row.operator = this.FILE_OPERATORS[0].value
      } else if (row.variable === 'FILE_NAMES') {
        row.operator = this.FILE_NAMES_OPERATORS[0].value
      }
    }
  }
}
</script>

<style scoped>
::v-deep .el-dialog__body{padding: 0 20px;}
::v-deep .el-table th, .el-table tr .el-form-item{margin-bottom: 0}
::v-deep .el-input--mini .el-input__inner{ border-radius: 0;}
::v-deep .cell .el-form-item__content .el-form-item__error{left: 10px; top: 35%}
</style>
