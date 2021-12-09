<template>
  <div>
    <div class="app-header">
      <div class="logo">
        <a href="https://github.com/jixindatech/shenshu" title="梦学谷">
          <img src="@/assets/images/logo.png" height="50px">
        </a>
      </div>
    </div>

    <div class="login_page">
      <div class="login_box">
        <div class="center_box">
          <div :class="{login_form: true, rotate: tab == 2}">
            <div :class="{tabs: true, r180: reverse == 2}">
              <div class="fl tab" @click="changetab(1)">
                <span :class="{on: tab == 1}">登录</span>
              </div>
            </div>
            <div v-if="reverse == 1" class="form_body">
              <form @submit.prevent="loginSubmit">
                <input v-model="loginData.username" type="text" placeholder="请输入用户名" autocomplete="off">
                <input v-model="loginData.password" type="password" placeholder="请输入密码" autocomplete="off">
                <div class="error_msg">{{ loginMessage }}</div>
                <input v-if="subState" type="submit" disabled="disabled" value="登录中···" class="btn">
                <input v-else type="submit" value="登录" class="btn" @submit="loginSubmit">
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="mxg-footer">
      <div class="footer-info">
        Copyright &copy;All Rights Reserved&nbsp;
        <a href="https://github.com/jixindatech/shenshu" target="_blank" rel="nofollow">神荼后台管理系统</a>
      </div>
    </div>
  </div>
</template>
<script >

export default {

  data() {
    return {
      tab: 1,
      reverse: 1,
      loginMessage: '',
      regMessage: '',
      subState: false,
      loginData: {
        username: '',
        password: ''
      }
    }
  },
  watch: {
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    }
  },

  async created() {
    if (this.$route.query.redirectURL) {
      this.redirectURL = this.$route.query.redirectURL
    }
  },
  methods: {
    changetab(int) {
      this.tab = int
      setTimeout(() => {
        this.reverse = int
      }, 200)
    },

    loginSubmit() {
      if (this.subState) {
        return false
      }
      if (this.loginData.password.length < 6) {
        this.loginMessage = '请输入正确用户名'
        return false
      }

      if (this.loginData.password.length < 6) {
        this.loginMessage = '请输入正确的用户名或密码'
        return false
      }

      this.subState = true
      this.$store.dispatch('user/login', this.loginData).then((response) => {
        this.$router.push({ path: this.redirect || '/' })
        this.subState = false
      }).catch(() => {
        this.subState = false
        this.loginMessage = '系统繁忙，请稍后重试'
      })
    }
  }
}
</script>
<style scoped>
@import '../../assets/style/login.css';
.app-header {
    width: 100%;
    height: 80px;
    border-top: 3px solid #345dc2;
    z-index: 10;
}

.logo {
    width: 1200px;
    margin: 0 auto; /* 居中 */
    overflow: hidden;
    margin-top: 15px;
}
.mxg-footer {
  width: 1200px;
  margin: 0 auto; /* 居中 */
  line-height: 60px;
  border-top: 1px solid #ddd;
}
.footer-info {
  text-align: center;
  font-size: 13px;
  color: #2C2C40;
}
.footer-info a {
  color: #2C2C40;
  text-decoration: none;
}

</style>

