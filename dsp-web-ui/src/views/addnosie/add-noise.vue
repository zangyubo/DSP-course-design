<template>
  <div style="height: 100%; width: 100%">
    <div style="height: 50%; width: 100%">
      <dealNoise />
    </div>
    <div style="height: 50%; width: 100%">
      <dealFilliter />
    </div>
  </div>
</template>

<script setup lang="ts">
import dealFilliter from './components/deal-filliter.vue'
import dealNoise from './components/deal-noise.vue'
</script>
