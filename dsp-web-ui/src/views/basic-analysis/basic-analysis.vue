<template>
  <div
    style="
      height: 100%;
      width: 100%;
      display: flex;
      flex-direction: row;
      align-items: center;
      justify-content: center;
      background-color: rgb(250, 250, 250);
    "
  >
    <!-- Left section -->
    <div style="width: 49%; height: 100%; display: flex; flex-direction: column">
      <div style="height: 45%; background-color: red; margin-bottom: 10px; margin-top: 10px">
        <FrequencyChart1 />
      </div>
      <div style="height: 45%; background-color: blue">
        <TimeChart1 />
      </div>
      <div
        style="
          justify-content: space-around;
          display: flex;
          flex-direction: row;
          align-items: center;
          margin-top: 10px;
        "
      >
        <a-button @click="triggerFileInput1" :disabled="disableButton1">上传音频文件</a-button>
        <input
          type="file"
          id="audio1"
          @change="handleAudio1Upload"
          accept="audio/*"
          ref="fileInput1"
          style="display: none"
        />
        <a-button type="primary" @click="createChart1" :disabled="disableButton1">
          生成频谱图 1
        </a-button>
        <div style="width: 50%; height: 40px">
          <audio
            v-if="audio1"
            controls
            :src="audio1"
            ref="audioPlayer1"
            style="height: 35px; width: 300px"
          ></audio>
        </div>
      </div>
    </div>

    <!-- Divider -->
    <div style="width: 1%; height: 100%"></div>

    <!-- Right section -->
    <div style="width: 49%; height: 100%; display: flex; flex-direction: column">
      <div style="height: 45%; background-color: red; margin-bottom: 10px; margin-top: 10px">
        <FrequencyChart2 />
      </div>
      <div style="height: 45%; background-color: blue">
        <TimeChart2 />
      </div>
      <div
        style="
          justify-content: space-around;
          display: flex;
          flex-direction: row;
          align-items: center;
          margin-top: 10px;
        "
      >
        <a-button @click="triggerFileInput2" :disabled="disableButton2">上传音频文件</a-button>
        <input
          type="file"
          id="audio2"
          @change="handleAudio2Upload"
          accept="audio/*"
          ref="fileInput2"
          style="display: none"
        />
        <a-button type="primary" @click="upload" :disabled="disableButton2">生成频谱图 2</a-button>
        <div style="width: 50%; height: 40px">
          <audio
            v-if="audio2"
            controls
            :src="audio2"
            ref="audioPlayer2"
            style="height: 35px; width: 300px"
          ></audio>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import FrequencyChart1 from './components/frequency-chart-1.vue'
import FrequencyChart2 from './components/frequency-chart-2.vue'
import TimeChart1 from './components/time-chart-1.vue'
import TimeChart2 from './components/time-chart-2.vue'

import { uploadFileRequest, getChartData } from '../../api/basicCharts'
import { disableButton1, disableButton2 } from './data'
import { audio1, audio2, fileInput1, fileInput2, xAxis1F, yAxis1F, xAxis1T, yAxis1T } from './data'
import { Message } from '@arco-design/web-vue'

const triggerFileInput1 = () => {
  fileInput1.value?.click()
}
const triggerFileInput2 = () => {
  fileInput2.value?.click()
}

// signal 1 upload
const handleAudio1Upload = async (event: Event) => {
  const target = event.target as HTMLInputElement
  const file = target.files?.[0]
  if (file) {
    audio1.value = URL.createObjectURL(file)
  }

  // api to upload file 1
  const uploadFile = fileInput1.value?.files?.[0] // get mp3 1 file
  if (!uploadFile) {
    return
  }

  const test = { file: uploadFile, channel: 1 }

  try {
    // 调用上传接口
    const response = await uploadFileRequest(test)
    console.log('upload file 1 success', response)
  } catch (error) {
    console.error('upload file 1 failed', error)
  }
}

const createChart1 = async () => {
  let response11, response12, response13, response14
  try {
    response11 = await getChartData(1, 1)
    console.log('createChart11', response11.data?.data)
  } catch {}

  try {
    response12 = await getChartData(1, 2)
    console.log('createChart12', response12.data?.data)
  } catch {}

  try {
    response13 = await getChartData(1, 3)
    console.log('createChart13', response13.data?.data)
  } catch {}

  try {
    response14 = await getChartData(1, 4)
    console.log('createChart14', response14.data?.data)
  } catch {}

  for (let i = 0; i < (response11!.data?.data as unknown as number[]).length; i++) {
    console.log(
      (Math.round(response11!.data?.data[i] as number) * 1000) / 1000,
      response12!.data?.data[i],
    )
    xAxis1F.value = []
    yAxis1F.value = []

    xAxis1F.value.push((Math.round(response11!.data?.data[i] as number) * 1000) / 1000)
    yAxis1F.value.push((Math.round(response12!.data?.data[i] as number) * 1000) / 1000)
  }

  // xAxis1F.value = response11 as never
  // yAxis1F.value = response12 as never

  xAxis1T.value = response13 as never
  yAxis1T.value = response14 as never
}

// 处理音频 2 上传
const handleAudio2Upload = (event: Event) => {
  const target = event.target as HTMLInputElement
  const file = target.files?.[0]
  if (file) {
    audio2.value = URL.createObjectURL(file)
  }
}

// ========================================================
// api 上传音乐 1

const upload = async () => {
  const file = fileInput2.value?.files?.[0] // 获取音频 2 的文件
  if (!file) {
    Message.error('请先选择音频文件 1')
    return
  }

  const test = { file: file, channel: 2 }

  try {
    // 调用上传接口
    const response = await uploadFileRequest(test)
    console.log('音频 2 上传成功:', response)
    // 如果需要处理上传后的响应数据，可以在这里继续操作
  } catch (error) {
    console.error('音频 2 上传失败:', error)
  }
}
</script>

<style scoped></style>
