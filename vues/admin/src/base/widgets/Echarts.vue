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
      this.setOption(newVal);
    },
  },

  mounted() {
    this.chart = echarts.init(this.$refs.chart);
    this.setOption(this.modelValue);
  },
  methods: {
    setOption(val) {
      if (val.tooltip.filter0) {
        val.tooltip.formatter = function (params) {
          var res = `${params[0].name} <br/>`;
          for (const item of params) {
            if (item.value !== 0) {
              res += `<span style="background: ${item.color}; height:10px; width: 10px; border-radius: 50%;display: inline-block;margin-right:10px;"></span> ${item.seriesName} ：${item.value}<br/>`;
            }
          }
          return res;
        };
      }
      this.chart.setOption(val);
    },
  },
};
</script>

<style lang="scss" scoped></style>
