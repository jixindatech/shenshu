<template>
  <div :class="{'has-logo':showLogo}">
    <logo v-if="showLogo" :collapse="isCollapse" />
    <el-scrollbar wrap-class="scrollbar-wrapper">
      <el-menu
        :default-active="activeMenu"
        :collapse="isCollapse"
        :background-color="variables.menuBg"
        :text-color="variables.menuText"
        :unique-opened="false"
        :active-text-color="variables.menuActiveText"
        :collapse-transition="false"
        mode="vertical"
      >
        <sidebar-item v-for="route in routes" :key="route.path" :item="route" :base-path="route.path" />
      </el-menu>
    </el-scrollbar>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Logo from './Logo'
import SidebarItem from './SidebarItem'
import variables from '@/styles/variables.scss'

export default {
  components: { SidebarItem, Logo },
  computed: {
    ...mapGetters([
      'sidebar'
    ]),
    routes() {
      return this.changeRoutes(this.$router.options.routes)
    },
    activeMenu() {
      const route = this.$route
      const { meta, path } = route
      // if set path, the sidebar will highlight the path you set
      if (meta.activeMenu) {
        return meta.activeMenu
      }
      return path
    },
    showLogo() {
      return this.$store.state.settings.sidebarLogo
    },
    variables() {
      return variables
    },
    isCollapse() {
      return !this.sidebar.opened
    }
  },
  methods: {
    mergeRoutes(router1, router2) {
      var newRouters = []
      if (router1.length > 0 && router2.length > 0) {
        for (let i = 0; i < router1.length; i++) {
          for (let j = 0; j < router2.length; j++) {
            if (router1[i].path === router2[j].path) {
              var item = {
                path: router2[j].path
              }

              if (router2[j].name !== undefined) {
                item['name'] = router2[j].name
              }
              if (router2[j].hidden === true) {
                item['hidden'] = true
              }

              if (router2[j].meta !== undefined) {
                item['meta'] = router2[j].meta
              }

              if (router2[j].redirect !== undefined) {
                item['redirect'] = router2[j].redirect
              }

              item['children'] = []
              if (router1[i].children !== undefined && router2[j].children !== undefined) {
                var router = this.mergeRoutes(router1[i].children, router2[j].children)
                if (router.length > 0) {
                  item['children'] = router
                }
              }
              if (item['children'].length === 0) {
                delete item['children']
              }

              newRouters.push(item)
            }
          }
        }
      }

      return newRouters
    },

    changeRoutes(routes) {
      const router = this.$store.getters.routes
      if (router.length === 0) {
        return []
      }

      var newRoutes = this.mergeRoutes(router, routes)
      return newRoutes
    }
  }
}
</script>
