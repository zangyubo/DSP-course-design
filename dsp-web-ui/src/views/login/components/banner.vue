<template>
  <div class="banner">
    <div class="banner-inner">
      <div
        ref="chartContainer"
        style="width: 100%; height: 100%; background-color: transparent"
      >
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { ref, onMounted, onBeforeUnmount } from 'vue';
  import * as echarts from 'echarts';

  import { use } from 'echarts/core';
  import { CanvasRenderer } from 'echarts/renderers';
  // @ts-ignore
  import { Line3DChart } from 'echarts-gl/charts';
  // @ts-ignore
  import { Grid3DComponent } from 'echarts-gl/components';

  // if
  // {// @ts-ignore
  // import { Line3DChart } from 'echarts-gl/charts';
  // // @ts-ignore
  // import { Grid3DComponent } from 'echarts-gl/components';
  // }
  // this part cound not run please use the following code

  // import 'echarts-gl';

  // but if you use this line of code, you have to bear a warn of "geo3D exists.",though the warn don't cause any problem in running

  use([Grid3DComponent, Line3DChart, CanvasRenderer]);

  const chartContainer = ref<HTMLDivElement | null>(null);
  let chart: echarts.ECharts | null = null;

  const signal1 = (x: number): number => {
    return Math.sin(5 * Math.PI * x);
  };

  const signal2 = (x: number): number => {
    return Math.sin(8 * Math.PI * x);
  };

  const complexSignal = (x: number): number => {
    return signal1(x) + signal2(x);
  };

  const getSignal = (mode: number): [number, string, number][] => {
    // mode == 1 => signal1;
    // mode == 2 => signal2;
    // mode == 3 => complexSignal;

    const data: [number, string, number][] = [];

    for (let i = -3; i <= 3; i += 0.001) {
      if (mode === 1) {
        data.push([i, 'signal1', signal1(i)]);
      } else if (mode === 2) {
        data.push([i, 'signal2', signal2(i)]);
      } else if (mode === 3) {
        data.push([i, 'complex signal', complexSignal(i)]);
      }
    }

    return data;
  };

  const initChart = (): void => {
    if (!chartContainer.value) return;

    chart = echarts.init(chartContainer.value);

    const option: echarts.EChartsOption = {
      grid3D: {
        boxWidth: 3000,
        boxHeight: 200,
        boxDepth: 300,
        axisLine: {
          show: true,
        },
        viewControl: {
          distance: 800,
          alpha: 20,
          beta: 60,
        },
      },
      xAxis3D: {
        type: 'category',
        name: 'X',
        nameTextStyle: {
          color: '#fff',
        },
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
        nameTextStyle: {
          color: '#fff',
        },
        axisLabel: {
          show: true,
          color: '#fff',
        },
      },
      zAxis3D: {
        type: 'value',
      },
      series: [
        {
          type: 'line3D',
          name: 'complex signal',
          data: getSignal(3),
        },
        {
          type: 'line3D',
          name: 'signal1',
          data: getSignal(1),
        },
        {
          type: 'line3D',
          name: 'signal2',
          data: getSignal(2),
        },
      ] as never,
    };

    chart.setOption(option);
  };

  onMounted(() => {
    initChart();
  });

  onBeforeUnmount(() => {
    if (chart) {
      chart.dispose();
    }
  });
</script>

<style lang="less" scoped>
  .banner {
    display: flex;
    align-items: center;
    justify-content: center;
    overflow: hidden;

    &-inner {
      flex: 1;
      height: 100%;
    }
  }
</style>
