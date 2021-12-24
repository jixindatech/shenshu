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
      <el-form-item label="IP名称：" prop="name">
        <el-input v-model="formData.name" maxlength="30" />
      </el-form-item>
      <el-form-item>
        <span slot="label">IP
          <el-tooltip placement="right">
            <div slot="content">
              支持CIDR方式
            </div>
            <i class="el-icon-question" />
          </el-tooltip>
        </span>

        <el-table
          :row-style="{height:'10px'}"
          :cell-style="{padding:'1px'}"
          style="font-size: 8px; margin-top: 0px;"
          size="mini"
          :show-header="false"
          :data="ips"
        >
          <el-table-column align="center" label="IP" width="270px">
            <template slot-scope="scope">
              <el-form-item :prop="'ips.' + scope.$index + '.ip'" :rules="rules.ips">
                <el-input v-model="scope.row.ip" size="mini" placeholder="请输入IP" />
              </el-form-item>
            </template>
          </el-table-column>
          <el-table-column align="center" width="30px">
            <template slot-scope="scope">
              <el-button type="text" icon="el-icon-delete" size="medium" @click="deleteIP(scope.row, scope.$index)" />
            </template>
          </el-table-column>
        </el-table>
        <el-button type="text" icon="el-icon-plus" size="mini" style="margin-bottom: 20px;" @click="addIP()">新增IP</el-button>
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
import * as globalIp from '@/api/globalip'
import { add, update } from '@/api/ip'

// import { validateIP } from '@/utils/validate'

export default {
  props: {
    title: {
      type: String,
      default: ''
    },
    type: {
      type: String,
      default: ''
    },
    site: {
      type: Number,
      default: 0
    },
    visible: {
      type: Boolean,
      default: false
    },
    formData: {
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
      ips: [],
      rules: {
        name: [{ required: true, message: '请输入名称', trigger: 'blur' }]
        /*
        ips: [
          { required: true, message: '请输入IP' },
          { validator: validateIP, tirgger: 'change' }
          ]*/
      }
    }
  },

  watch: {
    visible(newVal, oldVal) {
      if (newVal) {
        if (this.formData.ip === undefined) {
          this.formData.ip = []
          const item = { ip: '' }
          this.ips.push(item)
        } else {
          this.formData.ip.forEach(element => {
            const item = { ip: element }
            this.ips.push(item)
          })
        }
        this.formData.ips = this.ips
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
      this.formData.ip = []
      this.ips.forEach(element => {
        this.formData.ip.push(element.ip)
      })

      this.formData.type = Number(this.type)
      let response = null
      if (this.site === 0) {
        if (this.formData.id) {
          response = await globalIp.update(this.formData.id, this.formData)
        } else {
          response = await globalIp.add(this.formData)
        }
      } else {
        if (this.formData.id) {
          response = await update(this.formData.id, this.formData)
        } else {
          response = await add(this.site, this.formData)
        }
      }

      if ((response.code === 0)) {
        this.$message({ message: '保存成功', type: 'success' })
        this.handleClose()
      }
    },
    handleClose() {
      this.$refs['formData'].resetFields()
      this.ips = []
      this.remoteClose()
    },
    addIP() {
      const item = { ip: '' }
      this.ips.push(item)
    },
    deleteIP(row, index) {
      this.ips.splice(index, 1)
    }
  }
}
</script>
