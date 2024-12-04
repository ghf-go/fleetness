<template>
  <div>
    <el-table
      :data="tableData"
      border
      stripe
      style="width: 100%"
      :row-class-name="$tableRowClassName"
    >
      <el-table-column prop="id" label="ID" />
      <el-table-column label="账号">
        <template #default="scope">
          <UserInfo v-model="scope.row.user_info" />
        </template>
      </el-table-column>
      <el-table-column prop="content" label="内容" />
      <el-table-column prop="user_id" label="用户" />
      <el-table-column prop="create_at" label="时间" />
      <el-table-column prop="create_ip" label="IP" />
      <el-table-column label="操作" fixed="right">
        <template #default="scope">
          <el-button
            link
            type="primary"
            size="small"
            @click.prevent="deleteRow(scope.$index)"
          >
            Remove
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      style="margin-top: 1rem"
      v-model:current-page="queryData.page"
      v-model:page-size="queryData.page_size"
      :page-sizes="[10, 20, 50, 100]"
      layout="total, sizes, prev, pager, next, jumper"
      :total="totalRow"
      background
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    />
  </div>
</template>

<script>
export default {
  data() {
    return {
      tableData: [],
      totalRow: 0,
      queryData: {
        page: 1,
        page_size: 20,
      },
    };
  },
  mounted() {
    this.loadData();
  },
  methods: {
    handleSizeChange(pp) {
      this.queryData.page_size = pp;
      this.queryData.page = 1;
      this.loadData();
    },
    handleCurrentChange(pp) {
      this.queryData.page = pp;
      this.loadData();
    },
    async loadData() {
      const data = await this.$api("/comment/list", this.queryData);
      if (data.code != 200) {
        this.$message.error(data.msg);
      } else {
        this.tableData = data.data.list;
        this.totalRow = data.data.total;
      }
    },
  },
};
</script>

<style lang="scss" scoped></style>
