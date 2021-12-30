<template>
  <div
    v-permission="['GET:/nginx/site', 'GET:/shenshu/site/:id/cc', 'GET:/shenshu/site/cc/:id']"
    class="app-container"
  >
    <el-form :inline="true" :model="query" size="mini">
      <el-form-item label="域名:">
        <el-select v-model="siteId" placeholder="请选择域名" @change="selectChanged">
          <el-option v-for="(item,index) in sites" :key="index" :label="item.name" :value="item.id" />
        </el-select>
      </el-form-item>
      <el-form-item label="规则名称:">
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
          v-permission="['POST:/shenshu/site/:id/cc']"
          icon="el-icon-circle-plus-outline"
          type="primary"
          @click="openAdd"
        >新增</el-button>
      </el-form-item>
    </el-form>

    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      :header-cell-style="{'text-align':'center'}"
      :cell-style="{'text-align':'center'}"
      border
      fit
      highlight-current-row
      row-key="id"
    >
      <el-table-column prop="id" label="ID" />
      <el-table-column prop="name" label="规则名" />
      <el-table-column align="center" prop="mode" label="限定方式">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.mode === 'ip'" type="success">IP</el-tag>
          <el-tag v-if="scope.row.mode === ''" type="success">Session</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="method" label="请求方法" />
      <el-table-column align="center" prop="uri" label="URI" />
      <el-table-column align="center" prop="threshold" label="阈值" />
      <el-table-column align="center" prop="action" label="执行动作">
        <template slot-scope="scope">
          {{ CC_ACTION_TEXT[scope.row.action] }}
        </template>
      </el-table-column>
      <el-table-column prop="createdAt" label="创建时间" width="220">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.createdAt }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="updateAt" label="更新时间" width="220">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.updateAt }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作" width="250">
        <template slot-scope="scope">
          <el-button
            v-permission="['PUT:/shenshu/site/cc/:id']"
            type="success"
            size="mini"
            @click="handleEdit(scope.row.id)"
          >编辑</el-button>
          <el-button
            v-permission="['DELETE:/shenshu/site/cc/:id']"
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

    <edit
      :title="edit.title"
      :site="siteId"
      :form-data="edit.formData"
      :visible="edit.visible"
      :remote-close="remoteClose"
    />

  </div>
</template>

<script>
import * as site from '@/api/site'
import { getList, deleteById, getById } from '@/api/cc'
import Edit from './edit'
import { OPERATORS_TEXT, CC_ACTION_TEXT } from '@/utils/rule'
export default {
  name: 'CC',
  components: { Edit },
  data() {
    return {
      OPERATORS_TEXT,
      CC_ACTION_TEXT,
      siteId: 1,
      sites: [],
      query: {},
      edit: {
        title: '',
        visible: false,
        formData: {}
      },
      page: {
        current: 1,
        size: 10,
        total: 0
      },
      list: [],
      listLoading: true,
      checkedUserList: []
    }
  },
  watch: {
    '$route.path': {
      immediate: true,
      handler() {
        const id = this.$route.params.site
        if (id === undefined) {
          this.siteId = 1
        } else {
          this.siteId = id
        }
      }
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      site.getList({}, 0).then((response) => {
        this.sites = response.data.list
      })

      this.listLoading = true
      getList(
        this.siteId,
        this.query,
        this.page.current,
        this.page.size
      ).then(response => {
        const { data } = response
        this.list = data.list
        this.page.total = data.total
        this.listLoading = false
      })
    },
    selectChanged(id) {
      this.siteId = id
    },
    queryData() {
      this.page.current = 1
      this.fetchData()
    },
    reload() {
      this.query = {}
      this.fetchData()
    },
    openAdd() {
      this.edit.title = '新增'
      this.edit.visible = true
    },
    remoteClose() {
      this.edit.formData = {}
      this.edit.visible = false
      this.fetchData()
    },
    handleSizeChange(val) {
      this.page.size = val
      this.fetchData()
    },
    handleCurrentChange(val) {
      this.page.current = val
      this.fetchData()
    },
    async handleEdit(id) {
      getById(id).then((response) => {
        const { data } = response
        if (data.item.Site !== undefined) {
          delete data.item.Site
        }
        this.edit.formData = data.item
        this.edit.title = '编辑'
        this.edit.visible = true
      })
    },
    handleDelete(id) {
      this.$confirm('确认删除这条记录吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          deleteById(id).then((response) => {
            this.$message({
              type: response.code === 0 ? 'success' : 'error',
              message: '删除成功!'
            })
            this.fetchData()
          })
        })
        .catch(() => {
        })
    }
  }
}
</script>
