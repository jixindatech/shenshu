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
      <el-form-item label="规则名称：" prop="name">
        <el-input v-model="formData.name" />
      </el-form-item>
      <el-form-item label="限定方式：" prop="mode">
        <el-radio-group v-model="formData.mode">
          <el-radio :label="'ip'" border>IP</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="请求方法：" prop="method">
        <el-select v-model="formData.method" placeholder="请选择匹配方式">
          <el-option v-for="(item,index) in methodList" :key="index" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item prop="uri">
        <span slot="label">URI
          <el-tooltip placement="top" effect="light">
            <div slot="content">
              web服务的访问路径，支持通配符
            </div>
            <i class="el-icon-question" />
          </el-tooltip>
        </span>
        <el-input v-model="formData.uri" />
      </el-form-item>
      <el-row>
        <el-col :span="9">
          <el-form-item label="阈值：" prop="threshold">
            <el-input v-model.number="formData.threshold" oninput="value=value.replace(/[^\d]/g,'')" placeholder="阈值" style="width: 150px">
              <template slot="append">次</template>
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="9">
          <el-form-item prop="duration">
            <el-input v-model.number="formData.duration" oninput="value=value.replace(/[^\d]/g,'')" placeholder="单位时间" style="width: 150px">
              <template slot="append">秒</template>
            </el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item label="执行动作：" prop="action">
        <el-select v-model="formData.action" placeholder="请选择执行动作">
          <el-option v-for="(item,index) in CC_ACTIONS" :key="index" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item label="备注：" prop="remark">
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
import { add, update } from '@/api/cc'
import { CC_ACTIONS } from '@/utils/rule'

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
    site: {
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
      methodList: [
        { label: 'GET', value: 'GET' },
        { label: 'POST', value: 'POST' },
        { label: 'PUT', value: 'PUT' },
        { label: 'DELETE', value: 'DELETE' }
      ],
      URI_OPERATORS: [
        { value: 'str_equal', label: '等于' }
      ],
      CC_ACTIONS,
      rules: {
        name: [
          { required: true, message: '请输入规则名称', trigger: 'blur' }
        ],
        method: [
          { required: true, message: '请选择请求方法', trigger: 'change' }
        ],
        mode: [
          { required: true, message: '请选择限定方式', trigger: 'change' }
        ],
        match: [
          { required: true, message: '请选择匹配方式', trigger: 'change' }
        ],
        uri: [
          { required: true, message: '请输入uri', trigger: 'blur' }
        ],
        duration: [
          { required: true, message: '请输入时间间隔', trigger: 'blur' }
        ],
        threshold: [
          { required: true, message: '请输入阈值', trigger: 'blur' }
        ],
        action: [
          { required: true, message: '请选择匹配动作', trigger: 'change' }
        ]
      }
    }
  },
  watch: {
    visible(newVal, oldVal) {
      if (newVal) {
        console.log(newVal)
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
        response = await update(this.formData.id, this.formData)
      } else {
        response = await add(this.site, this.formData)
      }

      if (response.code === 0) {
        this.$message({ message: '保存成功', type: 'success' })
        this.handleClose()
      }
    },
    handleClose() {
      this.sites = []
      this.$refs['formData'].resetFields()
      this.remoteClose()
    }
  }
}
</script>
