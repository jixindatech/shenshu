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
            type="success"
            size="mini"
            @click="relateIPWithSites(scope.row.id)"
          >关联域名</el-button>
          <el-button
            type="danger"
            size="mini"
            :disabled="scope.row.id === 1"
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
      :form-data="edit.formData"
      :type="type"
      :visible="edit.visible"
      :remote-close="remoteClose"
    />

    <el-dialog title="关联站点" :visible.sync="site.visible" width="65%">
      <Site :ids="site.ids" @updateIpSites="updateIpSites" />
    </el-dialog>

  </div>
</template>

<script>
import { getList, deleteById, getById, update } from '@/api/ip'
import Edit from './edit'
import Site from '@/views/nginx/site'

export default {
  components: { Edit, Site },
  props: {
    params: {
      type: String,
      default: '1'
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
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.query.type = Number(this.type)
      this.listLoading = true
      getList(
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
    },
    relateIPWithSites(id) {
      this.site.ip = id
      this.site.ids = []
      this.site.visible = true

      getById(id).then((response) => {
        console.log(response)
        const ip = response.data.item
        ip.sites.forEach(element => {
          this.site.ids.push(element.id)
        })
        this.site.visible = true
      })
    },
    updateIpSites(ids) {
      const data = { sites: ids }
      update(this.site.ip, data).then((response) => {
        if (response.code === 0) {
          this.$message({ message: '关联成功', type: 'success' })
          this.site.visible = false
        }
      })
    }
  }
}
</script>
