<template>
  <div
    v-permission="[
      'GET:/nginx/site',
      'GET:/shenshu/event/cc']"
    class="app-container"
  >
    <el-form :inline="true" :model="query" size="mini">
      <el-form-item>
        <el-date-picker
          v-model="queryTime"
          type="datetimerange"
          :picker-options="pickerOptions"
          range-separator="-"
          start-placeholder=""
          end-placeholder=""
          value-format="timestamp"
          align="right"
        />
      </el-form-item>
      <el-form-item label="站点名称:">
        <el-select v-model="query.site" clearable placeholder="全部站点">
          <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="站点:">
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
      </el-form-item>
    </el-form>

    <el-table
      :data="list"
      stripe
      border
      style="width: 100%"
    >
      <el-table-column align="center" type="index" label="序号" width="60" />
      <el-table-column align="center" prop="_source.id" label="命中ID" />
      <el-table-column align="center" prop="_source.host" label="域名" />
      <el-table-column align="center" prop="_source.method" label="请求方法" />
      <el-table-column align="center" prop="_source.uri" label="Uri" />
      <el-table-column align="center" prop="_source.time" label="时间">
        <template slot-scope="scope">
          {{ new Date(scope.row._source.timestamp * 1000).toLocaleString() }}
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

  </div>
</template>

<script>
import * as site from '@/api/site'
import * as api from '@/api/ccevent'

export default {
  name: 'CCEvent',
  data() {
    return {
      list: [],
      page: {
        current: 1,
        size: 20,
        total: 0
      },
      options: [],
      query: {},
      queryTime: [],
      pickerOptions: {
        shortcuts: [{
          text: '最近30分钟',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 1800 * 1000)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '最近一小时',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '最近24小时',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24)
            picker.$emit('pick', [start, end])
          }
        },
        {
          text: '最近一周',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
            picker.$emit('pick', [start, end])
          }
        },
        {
          text: '最近一个月',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
            picker.$emit('pick', [start, end])
          }
        }]
      }
    }
  },

  created() {
    this.queryTime[0] = new Date().getTime() - 3600 * 1000 * 24 * 7
    this.queryTime[1] = new Date().getTime()
    this.fetchData()
  },

  methods: {
    async fetchData() {
      let response = null
      response = await site.getList({}, 0)
      const sites = response.data.list
      this.options = []
      for (const record of sites) {
        const item = {}
        item.label = record.name
        item.value = record.id
        this.options.push(item)
      }

      if (this.queryTime.length > 0) {
        this.query['start'] = this.queryTime[0]
        this.query['end'] = this.queryTime[1]
      }
      const { data } = await api.getList(this.query, this.page.current, this.page.size)
      this.list = data.list
      this.page.total = data.count
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
      this.queryTime[0] = new Date().getTime() - 3600 * 1000 * 24 * 7
      this.queryTime[1] = new Date().getTime()

      this.fetchData()
    }
  }
}
</script>
