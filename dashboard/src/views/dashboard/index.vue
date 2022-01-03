<template>
  <div class="dashboard-container">
    <panel-group
      :event-total="eventTotal"
      :allowed-total="allowedTotal"
      :denied-total="deniedTotal"
      :unknown-total="unknownTotal"
      :options="options"
      :query-data="queryData"
    />
    <el-row :gutter="40">
      <el-col :xs="24" :sm="24" :lg="12">
        <el-card>
          <EventPieChart
            v-if="flag"
            ref-name="batchPie"
            title="batch事件分布"
            :data="batchEvents"
          />
        </el-card>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="12">
        <el-card>
          <EventPieChart
            v-if="flag"
            ref-name="specificPie"
            title="specific事件分布"
            :data="specificEvents"
          />
        </el-card>
      </el-col>
    </el-row>

    <el-row style="margin-top:30px">
      <el-card>
        <EventLineChart
          ref-name="batch"
          title="batch事件分布"
          :data="batchEvents"
        />
      </el-card>
    </el-row>
    <el-row style="margin-top:30px">
      <el-card>
        <EventLineChart
          ref-name="specific"
          title="specific事件分布"
          :data="specificEvents"
        />
      </el-card>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import PanelGroup from './components/PanelGroup'
import EventLineChart from './components/EventLineChart'
import EventPieChart from './components/EventPieChart'
import * as api from '@/api/site'

export default {
  name: 'Dashboard',
  components: { PanelGroup, EventLineChart, EventPieChart },
  data() {
    return {
      eventTotal: 1000,
      allowedTotal: 1000,
      deniedTotal: 1000,
      unknownTotal: 1000,

      flag: false, // 判断是否显示图表组件
      categoryTotal: {}, // 每个分类下的文章数
      options: [],
      batchEvents: {},
      specificEvents: {}
    }
  },
  computed: {
    ...mapGetters([
      'name'
    ])
  },
  created() {
    // this.getEventInfo(null, 0, 0)
    this.fetchData()
    this.flag = true
  },
  methods: {
    async fetchData() {
      let response = null
      response = await api.getList({}, 0)
      const sites = response.data.list
      for (const record of sites) {
        const item = {}
        item.label = record.name
        item.value = record.id
        this.options.push(item)
      }
      const start = new Date().getTime() - 3600 * 1000 * 24 * 7
      const end = new Date().getTime()
      response = await this.queryData(0, start, end)

      this.batchEvents = {}
      this.batchEvents.items = response.data.item.batch
      this.batchEvents.start = start
      this.batchEvents.end = end

      this.specificEvents = {}
      this.specificEvents.items = response.data.item.specific
      this.specificEvents.start = start
      this.specificEvents.end = end
    },

    queryData(id, start, end) {
      const query = {}
      query.start = start
      query.end = end
      query.site = id

      return api.getInfo(query)
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
