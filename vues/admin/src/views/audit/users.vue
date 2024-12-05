<template>
  <div>
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

      <el-table-column prop="ukey" label="字段" />
      <el-table-column prop="newval" label="修改内容" />
      <el-table-column prop="uval" label="之前改前" />
      <el-table-column prop="create_at" label="注册时间" />
      <el-table-column prop="create_ip" label="注册IP" />
      <el-table-column label="操作" fixed="right">
        <template #default="scope">
          <el-button
            link
            type="success"
            size="small"
            @click.prevent="audit(scope.row, 'accept')"
          >
            通过
          </el-button>
          <el-button
            link
            type="danger"
            size="small"
            @click.prevent="audit(scope.row, 'del')"
          >
            拒绝
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
      addUserData: {
        isShowAddUser: false,
        name: "",
        pass: "",
      },
    };
  },
  mounted() {
    this.loadData();
  },
  methods: {
    async audit(r, act) {
      const data = await this.$api("/account/user_audit", {
        uid: r.user_id,
        key: r.ukey,
        act: act,
      });
      if (data.code == 200) {
        this.$message.success("操作完成");
        this.tableData = this.tableData.filter((item) => {
          return item.id != r.id;
        });
        if (this.tableData.length == 0) {
          this.loadData();
        }
      } else {
        this.$message.error(data.msg);
      }
    },
    handleSizeChange(pp) {
      this.queryData.page_size = pp;
      this.queryData.page = 1;
      this.loadData();
    },

    async loadData(page) {
      if (page) {
        this.queryData.page = page;
      }
      const data = await this.$api("/account/user_wait_audit", this.queryData);
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
