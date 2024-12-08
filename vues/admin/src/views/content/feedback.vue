<template>
  <div>
    <el-tabs
      v-model="queryData.tab"
      type="card"
      tab-position="top"
      @tab-change="loadData(1)"
    >
      <el-tab-pane label="未处理" name="noreply"> </el-tab-pane>
      <el-tab-pane label="已回复" name="replyed"> </el-tab-pane>
      <el-tab-pane label="全部" name="all"> </el-tab-pane>
    </el-tabs>

    <el-form
      style="text-align: right"
      :model="queryData"
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
      <el-table-column label="图片">
        <template #default="scope">
          <Imgs v-model="scope.row.imgs" />
        </template>
      </el-table-column>

      <el-table-column prop="content" label="内容" />
      <el-table-column prop="replay_content" label="回复内容" />
      <el-table-column prop="create_at" label="时间" />
      <el-table-column prop="create_ip" label="IP" />
      <el-table-column label="操作" fixed="right">
        <template #default="scope">
          <el-button
            v-if="scope.row.is_replay == 0"
            link
            size="small"
            @click.prevent="reply(scope.row)"
          >
            回复
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
        tab: "noreply",
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
    reply(row) {
      this.$prompt("回复反馈", {
        confirmButtonText: "回复",
        cancelButtonText: "取消",
        inputPlaceholder: "请输入答复内容",
        inputValidator: (v) => {
          if (!v || v == "" || v.length < 2) {
            return "请输入回复内容";
          }
          return true;
        },
      })
        .then(({ value }) => {
          this.$api("/feedback/send", {
            id: row.id,
            content: value,
          }).then((r) => {
            if (r.code == 200) {
              this.$message.success("成功");
              this.loadData();
            } else {
              this.$message.error(r.msg);
            }
          });
          console.log("c", value);
        })
        .catch((e) => {
          console.log(e);
        });
    },
    async loadData(page) {
      if (page) {
        this.queryData.page = page;
      }
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
