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
          <el-row>
            <el-button type="primary" @click="openCreateTask">创建Task</el-button>

            <el-dialog title="Task" v-model="dialogFormVisible" width="30%">
              <el-form :model="task">
                <el-form-item label="名称"
                              :rules="[
                      { required: true, message: '名称不能为空'},
                  ]">
                  <el-input v-model="task.name"></el-input>
                </el-form-item>
                <el-form-item label="Root">
                  <el-switch v-model="task.is_root"></el-switch>
                </el-form-item>
                <el-form-item label="父任务">
                  <el-select
                      v-model="parent_tasks"
                      multiple
                      clearable
                      placeholder="请选择类型查询">
                    <el-option v-for="(type,ind) in tasks"
                               :key=ind
                               :label="type.name"
                               :value="type.ID">
                    </el-option>
                  </el-select>
                </el-form-item>
                <el-form-item label="最大重试次数">
                  <el-select v-model="task.max_try_count" placeholder="请选择最大重试次数">
                    <el-option label="1次" value="1"></el-option>
                    <el-option label="2次" value="2"></el-option>
                    <el-option label="3次" value="3"></el-option>
                  </el-select>
                </el-form-item>
                <el-form-item label="等待时间"
                              placeholder="请输入等待时间,单位为秒"
                              :rules="[
                              { type: 'number', message: '等待时间为数字值'}
                            ]">
                  <el-input v-model="task.wait_time"></el-input>
                </el-form-item>
                <el-form-item label="HttpOperator">
                  <el-select v-model="task.operator_id" placeholder="请选择Operator">
                    <el-option v-for="operator in httpoperators" :key="operator.ID" :label="operator.name" :value="operator.ID"></el-option>
                  </el-select>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="createTask">立即创建</el-button>
                  <el-button @click="dialogFormVisible = false">取消</el-button>
                </el-form-item>
              </el-form>
            </el-dialog>
          </el-row>
          <Dag :key="timer"></Dag>
        </el-main>
        <el-footer>Footer</el-footer>
      </el-container>
    </el-container>
  </el-container>
</template>

<script>
import Dag from "@/components/Dag";
export default {
  name: "DashboardDag",
  components: {Dag},
  mounted() {

    return {}
  },
  data() {
    return {
      dialogFormVisible: false,
      timer: "",
      form: {
        name: ""
      },
      formLabelWidth: '80px',
      task: {
        name: '',
        is_root: 0,
        max_try_count: 0,
        wait_time: 0,
        operator_id: 0,
        operator_type: "httpoperator",
        parent_id: 0
      },
      httpoperators: [
        {
          id: "请选择",
          name: ""
        }
      ],
      tasks: [
        {
          id: "请选择",
          name: ""
        }
      ],
      parent_tasks: [],
      old_parent_tasks: [],
      rules: {
        region: [
          { required: true, message: '请选择活动区域', trigger: 'change' }
        ],
      }
    }
  },
  methods: {
    createTask() {
      let that = this;
      console.log(this.task)
      console.log(this.parent_tasks)
      this.axios.post("http://127.0.0.1:8000/create_task", {
        name: that.task.name,
        is_root: that.task.is_root,
        dag_id: Number(that.$route.params.id),
        max_try_count: Number(that.task.max_try_count),
        wait_time: Number(that.task.wait_time),
        operator_type: that.task.operator_type,
        operator_id: that.task.operator_id,
        parent_tasks: that.parent_tasks
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
    },

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