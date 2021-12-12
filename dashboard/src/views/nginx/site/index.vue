<template>
  <div class="app-container">
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
          v-if="ids === null"
          icon="el-icon-circle-plus-outline"
          type="primary"
          @click="openAdd"
        >新增</el-button>
        <el-button
          v-if="ids !== null"
          icon="el-icon-circle-plus-outline"
          type="success"
          @click="ipdateIpSites"
        >关联站点</el-button>
      </el-form-item>
    </el-form>
    <el-table
      ref="dataTable"
      :data="list"
      stripe
      border
      style="width: 100%"
      row-key="id"
      @selection-change="handleSelectionChange"
    >
      <el-table-column
        v-if="ids !== null"
        align="center"
        reserve-selection
        type="selection"
        width="55"
      />
      <el-table-column align="center" type="index" label="序号" width="60px" />
      <el-table-column align="center" prop="name" label="名称" width="150px" />
      <el-table-column align="center" prop="host" label="域名" width="200px" />
      <el-table-column align="center" prop="path" label="路径" width="150px" />
      <el-table-column align="center" prop="upstreamRef" label="Upstream" width="100px">
        <template v-if="scope.row.upstreamRef.length === 1" slot-scope="scope">
          {{ scope.row.upstreamRef[0].name }}
        </template>
      </el-table-column>
      <el-table-column align="center" prop="remark" label="备注" width="200px" />
      <el-table-column v-if="ids === null" align="center" label="操作">
        <template slot-scope="scope">
          <el-button
            type="success"
            size="mini"
            @click="handleEdit(scope.row.id)"
          >编辑</el-button>
          <el-button
            type="primary"
            size="mini"
            @click="handleDelete(scope.row.id)"
          >规则管理</el-button>
          <el-button
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

  </div>
</template>

<script>
import * as api from '@/api/site'
import Edit from './edit'

export default {
  name: 'Site',
  components: { Edit },
  props: {
    ids: {
      type: Array,
      default: function() { return null }
    }
  },
  data() {
    return {
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
      checkedSitesList: []
    }
  },
  watch: {
    ids() {
      this.query = {}
      this.queryData()
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

      this.chekedSites()
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
              type: response.code === 200 ? 'success' : 'error',
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
    chekedSites() {
      this.$refs.dataTable.clearSelection()
      if (this.ids) {
        this.list.forEach((item) => {
          if (this.ids.indexOf(item.id) !== -1) {
            this.$refs.dataTable.toggleRowSelection(item, true)
          }
        })
      }
    },
    handleSelectionChange(val) {
      this.checkedSitesList = val
    },
    ipdateIpSites() {
      const checkedSites = []
      this.checkedSitesList.forEach((item) => {
        checkedSites.push(item.id)
      })

      this.checkedSitesList = []
      this.$emit('updateIpSites', checkedSites)
    }
  }
}
</script>
