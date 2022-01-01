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

  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import PanelGroup from './components/PanelGroup'
import * as site from '@/api/site'

export default {
  name: 'Dashboard',
  components: { PanelGroup },
  computed: {
    ...mapGetters([
      'name'
    ])
  },
  data() {
    return {
      eventTotal: 1000,
      allowedTotal: 1000,
      deniedTotal: 1000,
      unknownTotal: 1000,

      flag: false, // 判断是否显示图表组件
      categoryTotal: {}, // 每个分类下的文章数
      options: []
    }
  },
  created() {
    // this.getEventInfo(null, 0, 0)
    this.fetchData()
    this.flag = true
  },
  methods: {
    async fetchData() {
      const { data } = await site.getList({}, 0)
      const sites = data.list
      for (const record of sites) {
        const item = {}
        item.label = record.name
        item.value = record.id
        this.options.push(item)
      }
      const start = new Date().getTime() - 3600 * 1000 * 24 * 7
      const end = new Date().getTime()
      const { info } = await this.queryData(0, start ,end)
      console.log(info)
    },

    async queryData(id, start, end) {
      const query = {}
      query.start = start
      query.end = end
      query.site = id

      await site.getInfo(query)
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
