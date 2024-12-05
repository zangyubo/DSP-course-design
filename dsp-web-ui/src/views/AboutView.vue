<template>
  <div style="align-items: center; justify-content: center; width: 100%; height: 100%">
    <button @click="test">test</button>
  </div>
</template>

<script setup lang="ts">
import { sendGetRequest } from '../api/helloWorld'

const test = async () => {
  try {
    // 等待请求完成并接收服务端响应
    const response = await sendGetRequest()

    console.log('Received response:', response)
  } catch (error) {
    console.error('Error:', error) // 捕获并处理错误
  }
}
</script>

<style></style>
