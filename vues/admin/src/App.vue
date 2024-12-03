<template>
  <router-view />
</template>

<style>
* {
  box-sizing: border-box;
}
body {
  margin: 0px;
  padding: 0px;
}
.el-table .warning-row {
  --el-table-tr-bg-color: var(--el-color-warning-light-9);
}
.el-table .success-row {
  --el-table-tr-bg-color: var(--el-color-success-light-9);
}
</style>
