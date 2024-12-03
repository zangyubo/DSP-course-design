<template>
  <div style="height: 100%; width: 100%; background-color: white">
    <v-chart :option="option" style="width: 100%; height: 100%" />
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { exampleLable } from '../data'

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
    data: ['0Hz', '10Hz', '20Hz', '30Hz', '40Hz', '50Hz', '60Hz', '70Hz', '80Hz', '90Hz'], // Frequency bins
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
      data: [10, 20, 35, 50, 70, 90, 110, 150, 180, 200], // Example amplitude data
      barWidth: 1, // Set bar width to a very small value
      itemStyle: {
        color: '#00f', // Color of the bars (lines)
      },
    },
  ],
})
</script>
