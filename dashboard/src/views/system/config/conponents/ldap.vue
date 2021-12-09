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
      <el-form-item label="Ldap类型：" prop="type" style="width: 450px">
        <el-select v-model="formData.type" placeholder="请选择">
          <el-option
            v-for="item in LDAP_TYPE_OPTIONS"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="Ldap服务器：" prop="host" style="width: 450px">
        <el-input v-model="formData.host" placeholder="请输入服务器" />
      </el-form-item>
      <el-form-item label="端口：" prop="port" style="width: 450px">
        <el-input v-model.number="formData.port" placeholder="请输入端口" />
      </el-form-item>
      <el-form-item label="DN用户：" prop="dn" style="width: 450px">
        <el-input v-model="formData.dn" placeholder="请输入DN用户" />
      </el-form-item>
      <el-form-item label="密码：" prop="password" style="width: 450px">
        <el-input v-model="formData.password" show-password placeholder="请输入密码" />
      </el-form-item>
      <el-form-item label="BaseDn：" prop="basedn" style="width: 450px">
        <el-input v-model="formData.basedn" placeholder="请输入BaseDN" />
      </el-form-item>
      <el-form-item v-permission="['POST:/system/ldap', 'PUT:/system/ldap/:id']">
        <el-button type="primary" @click="submit">保存</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { LDAP_TYPE_OPTIONS } from '@/utils/const'
import { get, add, update } from '@/api/ldap'
export default {
  name: 'Ldap',
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
      LDAP_TYPE_OPTIONS,
      formData: {},
      rules: {
        type: [
          { required: true, message: '请选择类型', trigger: 'change' }
        ],
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
