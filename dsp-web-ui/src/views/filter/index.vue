<template>
  <div class="container">
    <div style="height: 70vh; width: 100%">
      <Chart />
    </div>
    <div
      style="
        height: 16vh;
        width: 99%;
        align-items: center;
        justify-content: center;
        display: flex;
      "
    >
      <Function />
    </div>
  </div>
</template>

<script setup lang="ts">
  import Function from './components/function.vue';
  import Chart from './components/chart.vue';
</script>

<style scoped lang="less">
  .container {
    background-color: var(--color-fill-2);
    padding: 8px 10px;
    padding-bottom: 0;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }
</style>
