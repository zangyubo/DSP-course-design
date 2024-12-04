<template>
  <div style="height: 100%; width: 100%">
    <div style="height: 50%; width: 100%">
      <codeSignal />
    </div>
    <div style="height: 50%; width: 100%">
      <decodeSignal />
    </div>
  </div>
</template>

<script setup lang="ts">
import codeSignal from './components/code-signal.vue'
import decodeSignal from './components/decode-signal.vue'
</script>
