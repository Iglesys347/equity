// User infos
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    userName: "defaultuser",
  }),
  getters: {
    getUserName(): string {
      return this.userName
    },
    getShortUserName(): string {
      return String(this.userName).slice(0, 2)
    }
  },
  actions: {
    setUserName(userName: string) {
      this.userName = userName
    }
  }
})
