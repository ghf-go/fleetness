<template>
  <div class="user_info">
    <el-avatar :src="modelValue.avatar" fit="cover"></el-avatar>
    <div class="name">{{ modelValue.nick_name }}</div>
  </div>
</template>

<script>
export default {
  props: ["modelValue"],
  emits: ["update:modelValue"],
};
</script>

<style lang="scss" scoped>
.user_info {
  display: block;
  text-align: center;
  .avatar {
    width: 80%;
  }
  .name {
    margin: 0.5rem;
  }
}
</style>
