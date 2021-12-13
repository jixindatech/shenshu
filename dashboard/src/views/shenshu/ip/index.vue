<template>
  <div
    class="app-container"
  >
    <el-form :inline="true" size="mini">
      <el-form-item label="域名:">
        <el-select v-model="siteId" placeholder="请选择域名" @change="selectChanged">
          <el-option v-for="(item,index) in sites" :key="index" :label="item.name" :value="item.id" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button
          icon="el-icon-search"
          type="primary"
          @click="queryData"
        >查询
        </el-button>
      </el-form-item>
    </el-form>
    <el-tabs v-model="activeName" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="Allow名单" name="1">
        <Item :params="activeName" :site-id="siteId" :type="IP_TYPE.ALLOW" />
      </el-tab-pane>
      <el-tab-pane label="Deny名单" name="2">
        <Item :params="activeName" :site-id="siteId" :type="IP_TYPE.DENY" />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import * as site from '@/api/site'
import { IP_TYPE } from '@/utils/const'
import Item from './components/ip.vue'

export default {
  name: 'IP',
  components: { Item },
  data() {
    return {
      siteId: 1,
      sites: [],
      IP_TYPE,
      activeName: '1'
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
    },
    selectChanged(id) {
      this.site = id
    },
    queryData() {
      console.log('query data')
    },
    handleClick(tab, event) {
      this.activeName = tab.name
    }
  }
}
</script>
