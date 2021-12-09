<template>
  <div
    v-permission="['GET:/system/user', 'GET:/system/user/:id']"
    class="app-container"
  >
    <el-form :inline="true" :model="query" size="mini">
      <el-form-item label="用户名:">
        <el-input v-model.trim="query.username" />
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
          v-if="msgId == 0"
          v-permission="['POST:/system/user']"
          icon="el-icon-circle-plus-outline"
          type="primary"
          @click="openAdd"
        >新增</el-button>
        <el-button
          v-if="msgId != 0"
          icon="el-icon-circle-plus-outline"
          type="success"
          @click="handleUserMsg"
        >群发短信</el-button>
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
      @selection-change="handleSelectionChange"
    >
      <el-table-column
        v-if="msgId != 0"
        align="center"
        reserve-selection
        type="selection"
        width="55"
      />
      <el-table-column prop="username" label="用户名" />
      <el-table-column prop="phone" label="手机号" width="150" />
      <el-table-column prop="email" label="邮箱" width="180" />
      <el-table-column prop="role" label="角色" />
      <el-table-column prop="loginType" label="认证方式" />
      <el-table-column v-if="msgId == 0" prop="createdAt" label="创建时间" width="220">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.createdAt }}</span>
        </template>
      </el-table-column>
      <el-table-column v-if="msgId == 0" prop="updateAt" label="更新时间" width="220">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.updateAt }}</span>
        </template>
      </el-table-column>
      <el-table-column v-if="msgId == 0" align="center" label="操作" width="250">
        <template slot-scope="scope">
          <el-button
            v-permission="['PUT:/system/user']"
            type="success"
            size="mini"
            @click="handleEdit(scope.row.id)"
          >编辑</el-button>
          <el-button
            v-permission="['PUT:/system/user/password/:id']"
            type="danger"
            size="mini"
            :disabled="scope.row.id === 1"
            @click="handleResetPassword(scope.row.id)"
          >密码重置</el-button>
          <el-button
            v-permission="['DELETE:/system/user/:id']"
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
      v-permission="['POST:/system/user', 'PUT:/system/user/:id']"
      :title="edit.title"
      :form-data="edit.formData"
      :visible="edit.visible"
      :remote-close="remoteClose"
    />

  </div>
</template>

<script>
import { getList, deleteById, getById, resetUserPasswordById } from '@/api/user'
import Edit from './edit'
export default {
  components: { Edit },
  props: {
    msgId: {
      type: Number,
      default: 0
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
      checkedUserList: []
    }
  },
  watch: {
    msgId(newVal, oldVal) {
      if (newVal !== 0) {
        this.query = {}
        this.queryData()
      }
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
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
      this.edit.title = '新增用户'
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
    handleResetPassword(id) {
      this.$confirm('确认重置密码吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          resetUserPasswordById(id).then((response) => {
            this.$message({
              type: 'success',
              message: '重置成功!'
            })
          })
        })
        .catch(() => {
          this.$message({
            type: 'error',
            message: '重置失败!'
          })
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
    handleSelectionChange(val) {
      this.checkedUserList = val
    },
    handleUserMsg() {
      const checkedUserIds = []
      this.checkedUserList.forEach((item) => {
        checkedUserIds.push(item.id)
      })

      this.$emit('sendUserMsg', checkedUserIds)
      this.checkedUserList = []
    }
  }
}
</script>
