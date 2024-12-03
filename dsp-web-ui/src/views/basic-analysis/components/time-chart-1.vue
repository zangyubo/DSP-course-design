<template>
  <div style="height: 100%; width: 100%; background-color: white">
    <v-chart :option="option" style="width: 100%; height: 100%" />
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { exampleLable } from '../data'

// Generate sine wave data for the x-axis values
const xData = Array.from({ length: 100 }, (_, i) => (i + 1) * 0.1) // Generate x-values for the sine wave
const yData = xData.map((x) => Math.sin(x)) // Generate corresponding y-values as the sine of x

const option = ref({
  graphic: [exampleLable],
  title: {
    text: '正弦波图', // "Sine Wave Chart" in Chinese
    left: 'center',
  },
  tooltip: {
    trigger: 'axis',
    formatter: '{b}: {c}', // Shows the x and y values
  },
  xAxis: {
    type: 'category',
    data: xData.map((x) => x.toFixed(1)), // Convert x-values to string for display (rounded to 1 decimal place)
    name: 'x',
  },
  yAxis: {
    type: 'value',
    name: 'Amplitude',
    axisLabel: {
      formatter: '{value}', // Format the amplitude
    },
  },
  series: [
    {
      name: 'Sine Wave',
      type: 'line', // Use line chart type to plot the sine wave
      data: yData, // y-values of the sine wave
      smooth: true, // Make the line smoother for a continuous sine wave effect
      lineStyle: {
        color: '#00f', // Color of the sine wave line
      },
      symbol: 'none', // Remove the circle markers
    },
  ],
})
</script>
