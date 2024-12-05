<template>
  <div class="about">
    <h1>
      <button @click="test">test</button>
    </h1>
  </div>
</template>

<script setup lang="ts">
import { sendPostRequest } from '../api/helloWorld'

const test = async () => {
  try {
    // 等待请求完成并接收服务端响应
    const response = await sendPostRequest({ name: '111' })
    console.log(response.message) // 输出服务端返回的消息
  } catch (error) {
    console.error('Error:', error) // 捕获并处理错误
  }
}
</script>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
