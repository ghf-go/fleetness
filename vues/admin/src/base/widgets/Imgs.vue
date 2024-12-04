<template>
  <div class="imgs">
    <el-image
      v-for="(i, url) in imglist"
      :key="i"
      :src="url"
      fit="cover"
      :hide-on-click-modal="true"
      :preview-src-list="imglist"
      :lazy="true"
      :style="{ width: maxw }"
      class="item"
    ></el-image>
  </div>
</template>

<script>
export default {
  props: ["modelValue", "width"],
  emits: ["update:modelValue"],
  computed: {
    imglist() {
      if (this.modelValue instanceof Array) {
        return this.modelValue;
      } else {
        return this.modelValue.split(",");
      }
    },
    maxw() {
      if (this.width) {
        return this.width;
      } else {
        return "100px";
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.imgs {
  display: flex;
  flex-wrap: wrap;
  .item {
    margin: 0.5rem;
  }
}
</style>
