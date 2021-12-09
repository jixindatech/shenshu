<template>
  <div class="app-container">
    <el-tabs v-model="activeName" type="border-card" @tab-click="handleTabChange">
      <el-tab-pane v-permission="['GET:/system/email']" label="邮件配置" name="email">
        <Email :params="activeName" name="email" />
      </el-tab-pane>
      <el-tab-pane v-permission="['GET:/system/ldap']" label="Ldap配置" name="ldap">
        <Ldap :params="activeName" name="ldap" />
      </el-tab-pane>
      <!--
      <el-tab-pane label="Txsms配置" name="txsms">
        <Txsms :params="activeName" name="txsms" />
      </el-tab-pane>
      -->
    </el-tabs>
  </div>
</template>

<script>
import Email from './conponents/email'
import Ldap from './conponents/ldap'
// import Txsms from './conponents/txsms'
export default {
  components: { Email, Ldap },
  data() {
    return {
      activeName: 'email'
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
    },
    handleTabChange(tab, event) {
      this.activeName = tab.name
    }
  }
}
</script>
