<template>
  <div :ref="refName" :class="className" :style="{height: height, width: width}" />
</template>

<script>
import * as echarts from 'echarts'
require('echarts/theme/macarons')
import resize from './mixins/resize'

export default {
  mixins: [resize],
  props: {
    className: {
      type: String,
      default: 'chart'
    },
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '400px'
    },
    refName: {
      type: String,
      default: 'main'
    },
    title: {
      type: String,
      default: '事件分布'
    },
    data: {
      type: Object,
      default: function() { return {} }
    }
  },
  data() {
    return {
      chart: null,
      legend: [],
      series: []
    }
  },
  watch: {
    data: {
      handler(newValue, oldValue) {
        this.legend = []
        this.series = []
        for (var item in this.data.items) {
          this.legend.push(item)
          const tmp = {
            name: item,
            value: this.data.items[item].count
          }
          this.series.push(tmp)
        }
        this.initPieChart()
      },
      deep: true
    }
  },
  mounted() {
    this.$nextTick(() => {
      this.initPieChart()
    })
  },
  beforeDestroy() {
    if (!this.chart) {
      return
    }
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    async initPieChart() {
      this.chart = echarts.init(this.$refs[this.refName], 'macarons')
      this.chart.setOption({
        title: {
          text: this.title,
          left: 'center'
        },
        tooltip: {
          trigger: 'item',
          formatter: '{a} <br/>{b} : {c} ({d}%)'
        },
        legend: {
          orient: 'vertical',
          left: 'left',
          data: this.legend
        },
        series: [
          {
            name: '统计内容',
            type: 'pie',
            radius: '55 %',
            center: ['50%', '50%'],
            data: this.series,
            emphasis: {
              itemStyle: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: 'rgba(0, 0, 0, 0.5)'
              }
            }
          }
        ]
      })
    }
  }
}
</script>
