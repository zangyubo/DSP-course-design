<template>
  <div ref="chartContainer" style="width: 100%; height: 100%"></div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'
import * as echarts from 'echarts'
import 'echarts-gl'

// 定义容器引用
const chartContainer = ref<HTMLDivElement | null>(null)
let chart: echarts.ECharts | null = null

// 定义函数
const func = (x: number): number => {
  x /= 10
  return Math.sin(x) * Math.cos(x * 2 + 1) * Math.sin(x * 3 + 2) * 50
}

const func2 = (x: number): number => {
  x /= 10
  return Math.sin(x) * 50
}

const generateData = (funtype: number): [number, string, number][] => {
  const data: [number, string, number][] = []
  for (let i = -200; i <= 200; i += 0.1) {
    data.push([i, funtype === 1 ? 'test1' : 'test2', funtype === 1 ? func(i) : func2(i)])
  }
  return data
}

// 初始化图表
const initChart = (): void => {
  if (!chartContainer.value) return

  chart = echarts.init(chartContainer.value)

  const option: echarts.EChartsOption = {
    grid3D: {
      boxWidth: 300,
      boxHeight: 120,
      boxDepth: 200,
      axisLine: {
        show: true,
      },
      viewControl: {
        distance: 400,
      },
    },
    xAxis3D: {
      type: 'category',
      name: 'X',
      axisLabel: {
        show: false,
      },
      splitLine: {
        show: true,
      },
    },
    yAxis3D: {
      type: 'category',
      name: 'Y',
      axisLabel: {
        show: true,
      },
    },
    zAxis3D: {
      type: 'value',
    },
    series: [
      {
        type: 'line3D',
        name: 'test1',
        data: generateData(1),
      },
      {
        type: 'line3D',
        name: 'test2',
        data: generateData(2),
      },
    ] as never,
  }

  chart.setOption(option)
}

// 生命周期钩子
onMounted(() => {
  initChart()
})

onBeforeUnmount(() => {
  if (chart) {
    chart.dispose()
  }
})
</script>

<style scoped>
/* 让图表容器有合适的大小 */
</style>
