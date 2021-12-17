<template>
  <div
    v-permission="['GET:/shenshu/specificgroup/:id/rule', 'GET:/shenshu/specificgroup/rule/:id']"
    class="app-container"
  >
    <el-form :inline="true" :model="query" size="mini">
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
          v-permission="['POST:/shenshu/specificgroup/:id/rule']"
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
      <el-table-column prop="name" label="规则名称" />
      <el-table-column align="left" prop="rules" label="规则内容" width="200">
        <template slot-scope="scope">
          <div v-for="(item, index) in scope.row.rules" :key="index">
            <el-input v-if="item.variable === 'REQ_HEADER'" :value="VARIABLES_TEXT[item.variable] + ':' + item.header + ' ' + OPERATORS_TEXT[item.operator] + ' ' + item.pattern" size="mini" />
            <el-input v-else type="success" size="mini" :value="VARIABLES_TEXT[item.variable] + ' ' + OPERATORS_TEXT[item.operator] + ' ' + item.pattern" />
          </div>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="action" label="匹配动作">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.action === 1" type="success">允许</el-tag>
          <el-tag v-if="scope.row.action === 2" type="danger">阻断</el-tag>
          <el-tag v-if="scope.row.action === 4" type="primary">日志</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="status" label="状态">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === 1" type="success">启用</el-tag>
          <el-tag v-if="scope.row.status === 2" type="danger">停用</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="priority" label="优先级" />
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
            v-permission="['PUT:/shenshu/specificgroup/rule/:id']"
            type="success"
            size="mini"
            @click="handleEdit(scope.row.id)"
          >编辑</el-button>
          <el-button
            v-permission="['DELETE:/shenshu/specificgroup/rule/:id']"
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
      :id="edit.id"
      :title="edit.title"
      :form-data="edit.formData"
      :visible="edit.visible"
      :remote-close="remoteClose"
    />

  </div>
</template>

<script>
import { getList, deleteById, getById } from '@/api/rulespecific'
import Edit from './edit'
import { ACTION_TYPES, OPERATORS_TEXT, VARIABLES_TEXT } from '@/utils/rule'

export default {
  name: 'SpecificRule',
  components: { Edit },
  data() {
    return {
      ACTION_TYPES,
      OPERATORS_TEXT,
      VARIABLES_TEXT,

      groupId: 0,
      group: [],
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
      listLoading: false
    }
  },
  watch: {
    '$route.path': {
      immediate: true,
      handler() {
        const id = this.$route.params.groupId
        if (id === undefined) {
          this.fetchData()
        } else {
          this.groupId = id
        }
      }
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    async fetchData() {
      this.listLoading = true
      getList(
        this.groupId,
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
      this.groupId = id
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
    openAdd() {
      this.edit.id = this.groupId
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
        this.edit.id = id
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
    handleClose() {
      this.remoteClose()
    }
  }
}
</script>

<style scoped>
::v-deep .input_size .el-input { width: 250px; }

</style>
