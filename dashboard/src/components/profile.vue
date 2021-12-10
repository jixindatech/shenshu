<template>
  <el-dialog
    :title="title"
    :visible.sync="visible"
    :modal-append-to-body="false"
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
      <el-form-item label="昵称：" prop="displayName">
        <el-input v-model="formData.displayName" maxlength="50" />
      </el-form-item>
      <el-form-item label="旧密码：" prop="oldpassword">
        <el-input v-model="formData.oldpassword" maxlength="50" />
      </el-form-item>
      <el-form-item label="新密码：" prop="newpassword">
        <el-input v-model="formData.newpassword" maxlength="50" />
      </el-form-item>
      <el-form-item label="手机号：" prop="phone">
        <el-input v-model="formData.phone" maxlength="11" />
      </el-form-item>
      <el-form-item label="邮箱：" prop="email">
        <el-input v-model="formData.email" maxlength="30" />
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
import { updateInfo } from '@/api/user'
export default {
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    remoteClose: {
      type: Function,
      default: function() {}
    }
  },

  data() {
    return {
      title: '个人设置',
      formData: {},
      rules: {
        displayName: [{ required: true, message: '请输入昵称', trigger: 'blur' }],
        phone: [{ required: true, message: '请输入手机号', trigger: 'blur' }],
        email: [{ required: true, message: '请输入邮箱', trigger: 'blur' }]
      }
    }
  },
  watch: {
    visible(newVal, oldVal) {
      if (newVal === true) {
        this.formData = this.$store.getters.info
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
      response = await updateInfo(this.formData)
      console.log(response)
      if ((response.code === 0)) {
        this.$message({ message: '保存成功', type: 'success' })
        this.handleClose()
      } else {
        this.$message({ message: '保存失败', type: 'error' })
      }
    },

    handleClose() {
      this.$refs['formData'].resetFields()
      this.remoteClose()
    }
  }
}
</script>
