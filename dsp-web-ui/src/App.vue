<template>
  <div style="height: 100vh; width: 95vw; background-color: white">
    <div style="height: 30px; background-color: rgb(250, 250, 250)">
      <nav>
        <RouterLink to="/" exact-active-class="active-link">Home</RouterLink>
        <RouterLink to="/basic-analysis">Basic Analysis</RouterLink>
        <RouterLink to="/add-noise">Add Noise</RouterLink>
        <RouterLink to="/code-decode">Code and Decode</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </div>

    <div style="height: calc(100vh - 30px); width: 100%; background-color: rgb(255, 255, 255)">
      <RouterView />
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<style scoped>
nav {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
  font-size: 16px;
}

nav a.router-link-exact-active {
  color: black; /* 设置选中状态的颜色 */
}

nav a.router-link-exact-active:hover {
  background-color: transparent; /* 保持 hover 时不变 */
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

/* 自定义的选中状态颜色 */
.active-link {
  color: black; /* 刷新时设置 Home 链接为红色 */
}
</style>
