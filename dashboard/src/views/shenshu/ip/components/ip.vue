<template>
  <div
    class="app-container"
  >
    <el-form :inline="true" :model="query" size="mini">
      <el-form-item label="IP名称:">
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
          v-permission="['POST:/shenshu/site/:id/ip']"
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
      <el-table-column prop="name" label="IP名称" />
      <el-table-column align="center" prop="ip" label="IP">
        <template slot-scope="scope">
          <div v-for="(item, index) in scope.row.ip" :key="index">
            <el-input :value="item" size="mini" />
          </div>
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
            v-permission="['PUT:/shenshu/site/ip/:id']"
            type="primary"
            size="mini"
            @click="handleEdit(scope.row.id)"
          >编辑</el-button>
          <el-button
            v-permission="['DELETE:/shenshu/site/ip/:id']"
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
      :type="type"
      :visible="edit.visible"
      :remote-close="remoteClose"
    />

  </div>
</template>

<script>
import * as globalIp from '@/api/globalip'
import { getList, deleteById, getById } from '@/api/ip'
import Edit from './edit'

export default {
  components: { Edit },
  props: {
    params: {
      type: String,
      default: '1'
    },
    siteId: {
      type: Number,
      default: 0
    },
    type: {
      type: String,
      default: '1'
    }
  },
  data() {
    return {
      query: {},
      edit: {
        title: '',
        type: '',
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
      checkedUserList: [],
      site: {
        visible: false,
        ids: [],
        ip: null
      }
    }
  },
  watch: {
    params(newVal, oldVal) {
      if (newVal === this.type) {
        this.fetchData()
      }
    },
    siteId(newVal, oldVal) {
      this.fetchData()
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.query.type = Number(this.type)
      this.listLoading = true
      if (this.siteId === 0) {
        globalIp.getList(
          this.query,
          this.page.current,
          this.page.size
        ).then(response => {
          const { data } = response
          this.list = data.list
          this.page.total = data.total
          this.listLoading = false
        })
      } else {
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
      }
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
    handleDelete(id) {
      this.$confirm('确认删除这条记录吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          if (this.siteId === 0) {
            globalIp.deleteById(id).then((response) => {
              this.$message({
                type: response.code === 0 ? 'success' : 'error',
                message: '删除成功!'
              })
              this.fetchData()
            })
          } else {
            deleteById(id).then((response) => {
              this.$message({
                type: response.code === 0 ? 'success' : 'error',
                message: '删除成功!'
              })
              this.fetchData()
            })
          }
        })
        .catch(() => {
        })
    },
    handleEdit(id) {
      if (this.siteId === 0) {
        globalIp.getById(id).then((response) => {
          this.edit.formData = response.data.item
        })
      } else {
        getById(id).then((response) => {
          this.edit.formData = response.data.item
        })
      }
      this.edit.title = '编辑'
      this.edit.visible = true
    }
  }
}
</script>
