<template>
  <el-dialog
    :title="title"
    :visible.sync="visible"
    width="500px"
    :before-close="handleClose"
  >
    <el-form
      ref="formData"
      :rules="rules"
      :model="formData"
      label-width="100px"
      label-position="right"
      style="width: 400px"
      status-icon
    >
      <el-form-item label="组名称：" prop="name">
        <el-input v-model="formData.name" maxlength="30" />
      </el-form-item>
      <el-form-item prop="action">
        <span slot="label">动作
          <el-tooltip placement="top" effect="light">
            <div slot="content">
              短路或全量标识规则短路判断或者规则全部判断，日志分别代表其短路/全量记录日志,规则组合并时，优先级：短路日志>全量日志>短路>全量
            </div>
            <i class="el-icon-question" />
          </el-tooltip>
        </span>
        <el-radio-group v-model="formData.action" size="mini">
          <el-radio-button :label="1" border>日志短路</el-radio-button>
          <el-radio-button :label="2" border>日志全量</el-radio-button>
          <el-radio-button :label="3" border>拦截短路</el-radio-button>
          <el-radio-button :label="4" border>拦截全量</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="解码配置：" prop="decoder">
        <el-checkbox-group v-model="formData.decoder" @change="handleCheckedDecoder">
          <el-checkbox v-for="(item, index) in decoderOptions" :key="index" size="mini" :label="item.name">{{ item.value }}</el-checkbox>
        </el-checkbox-group>
      </el-form-item>
      <el-form-item label="规则权重：" prop="priority">
        <el-input v-model.number="formData.priority" maxlength="30" type="number" />
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
      <el-form-item label="备注：" prop="remark">
        <el-input v-model="formData.remark" type="textarea" />
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button
        type="primary"
        size=""
        @click="submitForm('formData')"
      >确定</el-button>
      <el-button size="" @click="handleClose">取消</el-button>
    </div>
  </el-dialog>
</template>

<script>
import { add, update } from '@/api/specificgroup'
import { RULE_TYPES } from '@/utils/rule'
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
    data: {
      type: Object,
      default: function() { return {} }
    },
    remoteClose: {
      type: Function,
      default: function() {}
    }
  },

  data() {
    return {
      RULE_TYPES,
      formData: {
        decoder: []
      },
      decoderOptions: [
        { name: 'multipart', value: 'Multipart解析' },
        { name: 'json', value: 'JSON解析' },
        { name: 'form', value: 'Form解析' }],
      rules: {
        name: [{ required: true, message: '请输入名称', trigger: 'blur' }],
        type: [{ required: true, message: '请选择类型', trigger: 'blur' }],
        priority: [
          { required: true, message: '请选择类型', trigger: 'blur' },
          { validator: isInteger }],
        status: [
          { required: true, message: '请选择模式', trigger: 'change' }],
        decoder: [
          { type: 'array', required: true, message: '请选择解码方式', trigger: 'change' }
        ]
      }
    }
  },
  watch: {
    visible(newVal, oldVal) {
      if (newVal === true) {
        if (this.data.id !== undefined) {
          this.formData = JSON.parse(JSON.stringify(this.data))
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
          return false
        }
      })
    },

    async submitData() {
      let response = null
      if (this.formData.id) {
        response = await update(this.formData.id, this.formData)
      } else {
        response = await add(this.formData)
      }

      if ((response.code === 0)) {
        this.$message({ message: '保存成功', type: 'success' })
        this.handleClose()
      }
    },
    handleClose() {
      this.$refs['formData'].resetFields()
      this.remoteClose()
    },
    handleCheckedDecoder(value) {
      console.log(value)
    }
  }
}
</script>

<style scoped>
::v-deep .el-radio-button--mini .el-radio-button__inner  {padding: 5 7px;}
</style>
