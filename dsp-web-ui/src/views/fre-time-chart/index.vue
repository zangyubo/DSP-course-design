<template>
  <div class="container">
    <div style="height: 80vh; width: 100%">
      <div
        class="fre-time-chart-full-container"
        :style="{ display: showChart ? 'flex' : 'none' }"
      >
        <div class="fre-time-chart-container" style="margin-bottom: 5px">
          <vueChart :option="freOption" style="height: 100%; width: 100%" />
        </div>
        <div class="fre-time-chart-container">
          <vueChart :option="timeOption" style="height: 100%; width: 100%" />
        </div>
      </div>
      <div
        class="fre-time-chart-full-container"
        :style="{ display: showChart ? 'none' : 'flex' }"
      >
        <img
          src="./dist/image.png"
          alt="当前未生成语谱图"
          style="height: 90%; width: 90%"
        />
      </div>
      <div class="fre-time-chart-function-part">
        <a-button
          type="primary"
          class="fre-time-chart-button"
          @click="triggerFileInput"
        >
          上传音频
        </a-button>
        <input
          id="audio"
          ref="fileInput"
          type="file"
          accept="audio/*"
          style="display: none"
          @change="handleAudioUpload"
        />

        <a-button
          type="primary"
          class="fre-time-chart-button"
          @click="getFreAndTimeChart"
        >
          分析信号
        </a-button>

        <audio style="height: 30px; width: 500px" controls :src="audio"></audio>

        <a-button
          type="primary"
          class="fre-time-chart-button"
          @click="hideChart"
        >
          {{ showChartMessage }}
        </a-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
  import { ref } from 'vue';
  import exampleLable from '@/components/chart-example-label/example-label';
  import { freX, freY, timeX, timeY } from './data';

  const exampleLableUse = exampleLable;

  const audio = ref<string | undefined>(undefined);
  const fileInput = ref<HTMLInputElement | null>(null);

  const handleAudioUpload = (event: Event) => {
    const target = event.target as HTMLInputElement;
    const file = target.files?.[0];
    if (file) {
      audio.value = URL.createObjectURL(file);
    }
  };

  const triggerFileInput = () => {
    fileInput.value?.click();
  };

  const showChart = ref(true);
  const showChartMessage = ref('查看语谱图');

  const hideChart = () => {
    showChart.value = !showChart.value;
    if (showChart.value) {
      showChartMessage.value = '查看语谱图';
    } else {
      showChartMessage.value = '查看时频域图';
    }
  };

  timeX.value = Array.from({ length: 100 }, (_, i) =>
    ((i + 1) * 0.1).toFixed(3)
  ) as never[];
  timeY.value = timeX.value.map((x) => Math.sin(x).toFixed(3)) as never[];

  const freOption = ref({
    graphic: exampleLableUse,
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
  const timeOption = ref({
    graphic: exampleLableUse,
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
      data: timeX,
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
        type: 'line',
        data: timeY,
        smooth: true,
        itemStyle: {
          color: '#00f',
        },
        symbol: 'none',
      },
    ],
  });

  const getFreAndTimeChart = async () => {
    freOption.value = {
      graphic: [],
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
    };

    freOption.value.xAxis.data[0] += 10;
  };
</script>

<style scoped lang="less">
  .container {
    background-color: var(--color-fill-2);
    padding: 10px 20px;
    padding-bottom: 0;
    display: flex;
  }
  .fre-time-chart-function-part {
    height: 40px;
    align-items: center;
    justify-content: center;
    display: flex;
    margin-top: 5px;
  }
  .fre-time-chart-button {
    height: 30px;
    width: 100px;
    margin-left: 10px;
    margin-right: 10px;
  }
  .fre-time-chart-full-container {
    height: 100%;
    align-items: center;
    justify-content: center;
    display: flex;
    flex-direction: column;
  }
  .fre-time-chart-container {
    height: 50%;
    width: 90%;
    background-color: white;
  }
</style>
