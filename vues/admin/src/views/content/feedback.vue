<template>
  <div>
    <el-form
      style="text-align: right"
      :model="form"
      ref="queryData"
      :inline="true"
      size="normal"
    >
      <el-form-item>
        <el-date-picker
          v-model="queryData.range_date"
          type="daterange"
          start-placeholder="开始时间"
          end-placeholder="结束时间"
          range-separator="到"
          format="YYYY-MM-DD"
          value-format="YYYY-MM-DD"
          style="width: 220px"
        />
        <el-button plain @click="loadData(1)">搜索</el-button>
      </el-form-item>
    </el-form>
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
      <el-table-column prop="replay_content" label="回复内容" />
      <el-table-column prop="create_at" label="注册时间" />
      <el-table-column prop="create_ip" label="注册IP" />
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
      @current-change="loadData"
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
        id: 0,
        key: "",
        page: 1,
        page_size: 20,
        range_date: ["", ""],
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
      const data = await this.$api("/feedback/list", this.queryData);
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
