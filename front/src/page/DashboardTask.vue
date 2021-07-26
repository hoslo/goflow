<template>
  <el-container>
    <el-header>
      <el-menu
          :default-active="1"
          class="el-menu-horizontal"
          mode="horizontal"
          background-color="#23262E"
          text-color="#fff"
          active-text-color="#ffd04b"
          router="true">
        <el-menu-item index="/dag_list">任务中心</el-menu-item>
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
          <Task :key="timer"></Task>
        </el-main>
        <el-footer>Footer</el-footer>
      </el-container>
    </el-container>
  </el-container>
</template>

<script>
import Task from "@/components/Task";
export default {
  name: "DashboardDag",
  components: {Task},
  mounted() {

    return {}
  },
  data() {
    return {}
  },
  methods: {
    createTask() {
      let that = this;
      console.log(this.task)
      this.axios.post("http://127.0.0.1:8000/create_task", {
        name: that.task.name,
        is_root: that.task.is_root,
        dag_id: Number(that.$route.params.id),
        max_try_count: Number(that.task.max_try_count),
        wait_time: Number(that.task.wait_time),
        operator_type: that.task.operator_type,
        operator_id: that.task.operator_id,
        parent_id: that.task.parent_id
      }, {
        headers: {
          token: localStorage.token
        },
      })
      this.dialogFormVisible = false;
      this.timer = new Date().getTime();
    },
    openCreateTask () {
      this.dialogFormVisible = true;
      let that = this;
      this.axios.get("http://127.0.0.1:8000/get_httpoperators", {
        headers: {
          token: localStorage.token
        }
      }).then(function (response) {
        console.log(response)
        that.httpoperators = response.data.data.httpoperators;
      })
      this.axios.get("http://127.0.0.1:8000/get_tasks", {
        headers: {
          token: localStorage.token
        },
        params: {
          id: that.$route.params.id
        }
      }).then(function (response) {
        console.log(response)
        that.tasks = response.data.data.tasks;
      })
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

.el-header {
  background-color: #23262E;
  color: #333;
  text-align: center;
  line-height: 60px;
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