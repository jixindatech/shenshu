<template>
  <div
    class="app-container"
  >
    <el-tabs v-model="activeName" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="Allow名单" name="1">
        <Allow :params="activeName" :type="IP_TYPE.ALLOW" />
      </el-tab-pane>
      <el-tab-pane label="Deny名单" name="2">
        <Deny :params="activeName" :type="IP_TYPE.DENY" />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import { IP_TYPE } from '@/utils/const'
import Allow from './components/allow.vue'
import Deny from './components/deny.vue'

export default {
  name: 'IP',
  components: { Allow, Deny },
  data() {
    return {
      IP_TYPE,
      activeName: '1'
    }
  },
  methods: {
    handleClick(tab, event) {
      this.activeName = tab.name
    }
  }
}
</script>
