import { defineStore } from 'pinia'
import axios from 'axios'
import type { State } from '../types/game'

export const useGameStore = defineStore('game', {
  state: () => ({ state: null as State | null }),
  actions: {
    async fetchState() {
      const res = await axios.get<State>('/api/state')
      this.state = res.data
    }
  }
})
