<template>
  <div id="app">
    <h1>上传并播放音频</h1>

    <!-- 上传音频 1 -->
    <div>
      <label for="audio1" class="upload-button">选择音频 1: </label>
      <input
        type="file"
        id="audio1"
        @change="handleAudio1Upload"
        accept="audio/*"
        ref="fileInput1"
        style="display: none"
      />
      <button @click="triggerFileInput1" class="upload-trigger">上传音频 1</button>
      <audio v-if="audio1" controls :src="audio1" ref="audioPlayer1"></audio>
      <button v-if="audio1" @click="playAudio1">播放音频 1</button>
    </div>

    <!-- 上传音频 2 -->
    <div>
      <label for="audio2" class="upload-button">选择音频 2: </label>
      <input
        type="file"
        id="audio2"
        @change="handleAudio2Upload"
        accept="audio/*"
        ref="fileInput2"
        style="display: none"
      />
      <button @click="triggerFileInput2" class="upload-trigger">上传音频 2</button>
      <audio v-if="audio2" controls :src="audio2" ref="audioPlayer2"></audio>
      <button v-if="audio2" @click="playAudio2">播放音频 2</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// 创建响应式变量来存储音频文件的 URL
const audio1 = ref<string | null>(null)
const audio2 = ref<string | null>(null)

// 获取文件输入框的引用
const fileInput1 = ref<HTMLInputElement | null>(null)
const fileInput2 = ref<HTMLInputElement | null>(null)

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

// 播放音频 1
const playAudio1 = () => {
  const audioPlayer1 = document.querySelector('#audioPlayer1') as HTMLAudioElement | null
  if (audioPlayer1) {
    audioPlayer1.play()
  }
}

// 播放音频 2
const playAudio2 = () => {
  const audioPlayer2 = document.querySelector('#audioPlayer2') as HTMLAudioElement | null
  if (audioPlayer2) {
    audioPlayer2.play()
  }
}
</script>

<style>
#app {
  text-align: center;
  margin-top: 20px;
}

.upload-button {
  cursor: pointer;
  margin-right: 10px;
}

.upload-trigger {
  margin-left: 10px;
}

input[type='file'] {
  margin: 10px;
}
</style>
