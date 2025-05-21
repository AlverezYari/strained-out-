<script setup lang="ts">
import { onMounted } from 'vue'
import { useGameStore } from './store/game'

const store = useGameStore()

onMounted(() => {
  store.fetchState()
})
</script>

<template>
  <router-view />
</template>
