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
        <frequencyChart1 />
      </div>
      <div style="height: 45%; background-color: blue">
        <timeChart1 />
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
        <a-button type="primary" @click="test" :disabled="disableButton1">生成频谱图 1</a-button>
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
        <frequencyChart2 />
      </div>
      <div style="height: 45%; background-color: blue">
        <timeChart2 />
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
import { disableButton1, disableButton2 } from './data'

// 创建响应式变量来存储音频文件的 URL, 获取文件输入框的引用
import { audio1, audio2, fileInput1, fileInput2 } from './data'

// 触发音频 1 文件选择框
const triggerFileInput1 = () => {
  fileInput1.value?.click()
}

// 触发音频 2 文件选择框
const triggerFileInput2 = () => {
  fileInput2.value?.click()
}

// 处理音频 1 上传
const handleAudio1Upload = (event: Event) => {
  const target = event.target as HTMLInputElement
  const file = target.files?.[0]
  if (file) {
    audio1.value = URL.createObjectURL(file)
  }
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
import { uploadFileRequest } from '../../api/basicCharts'

const upload = async () => {
  const file = fileInput2.value?.files?.[0] // 获取音频 2 的文件
  if (!file) {
    console.error('请先选择音频文件 2')
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
// ========================================================
import { yAxis1F } from './data'

import frequencyChart1 from './components/frequency-chart-1.vue'
import frequencyChart2 from './components/frequency-chart-2.vue'
import timeChart1 from './components/time-chart-1.vue'
import timeChart2 from './components/time-chart-2.vue'

const test = () => {
  yAxis1F.value[1] += 100
}
</script>

<style scoped></style>
