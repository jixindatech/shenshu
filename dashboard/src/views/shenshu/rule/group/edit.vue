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
      <el-form-item label="规则类型：" prop="type">
        <el-select v-model="formData.type" placeholder="请选择规则类型">
          <el-option v-for="(item,index) in RULE_TYPES " :key="index" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item label="规则权重：" prop="priority">
        <el-input v-model.number="formData.priority" maxlength="30" type="number" />
      </el-form-item>
      <el-form-item label="模式：" prop="status">
        <el-radio-group v-model="formData.status" size="mini">
          <el-radio :label="1" border>阻断</el-radio>
          <el-radio :label="2" border>关闭</el-radio>
          <el-radio :label="4" border>日志</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="解码配置：" prop="decoder">
        <el-checkbox-group v-model="formData.decoder" @change="handleCheckedDecoder">
          <el-checkbox v-for="(item, index) in decoderOptions" :key="index" size="mini" :label="item.name">{{ item.value }}</el-checkbox>
        </el-checkbox-group>
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
import { add, update } from '@/api/rulegroup'
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
