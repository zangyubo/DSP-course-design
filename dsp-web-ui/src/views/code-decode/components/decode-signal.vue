<template>
  <!-- <a-button type="primary" @click="goToHome">Primary Button</a-button> -->

  <div
    style="
      height: 100%;
      width: 100%;
      background-color: rgb(250, 250, 250);
      display: flex;
      align-items: center;
      justify-content: center;
      flex-direction: column;
    "
  >
    <div style="margin: 5px; flex-direction: row; display: flex; height: 50px; width: 600px">
      <a-button type="primary" style="height: 100%; width: 200px">点击上传音乐</a-button>
      <a-input placeholder="输入密钥" style="width: 400px" />
      <a-button type="primary" style="height: 100%; width: 200px">解密</a-button>
    </div>
    <div
      style="
        width: 600px;
        height: 50px;
        margin: 5px;
        display: flex;
        flex-direction: row;
        align-items: center;
      "
    >
      <div style="color: black">原始音频</div>
      <audio controls ref="audioPlayer1" style="height: 100%; width: 90%"></audio>
    </div>
    <div
      style="
        width: 600px;
        height: 50px;
        margin: 5px;
        display: flex;
        flex-direction: row;
        align-items: center;
      "
    >
      <div style="color: black">解密信号</div>
      <audio controls ref="audioPlayer1" style="height: 100%; width: 90%"></audio>
    </div>
    <div>
      <a-button type="primary" style="height: 50px" @click="goToHome">查看频谱图对比</a-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { yAxis1F } from '../../basic-analysis/data'

const router = useRouter()

// Function to navigate to /home
const goToHome = () => {
  yAxis1F.value[1] = 500
  router.push('/basic-analysis')
}
</script>
