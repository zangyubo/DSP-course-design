<template>
  <div class="container" style="width: 100%; height: 100%">
    <div style="width: 100%; height: 50%; display: flex">
      <div
        style="
          width: 50%;
          height: 100%;
          background-color: white;
          border: 1px solid black;
        "
      >
        <vueChart :option="freOption" style="height: 100%; width: 100%" />
      </div>

      <div
        style="
          width: 50%;
          height: 100%;
          background-color: white;
          border-right: 1px solid black;
          border-top: 1px solid black;
          border-bottom: 1px solid black;
        "
      >
        <vueChart :option="freOption" style="height: 100%; width: 100%" />
      </div>
    </div>
    <div style="width: 100%; height: 50%; display: flex">
      <div
        style="
          width: 50%;
          height: 100%;
          background-color: white;
          border-right: 1px solid black;
          border-left: 1px solid black;
          border-bottom: 1px solid black;
        "
      >
        <vueChart :option="freOption" style="height: 100%; width: 100%" />
      </div>
      <div
        style="
          width: 50%;
          height: 100%;
          background-color: white;
          border-right: 1px solid black;
          border-bottom: 1px solid black;
        "
      >
        <vueChart :option="freOption" style="height: 100%; width: 100%" /></div
    ></div>
  </div>
</template>

<script setup lang="ts">
  import { ref } from 'vue';
  import { exampleLable, freX, freY } from '../data';

  const freOption = ref({
    graphic: [exampleLable],
    title: {
      text: '频域图',
      left: 'center',
    },
    tooltip: {
      trigger: 'axis',
      formatter: '{b}: {c}',
    },
    xAxis: {
      type: 'category',
      data: freX,
      name: 'F(Hz)',
    },
    yAxis: {
      type: 'value',
      name: 'Amplitude',
      axisLabel: {
        formatter: '{value}',
      },
    },
    series: [
      {
        name: 'Amplitude Spectrum',
        type: 'bar',
        data: freY,
        barWidth: 1,
        itemStyle: {
          color: '#00f',
        },
      },
    ],
  });
</script>

<style scoped lang="less">
  .container {
    background-color: var(--color-fill-2);
    padding: 1px 1px;
    padding-bottom: 0;
    display: flex;
  }
</style>
