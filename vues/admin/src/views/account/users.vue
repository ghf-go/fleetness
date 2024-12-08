<template>
  <div>
    <el-button type="primary" size="default" @click="addNewUser"
      >添加账号</el-button
    >
    <el-form
      style="text-align: right"
      :model="queryData"
      :inline="true"
      size="normal"
    >
      <el-form-item style="text-align: right">
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
          <UserInfo v-model="scope.row" />
        </template>
      </el-table-column>
      <el-table-column prop="create_at" label="注册时间" />
      <el-table-column prop="create_ip" label="注册IP" />
      <el-table-column label="操作" fixed="right">
        <template #default="scope">
          <el-button
            link
            type="primary"
            size="small"
            @click.prevent="resetPass(scope.row)"
          >
            重置密码
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

    <el-dialog v-model="addUserData.isShowAddUser" title="添加账号" width="500">
      <el-form :model="form" label-position="top">
        <el-form-item label="登录名称">
          <el-input v-model="addUserData.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="密码">
          <el-input
            v-model="addUserData.pass"
            type="password"
            autocomplete="off"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="addUserData.isShowAddUser = false">取消</el-button>
          <el-button type="primary" @click="createUser"> 添加 </el-button>
        </div>
      </template>
    </el-dialog>
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
    //添加新用户
    addNewUser() {
      this.addUserData = {
        isShowAddUser: true,
        name: "",
        pass: "",
      };
    },
    //前进账号
    async createUser() {
      if (this.addUserData.name == "" || this.addUserData.pass == "") {
        this.$message.error("请输入正确的账号");
        return;
      }
      const res = await this.$api("/account/user_add", this.addUserData);
      if (res.code == 200) {
        this.addUserData.isShowAddUser = false;
        this.$message.success("添加账号成功");
        this.loadData();
      } else {
        this.$message.error(res.msg);
      }
    },
    //修改账号密码
    resetPass(uinfo) {
      this.$prompt("修改密码", "新密码", {
        confirmButtonText: "确认修改密码",
        cancelButtonText: "取消",
        inputPlaceholder: "请输入新密码",
        inputValidator: (v) => {
          if (!v || v.length < 6) {
            return "新密码最小6位字符";
          }
          return true;
        },
      })
        .then(({ value }) => {
          this.$api("/account/user_changepass", {
            uid: uinfo.id,
            pass: value,
          }).then((r) => {
            if (r.code == 200) {
              this.$message.success("修改成功");
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
    handleSizeChange(pp) {
      this.queryData.page_size = pp;
      this.queryData.page = 1;
      this.loadData();
    },

    async loadData(page) {
      if (page) {
        this.queryData.page = page;
      }
      const data = await this.$api("/account/user_list", this.queryData);
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
