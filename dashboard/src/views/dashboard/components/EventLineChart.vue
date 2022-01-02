<template>
  <div :ref="refName" :class="className" :style="{height:height,width:width}" />
</template>

<script>
import * as echarts from 'echarts'
require('echarts/theme/macarons') // echarts theme
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
      default: '450px'
    },
    autoResize: {
      type: Boolean,
      default: true
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
      default: function() {
        return {}
      }
    }
  },
  data() {
    return {
      chart: null,
      selectData: [130, 140, 141, 142, 145, 150, 160, 11, 87, 20],
      unionData: [120, 82, 91, 154, 162, 140, 130, 140, 141, 142],
      legend: [],
      series: [],
      timeData: []
    }
  },
  watch: {
    data: {
      handler(newValue, oldValue) {
        console.log('newvalue:', newValue)
        const defaultData = [1, 2, 3, 4, 5, 7, 4, 3, 2, 0, 11]
        this.data = newValue
        this.selectData = defaultData
        this.unionData = defaultData

        const intervalTime = (this.data.end - this.data.start) / 10
        const timesSplice = []
        for (var i = 0; i < 10; i++) {
          const timeData = this.data.start + i * intervalTime
          timesSplice.push((new Date(timeData)).toLocaleString())
        }
        timesSplice.push((new Date(this.data.end)).toLocaleString())
        this.timeData = timesSplice
        this.legend = []
        this.series = []
        for (var item in this.data.items) {
          console.log(item)
          this.legend.push(item)
          const tmp = {
            name: item,
            smooth: true,
            type: 'line',
            data: this.data.items[item].interval
          }
          this.series.push(tmp)
        }
        this.initChart()
      },
      deep: true
    }
  },
  mounted() {
    this.$nextTick(() => {
      this.initChart()
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
    initChart() {
      this.chart = echarts.init(this.$refs[this.refName], 'macarons')
      this.setOptions(this.chartData)
    },
    setOptions({ expectedData, actualData } = {}) {
      this.chart.setOption({
        title: {
          text: this.title,
          left: 'left'
        },

        xAxis: {
          name: '日期',
          data: this.timeData,
          boundaryGap: false,
          axisTick: {
            show: false
          },
          axisLabel: {
            interval: 0,
            rotate: -30
          }
        },
        grid: {
          left: 10,
          right: 10,
          bottom: 20,
          top: 30,
          containLabel: true
        },
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'cross'
          },
          padding: [5, 10]
        },
        yAxis: {
          axisTick: {
            show: false
          }
        },
        legend: {
          data: this.legend
        },
        series: this.series
      })
    }
  }
}
</script>
