<template>
  <div>
    <el-form :inline="false" size="normal">
      <el-form-item>
        <el-button type="primary" size="default" @click="addFriendLink"
          >添加</el-button
        >
        <el-button type="success" size="default" @click="save">保存</el-button>
      </el-form-item>
    </el-form>

    <el-table
      :data="tableData"
      border
      stripe
      style="width: 100%"
      :row-class-name="$tableRowClassName"
    >
      <el-table-column prop="id" label="id" fixed> </el-table-column>
      <el-table-column label="名称" width="200">
        <template #default="scpoe">
          <el-input
            v-model="scpoe.row.name"
            placeholder="名称"
            size="normal"
          ></el-input>
        </template>
      </el-table-column>
      <el-table-column label="类型">
        <template #default="scpoe">
          <el-select v-model="scpoe.row.link_type" placeholder="类型">
            <el-option label="WEB" value="WEB"> </el-option>
            <el-option label="APP" value="APP"> </el-option>
          </el-select>
        </template>
      </el-table-column>
      <el-table-column label="logo" width="200">
        <template #default="scpoe">
          <UploadImg v-model="scpoe.row.logo"></UploadImg>
        </template>
      </el-table-column>
      <el-table-column label="地址" width="200">
        <template #default="scpoe">
          <el-input
            v-model="scpoe.row.url"
            placeholder="地址"
            size="normal"
          ></el-input>
        </template>
      </el-table-column>
      <el-table-column label="IOS应用商店地址" width="200">
        <template #default="scpoe">
          <el-input
            v-model="scpoe.row.ios"
            placeholder="IOS应用商店地址"
            size="normal"
          ></el-input>
        </template>
      </el-table-column>
      <el-table-column label="GooglePlay地址" width="200">
        <template #default="scpoe">
          <el-input
            v-model="scpoe.row.google_pay"
            placeholder="GooglePlay地址"
            size="normal"
          ></el-input>
        </template>
      </el-table-column>
      <el-table-column label="背景图片" width="200">
        <template #default="scpoe">
          <UploadImg v-model="scpoe.row.bg_img"></UploadImg>
        </template>
      </el-table-column>
      <el-table-column label="描述" width="200">
        <template #default="scpoe">
          <el-input
            v-model="scpoe.row.content"
            placeholder="描述"
            size="normal"
          ></el-input>
        </template>
      </el-table-column>
      <el-table-column label="是否显示">
        <template #default="scpoe">
          <el-switch
            v-model="scpoe.row.is_show"
            :active-value="1"
            :inactive-value="0"
          >
          </el-switch>
        </template>
      </el-table-column>
      <el-table-column label="排序">
        <template #default="scpoe">
          <el-input
            v-model="scpoe.row.sort_index"
            placeholder="排序"
            size="normal"
          ></el-input>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      tableData: [],
    };
  },
  mounted() {
    this.loadData();
  },
  methods: {
    //添加
    addFriendLink() {
      this.tableData.push({
        id: 0,
        name: "",
        link_type: "WEB",
        logo: "",
        url: "",
        ios: "",
        google_pay: "",
        bg_img: "",
        content: "",
        is_show: 1,
        sort_index: 0,
      });
    },
    async save() {
      const res = await this.$api("/friendlink/save", this.tableData);
      if (res.code != 200) {
        this.$message.error(res.msg);
        return;
      }
      this.$message.success("修改成功");
      this.loadData();
    },
    //加载数据
    async loadData() {
      const res = await this.$api("/friendlink/list", {});
      if (res.code != 200) {
        this.$message.error(res.msg);
        return;
      }
      this.tableData = res.data;
    },
  },
};
</script>

<style lang="scss" scoped></style>
