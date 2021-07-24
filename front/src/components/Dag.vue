
<template>
  <div></div>
  <el-table
      :data="tableData"
      border
      fit=true
      style="width: 100%"
      class="el-table">
    <el-table-column
        prop="ID"
        label="ID"
    >
    </el-table-column>
    <el-table-column
        prop="Name"
        label="名称"
    >
    </el-table-column>
    <el-table-column
        prop="CreatedAt"
        label="创建时间"
    >
    </el-table-column>
    <el-table-column
        prop="UpdatedAt"
        label="更新时间"
    >
    </el-table-column>
    <el-table-column
        label="操作"
    >
      <template #default="scope">
        <el-button @click="handleClick(scope.row)" type="text" size="small">查看</el-button>
        <el-button type="text" size="small">编辑</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>

export default {
  mounted() {
    console.log(this.$route.params.id)

    this.getDAG();

    return {}
  },
  methods: {
    handleClick(row) {
      console.log(row);
    },

    getDAG() {
      let that = this;
      this.axios.get("http://127.0.0.1:8000/get_tasks", {
        headers: {
          token: localStorage.token
        },
        params: {
          id: that.$route.params.id
        }
      }).then(function (response) {
        console.log(response)
        that.tableData = response.data.data.tasks;
      })
    }
  },

  data() {
    return {
      tableData: []
    }
  },
}
</script>

<style>
.el-table {
  padding: 0;
  margin: 0;
}
</style>
