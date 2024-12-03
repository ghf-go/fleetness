<template>
  <div>
    <UploadImg v-model="img" style="width: 100px; height: 100px"></UploadImg>
    <Editor v-model="data"></Editor>
    ddd
  </div>
</template>

<script>
export default {
  data() {
    return {
      img: "",
      data: "<p>123</p>",
    };
  },
};
</script>

<style lang="scss" scoped></style>
