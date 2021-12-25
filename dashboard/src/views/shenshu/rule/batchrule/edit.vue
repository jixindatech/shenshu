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
        <el-input v-model="formData.name" placeholder="请输入规则名称" />
      </el-form-item>
      <el-form-item label="匹配内容：" prop="pattern" style="width: 300px">
        <el-input v-model="formData.pattern" placeholder="请输入匹配内容" />
      </el-form-item>
      <el-form-item label="匹配动作：" prop="action" style="width: 300px">
        <el-select v-model="formData.action" placeholder="请选择匹配动作">
          <el-option v-for="(item,index) in ACTION_TYPES" :key="index" :label="item.label" :value="item.value" />
        </el-select>
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
import * as api from '@/api/rulebatch'
// import { ACTION_TYPES } from '@/utils/rule'

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
    groupId: {
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
      ACTION_TYPES: [{ label: '阻断', value: 2 }],
      rules: {
        name: [
          { required: true, message: '请输入名称', trigger: 'change' }
        ],
        pattern: [
          { required: true, message: '请输入匹配内容', trigger: 'change' }
        ],
        action: [
          { required: true, message: '请选择规则动作', trigger: 'change' }
        ],
        status: [
          { required: true, message: '请选择状态', trigger: 'blur' }
        ]
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
        response = await api.add(this.groupId, this.formData)
      }

      if (response.code === 0) {
        this.$message({ message: '保存成功', type: 'success' })
        this.handleClose()
      }
    },
    handleClose() {
      this.$refs['formData'].resetFields()
      this.remoteClose()
    }
  }
}
</script>
