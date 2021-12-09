<template>
  <div class="app-container">
    <el-form
      ref="formData"
      :rules="rules"
      :model="formData"
      label-position="left"
      label-width="120px"
      status-icon
    >
      <el-form-item label="邮箱域名：" prop="host" style="width: 450px">
        <el-input v-model="formData.host" placeholder="请输入邮箱域名" />
      </el-form-item>
      <el-form-item label="端口：" prop="port" style="width: 450px">
        <el-input v-model.number="formData.port" placeholder="请输入端口" />
      </el-form-item>
      <el-form-item label="邮箱用户：" prop="sender" style="width: 450px">
        <el-input v-model="formData.sender" placeholder="请输入邮箱用户" />
      </el-form-item>
      <el-form-item label="邮箱密码：" prop="password" style="width: 450px">
        <el-input v-model="formData.password" show-password placeholder="请输入邮箱密码" />
      </el-form-item>
      <el-form-item v-permission="['POST:/system/mail', 'PUT:/system/mail/:id']">
        <el-button type="primary" @click="submit">保存</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { get, add, update } from '@/api/email'
export default {
  name: 'Email',
  components: { },
  props: {
    params: {
      type: String,
      default: '0'
    },
    name: {
      type: String,
      default: '0'
    }
  },
  data() {
    return {
      formData: {},
      rules: {
        host: [
          { required: true, message: '请输入域名', trigger: 'change' }
        ],
        port: [
          { required: true, message: '请输入端口', trigger: 'change' }
        ],
        sender: [
          { required: true, message: '请输入邮箱用户', trigger: 'change' }
        ],
        password: [
          { required: true, message: '请输入邮箱密码', trigger: 'change' }
        ]
      }
    }
  },
  watch: {
    params(newVal, oldVal) {
      if (newVal === this.name) {
        this.fetchData()
      }
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      get().then(response => {
        const { data } = response
        if (data.item && data.item.id !== 0) {
          this.formData = data.item
        }
      })
    },
    submit() {
      if (this.formData.id > 0) {
        update(this.formData.id, this.formData).then(response => {
          this.$message({ message: '保存成功', type: 'success' })
        })
      } else {
        add(this.formData).then(response => {
          this.$message({ message: '保存成功', type: 'success' })
        })
      }
    }
  }
}
</script>
