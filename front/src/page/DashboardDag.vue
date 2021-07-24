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
            active-text-color="#ffd04b">
          <el-menu-item index="3">任务中心</el-menu-item>
          <el-menu-item index="4">用户管理</el-menu-item>
        </el-menu>
      </el-aside>
      <el-container>
        <el-main>
          <Dag></Dag>
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
  components: {Dag}
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