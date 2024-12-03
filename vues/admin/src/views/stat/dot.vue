<template>
  <div>
    <Echarts v-model="statData" style="width: 100%; height: 600px"></Echarts>
  </div>
</template>

<script>
export default {
  data() {
    return {
      statData: {},
      queryData: {
        start: "1970-01-01",
        keys: {
          dddd: 1,
        },
      },
    };
  },
  mounted() {
    this.loadData();
  },
  methods: {
    async loadData() {
      const data = await this.$api("/metrics/stat", this.queryData);
      if (data.code != 200) {
        this.$message.error(data.msg);
      } else {
        this.statData = data.data;
      }
    },
  },
};
</script>

<style lang="scss" scoped></style>
