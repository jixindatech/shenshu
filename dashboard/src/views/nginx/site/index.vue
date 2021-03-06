<template>
  <div
    v-permission="['GET:/nginx/site', 'GET:/nginx/site/:id']"
    class="app-container"
  >
    <el-form :inline="true" :model="query" size="mini">
      <el-form-item
        label="Site名称:"
      >
        <el-input v-model.trim="query.name" />
      </el-form-item>
      <el-form-item>
        <el-button
          icon="el-icon-search"
          type="primary"
          @click="queryData"
        >查询</el-button>
        <el-button
          icon="el-icon-refresh"
          @click="reload"
        >重置</el-button>
        <el-button
          v-permission="['POST:/nginx/site']"
          icon="el-icon-circle-plus-outline"
          type="primary"
          @click="openAdd"
        >新增</el-button>
        <el-button
          v-permission="['POST:/shenshu/site/:id/enable']"
          type="success"
          size="mini"
          @click="enableConfig()"
        >配置下发</el-button>
      </el-form-item>
    </el-form>
    <el-table
      ref="dataTable"
      :data="list"
      stripe
      border
      style="width: 100%"
      row-key="id"
    >
      <el-table-column align="center" prop="name" label="名称" width="150px" />
      <el-table-column align="center" prop="host" label="域名" width="200px" />
      <el-table-column align="center" prop="path" label="路径" width="150px" />
      <el-table-column align="center" prop="upstreamRef" label="Upstream">
        <template v-if="scope.row.upstreamRef.length === 1" slot-scope="scope">
          {{ scope.row.upstreamRef[0].name }}
        </template>
      </el-table-column>
      <el-table-column align="center" prop="status" label="状态">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === 1" type="success">启用</el-tag>
          <el-tag v-if="scope.row.status === 2" type="danger">停用</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="remark" label="备注" />
      <el-table-column align="center" label="操作" width="550px">
        <template slot-scope="scope">
          <el-button
            v-permission="['PUT:/nginx/site/:id']"
            type="success"
            size="mini"
            @click="handleEdit(scope.row.id)"
          >编辑</el-button>
          <el-button
            v-permission="['GET:/shenshu/site/:id/ip', 'GET:/shenshu/site/ip/:id']"
            type="primary"
            size="mini"
            @click="ipConfig(scope.row.id)"
          >IP管理</el-button>
          <el-button
            v-permission="['GET:/shenshu/site/:id/cc', 'GET:/shenshu/site/cc/:id']"
            type="primary"
            size="mini"
            @click="ccConfig(scope.row.id)"
          >CC配置</el-button>
          <el-button
            v-permission="['GET:/shenshu/batchgroup']"
            type="primary"
            size="mini"
            @click="rulegroupConfig(scope.row.id, TYPE_RULE_GROUP.TYPE_BATCH_GROUP)"
          >Batch规则</el-button>
          <el-button
            v-permission="['GET:/shenshu/specificgroup']"
            type="primary"
            size="mini"
            @click="rulegroupConfig(scope.row.id, TYPE_RULE_GROUP.TYPE_SPECIFIC_GROUP)"
          >Sepcififc规则</el-button>
          <el-button
            v-permission="['DELETE:/nginx/site/:id']"
            type="danger"
            size="mini"
            @click="handleDelete(scope.row.id)"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page.current"
      :page-sizes="[10, 20, 50]"
      :page-size="page.size"
      layout="total, sizes, prev, pager, next, jumper"
      :total="page.total"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    />
    <Edit
      :title="edit.title"
      :form-data="edit.formData"
      :visible="edit.visible"
      :remote-close="remoteClose"
    />

    <el-dialog title="设置规则" :visible.sync="rulegroup.batchvisible" width="65%">
      <BatchRuleGroup
        :ids="rulegroup.ids"
        :site="rulegroup.site"
        @getRuleGroup="getRuleGroup"
      />
    </el-dialog>

    <el-dialog title="设置规则" :visible.sync="rulegroup.speicificvisible" width="65%">
      <SpecificGroup
        :ids="rulegroup.ids"
        :site="rulegroup.site"
        @getRuleGroup="getRuleGroup"
      />
    </el-dialog>

  </div>
</template>

<script>
import * as api from '@/api/site'
import Edit from './edit'
import BatchRuleGroup from '@/views/shenshu/rule/batchgroup'
import SpecificGroup from '@/views/shenshu/rule/specificgroup'
import { TYPE_RULE_GROUP } from '@/utils/const'

export default {
  name: 'Site',
  components: { Edit, BatchRuleGroup, SpecificGroup },
  data() {
    return {
      TYPE_RULE_GROUP,
      list: [],
      page: {
        current: 1,
        size: 20,
        total: 0
      },
      query: {},
      edit: {
        title: '',
        visible: false,
        formData: {}
      },
      checkedSitesList: [],
      rulegroup: {
        type: 0,
        batchvisible: false,
        speicificvisible: false,
        ids: [],
        site: 0
      }
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    async fetchData() {
      const { data } = await api.getList(
        this.query,
        this.page.current,
        this.page.size
      )

      this.list = data.list
      this.page.total = data.total
    },

    handleSizeChange(val) {
      this.page.size = val
      this.fetchData()
    },

    handleCurrentChange(val) {
      this.page.current = val
      this.fetchData()
    },

    queryData() {
      this.page.current = 1
      this.fetchData()
    },

    reload() {
      this.query = {}
      this.fetchData()
    },

    handleDelete(id) {
      this.$confirm('确认删除这条记录吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          api.deleteById(id).then((response) => {
            this.$message({
              type: 'success',
              message: '删除成功!'
            })
            this.fetchData()
          })
        })
        .catch(() => {
        })
    },

    openAdd() {
      this.edit.title = '新增'
      this.edit.visible = true
    },

    handleEdit(id) {
      api.get(id).then((response) => {
        this.edit.formData = response.data.item
        this.edit.title = '编辑'
        this.edit.visible = true
      })
    },
    remoteClose() {
      this.edit.formData = {}
      this.edit.visible = false
      this.fetchData()
    },
    ipConfig(id) {
      this.$router.push({ name: 'IP', params: { site: id }})
    },
    ccConfig(id) {
      this.$router.push({ name: 'CC', params: { site: id }})
    },
    async rulegroupConfig(id, type) {
      const { data } = await api.GetSiteRuleGroup(id, { type: type })
      if (type === TYPE_RULE_GROUP.TYPE_BATCH_GROUP) {
        this.rulegroup.batchvisible = true
      } else {
        this.rulegroup.speicificvisible = true
      }
      this.rulegroup.type = type
      this.rulegroup.site = id
      this.rulegroup.ids = data.ids
    },
    enableConfig() {
      const id = 1
      api.enable(id).then((response) => {
        this.$message({
          type: 'success',
          message: '下发成功!'
        })
      })
    },
    getRuleGroup(ids) {
      const data = { type: this.rulegroup.type, ids: ids }
      api.UpdateSiteRuleGroup(this.rulegroup.site, data).then((response) => {
        this.$message({
          type: 'success',
          message: '更新成功!'
        })
      })
      if (this.rulegroup.type === TYPE_RULE_GROUP.TYPE_BATCH_GROUP) {
        this.rulegroup.batchvisible = false
      } else {
        this.rulegroup.speicificvisible = false
      }
      this.rulegroup.type = 0
      this.rulegroup.site = 0
      this.rulegroup.ids = []
    }
  }
}
</script>
