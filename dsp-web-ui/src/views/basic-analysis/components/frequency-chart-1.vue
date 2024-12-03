<template>
  <div style="height: 100%; width: 100%; background-color: white">
    <v-chart :option="option" style="width: 100%; height: 100%" />
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { exampleLable, xAxis1F, yAxis1F } from '../data'

const option = ref({
  graphic: [exampleLable],
  title: {
    text: '频域图', // "Frequency Spectrum" in Chinese
    left: 'center',
  },
  tooltip: {
    trigger: 'axis',
    formatter: '{b}: {c}', // Shows the frequency value for each bin
  },
  xAxis: {
    type: 'category',
    data: xAxis1F, // Frequency bins
    name: 'F(Hz)',
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
      name: 'Amplitude Spectrum',
      type: 'bar',
      data: yAxis1F, // Example amplitude data
      barWidth: 1, // Set bar width to a very small value
      itemStyle: {
        color: '#00f', // Color of the bars (lines)
      },
    },
  ],
})
</script>
