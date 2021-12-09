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
      <el-form-item label="SMS域名：" prop="host" style="width: 500px">
        <el-input v-model="formData.host" placeholder="请输入SMS域名" />
      </el-form-item>
      <el-form-item label="secretId：" prop="secretId" style="width: 500px">
        <el-input v-model="formData.secretId" placeholder="请输入secretId" />
      </el-form-item>
      <el-form-item label="secretKey：" prop="secretKey" style="width: 500px">
        <el-input v-model="formData.secretKey" show-password placeholder="请输入secretKey" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submit">保存</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { get, add, update } from '@/api/txsms'
export default {
  name: 'Txsms',
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
          { required: true, message: '请输入SMS域名', trigger: 'change' }
        ],
        secretId: [
          { required: true, message: '请输入secretId', trigger: 'change' }
        ],
        secretKey: [
          { required: true, message: '请输入secretKey', trigger: 'change' }
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
