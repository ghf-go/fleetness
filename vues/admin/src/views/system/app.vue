<template>
  <div>
    <Editor v-model="aaa"></Editor>
    <div>{{ aaa }}</div>

    <div v-html="aaa"></div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      aaa: "1234",
    };
  },
};
</script>

<style lang="scss" scoped></style>
