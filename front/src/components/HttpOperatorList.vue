<template>
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
        prop="name"
        label="名称"
    >
    </el-table-column>
    <el-table-column
        prop="url"
        label="URL"
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
    this.getDAGList()

    return {}
  },
  methods: {
    handleClick(row) {
      console.log(row);
      this.$router.push({name: "dag", params: {id: row.ID}})
    },

    getDAGList() {
      let that = this;
      this.axios.get("http://127.0.0.1:8000/get_httpoperators", {
        headers: {
          token: localStorage.token
        }
      }).then(function (response) {
        console.log(response)
        that.tableData = response.data.data.httpoperators;
      })
    }
  },

  data() {
    return {
      tableData: [],
      comName: [],
      ID: 0
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
