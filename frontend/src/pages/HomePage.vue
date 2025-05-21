<script setup lang="ts">
import { computed } from 'vue'
import { useGameStore } from '../store/game'
import type { Mission } from '../types/game'

const store = useGameStore()
const mission = computed<Mission | undefined>(() => store.state?.mission)
</script>

<template>
  <div>
    <h1>{{ mission?.name || 'Loading...' }}</h1>
    <pre>{{ mission }}</pre>
  </div>
</template>
