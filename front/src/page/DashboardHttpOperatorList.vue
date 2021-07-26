<template>
  <el-container>
    <el-header>
      <el-menu
          :default-active="1"
          class="el-menu-horizontal"
          mode="horizontal"
          background-color="#23262E"
          text-color="#fff"
          active-text-color="#ffd04b">
        <el-menu-item index="3">任务中心</el-menu-item>
        <el-menu-item index="4">用户管理</el-menu-item>
      </el-menu>
    </el-header>
    <el-container>
      <el-aside style="width: 200px">
        <el-menu
            :default-active="1"
            class="el-menu-vertical"
            mode="vertical"
            background-color="#333333"
            text-color="#fff"
            active-text-color="#ffd04b"
            router="true">
          <el-menu-item index="/dag_list">任务中心</el-menu-item>
          <el-submenu index="/httpoperator_list">
            <template #title>
              <span>Operator管理</span>
            </template>
            <el-menu-item-group>
              <el-menu-item index="/httpoperator_list">HttpOperator管理</el-menu-item>
            </el-menu-item-group>
          </el-submenu>
        </el-menu>
      </el-aside>
      <el-container>
        <el-main>
          <el-row>
            <el-button type="primary" @click="dialogFormVisible = true">创建HttpOperator</el-button>

            <el-dialog title="DAG" v-model="dialogFormVisible" width="30%">
              <el-form :model="form">
                <el-form-item label="名称" :label-width="formLabelWidth">
                  <el-input v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item label="URL" :label-width="formLabelWidth">
                  <el-input v-model="form.url"></el-input>
                </el-form-item>
              </el-form>
              <template #footer>
    <span class="dialog-footer">
      <el-button @click="dialogFormVisible = false">取 消</el-button>
      <el-button type="primary" @click="createHttpOperator">确 定</el-button>
    </span>
              </template>
            </el-dialog>
          </el-row>
          <HttpOperatorList :key="timer"></HttpOperatorList>
        </el-main>
        <el-footer>Footer</el-footer>
      </el-container>
    </el-container>
  </el-container>

</template>

<script>
import HttpOperatorList from "@/components/HttpOperatorList";
export default {
  components: {HttpOperatorList},
  data () {
    return {
      dialogFormVisible: false,
      timer: "",
      formLabelWidth: '80px',
      form: {
        name: "",
        url: "",
      }
    }
  },
  methods: {
    createHttpOperator() {
      let that = this;
      this.axios.post("http://127.0.0.1:8000/create_httpoperator", {
        name: that.form.name,
        url: that.form.url,
      }, {
        headers: {
          token: localStorage.token
        }
      })
      this.dialogFormVisible = false;
      this.timer = new Date().getTime();
    }
  }
}
</script>

<style scoped>
.el-menu-horizontal {
  padding: 0;
  margin: 0;
  width: 100%;
}

.el-menu-vertical {
  padding: 0;
  margin: 0;
  width: 100%;
}

.el-header {
  background-color: #23262E;
  color: #333;
  text-align: center;
  line-height: 60px;
  padding: 0;
}

.el-aside {
  background-color: #333333;
  color: #333;
  text-align: center;
  line-height: 200px;
}

.el-main {
  background-color: #E9EEF3;
  color: #333;
  text-align: center;
  line-height: 160px;
  margin: 0;
  padding: 0;
}

.el-footer {
  background-color: #E9EEF3;
  color: #333;
  text-align: center;
  line-height: 60px;
}

body > .el-container {
  margin-bottom: 0;
}

.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  line-height: 260px;
}

.el-container:nth-child(7) .el-aside {
  line-height: 320px;
}
</style>