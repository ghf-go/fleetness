<template>
  <div ref="chart"></div>
</template>

<script>
//https://echarts.apache.org/examples/zh/index.html#chart-type-bar
import * as echarts from "echarts";
export default {
  data() {
    return {
      chart: {},
    };
  },
  model: { prop: "value" },
  props: ["value"],
  watch: {
    value(newVal, ov) {
      console.log(ov);
      this.chart.setOption(newVal);
    },
  },

  mounted() {
    this.chart = echarts.init(this.$refs.chart);
    this.chart.setOption(this.value);
  },
};
</script>

<style lang="scss" scoped></style>
