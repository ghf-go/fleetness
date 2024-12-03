<template>
  <div>
    <el-tag v-if="modelValue > 0" type="primary">是</el-tag>
    <el-tag v-else type="info">否</el-tag>
  </div>
</template>

<script>
export default {
  props: ["modelValue"],
  emits: ["update:modelValue"],
};
</script>

<style lang="scss" scoped></style>
