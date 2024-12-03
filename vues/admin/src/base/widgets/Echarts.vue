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
  props: ["modelValue"],
  emits: ["update:modelValue"],
  watch: {
    modelValue(newVal, ov) {
      console.log(ov);
      this.chart.setOption(newVal);
    },
  },

  mounted() {
    this.chart = echarts.init(this.$refs.chart);
    this.chart.setOption(this.modelValue);
  },
};
</script>

<style lang="scss" scoped></style>
