<template>
  <div
    class="app-container"
  >
    <el-form :inline="true" :model="query" size="mini">
      <el-form-item label="规则组名称:">
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
          v-if="!ids"
          icon="el-icon-circle-plus-outline"
          type="primary"
          @click="openAdd"
        >新增</el-button>
      </el-form-item>
      <el-form-item>
        <el-button
          v-if="ids"
          icon="el-icon-circle-plus-outline"
          type="success"
          @click="setRuleGroup"
        >设置规则组</el-button>
      </el-form-item>
    </el-form>

    <el-table
      ref="dataTable"
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      :header-cell-style="{'text-align':'center'}"
      :cell-style="{'text-align':'center'}"
      border
      fit
      highlight-current-row
      row-key="id"
      @selection-change="handleSelectionChange"
    >
      <el-table-column
        v-if="ids"
        align="center"
        reserve-selection
        type="selection"
        width="55"
      />
      <el-table-column prop="name" label="规则组名称" />
      <el-table-column prop="decoder" label="编码">
        <template slot-scope="scope">
          <div v-for="(item, index) in scope.row.decoder" :key="index">
            <el-input type="success" size="mini" :value="item" />
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="priority" label="优先级" />
      <el-table-column align="center" prop="status" label="状态">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === 1" type="success">启用</el-tag>
          <el-tag v-if="scope.row.status === 2" type="danger">停用</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="action" label="动作">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.action === 1" type="primary">日志短路</el-tag>
          <el-tag v-if="scope.row.action === 2" type="primary">日志全量</el-tag>
          <el-tag v-if="scope.row.action === 3" type="primary">短路</el-tag>
          <el-tag v-if="scope.row.action === 4" type="primary">全量</el-tag>
        </template>
      </el-table-column>
      <el-table-column v-if="!ids" prop="createdAt" label="创建时间" width="220">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.createdAt }}</span>
        </template>
      </el-table-column>
      <el-table-column v-if="!ids" prop="updateAt" label="更新时间" width="220">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.updateAt }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="remark" label="备注" />
      <el-table-column v-if="!ids" align="center" label="操作" width="250">
        <template slot-scope="scope">
          <el-button
            type="success"
            size="mini"
            @click="handleEdit(scope.row.id)"
          >编辑</el-button>
          <el-button
            type="primary"
            size="mini"
            @click="handleRule(scope.row.id)"
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

    <edit
      :title="edit.title"
      :data="edit.formData"
      :visible="edit.visible"
      :remote-close="remoteClose"
    />
  </div>
</template>

<script>
import { getList, deleteById, getById } from '@/api/batchgroup'
import Edit from './edit'
export default {
  name: 'BatchGroup',
  components: { Edit },
  props: {
    ids: {
      type: Array,
      default: function() { return null }
    }
  },
  data() {
    return {
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
      rule: {
        id: 0,
        title: '规则管理',
        visible: false
      },
      checkedList: []
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
      this.listLoading = true
      await getList(
        this.query,
        this.page.current,
        this.page.size
      ).then(response => {
        const { data } = response
        this.list = data.list
        this.page.total = data.total
        this.listLoading = false
      })

      this.chekedChoices()
    },
    chekedChoices() {
      this.$refs.dataTable.clearSelection()
      if (this.ids) {
        this.list.forEach((item) => {
          if (this.ids.indexOf(item.id) !== -1) {
            this.$refs.dataTable.toggleRowSelection(item, true)
          }
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
    handleEdit(id) {
      getById(id).then((response) => {
        const { data } = response
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
    },
    handleRule(id) {
      this.$router.push({ name: 'BatchRule', params: { groupId: id }})
    },
    handleSelectionChange(val) {
      this.checkedList = val
    },
    setRuleGroup() {
      const checkedIds = []
      this.checkedList.forEach((item) => {
        checkedIds.push(item.id)
      })

      this.$emit('getRuleGroup', checkedIds)
    }

  }
}
</script>

<style scoped>
::v-deep .input_size .el-input { width: 250px; }

</style>
