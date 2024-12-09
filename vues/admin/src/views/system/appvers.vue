<template>
  <div>
    <el-button type="primary" size="default" @click="add">添加</el-button>

    <el-table
      :data="tableData"
      border
      stripe
      style="width: 100%"
      :row-class-name="$tableRowClassName"
    >
      <el-table-column prop="id" label="ID" />
      <el-table-column prop="app_ver" label="版本" />
      <el-table-column prop="apk_url" label="apk地址" />
      <el-table-column prop="wgt_url" label="热更新地址" />
      <el-table-column label="是否发布">
        <template #default="scope">
          <YesNo v-model="scope.row.is_online"></YesNo>
        </template>
      </el-table-column>
      <el-table-column prop="ver_content" label="版本说明" />
      <el-table-column label="操作" fixed="right">
        <template #default="scope">
          <el-button
            v-if="scope.row.is_online == 0"
            link
            type="primary"
            size="small"
            @click.prevent="publish(scope.row)"
          >
            发布
          </el-button>
          <el-button
            v-if="scope.row.is_online == 0"
            link
            type="primary"
            size="small"
            @click.prevent="edit(scope.row)"
          >
            修改
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
        page: 1,
        page_size: 20,
      },
    };
  },
  mounted() {
    this.loadData();
  },
  methods: {
    //添加版本
    add() {},
    //发布版本
    publish(data) {},
    //编辑版本
    edit(data) {},
    //保存
    save() {},
    //分页编号
    handleSizeChange(pp) {
      this.queryData.page_size = pp;
      this.queryData.page = 1;
      this.loadData();
    },
    //加载数据
    async loadData(page) {
      if (page) {
        this.queryData.page = page;
      }
      const data = await this.$api("/appver/list", this.queryData);
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
