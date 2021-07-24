<template>
  <el-row justify="center">
    <el-form ref="form" :model="form" label-width="80px" class="el-form">
        <el-form-item label="用户名">
          <el-input v-model="form.username"></el-input>
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="form.password"></el-input>
        </el-form-item>
        <el-form-item size="large">
          <el-button type="primary" @click="login">登录</el-button>
          <el-button type="primary" @click="registerUser">注册</el-button>
        </el-form-item>
    </el-form>
  </el-row>
</template>

<script>

export default {
  data() {
    return {
      form: {
        username: '',
        password: '',
      }
    }
  },
  methods: {
    registerUser() {
      let that = this;
      this.axios.post("http://127.0.0.1:8000/create_user", {
        username: this.form.username,
        password: this.form.password
      }).then(function (response) {
        console.log(response)
        that.$router.push({name: "dag_list"})
      })
    },
    login() {
      let that = this;
      this.axios.post("http://127.0.0.1:8000/get_user", {
        username: this.form.username,
        password: this.form.password
      }).then(function (response) {
        console.log(response.data.data.token)
        localStorage.token = response.data.data.token;
        that.$router.push({name: "dag_list"})
      })
    }
  }
}
</script>

<style>
  .el-form {
    margin-top: 15%;
  }
</style>

