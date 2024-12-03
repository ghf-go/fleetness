<template>
  <div>
    <!-- <UploadImg v-model="img" style="width: 100px; height: 100px"></UploadImg> -->
    <UploadImgs v-model="imgs" />
    <div>{{ imgs }}</div>
    <!-- <Editor v-model="data"></Editor> -->

    ddd
  </div>
</template>

<script>
export default {
  data() {
    return {
      img: "",
      imgs: "",
      data: "<p>123</p>",
    };
  },
};
</script>

<style lang="scss" scoped></style>
