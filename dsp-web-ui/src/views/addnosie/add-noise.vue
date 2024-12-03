<template>
  <a-button type="primary" @click="goToHome">Primary Button</a-button>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { yAxis1F } from '../basic-analysis/data'

const router = useRouter()

// Function to navigate to /home
const goToHome = () => {
  yAxis1F.value[1] = 500
  router.push('/basic-analysis')
}
</script>
