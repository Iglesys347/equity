// Utilities
import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
  state: () => ({
    theme: "dark",
    showSideBar: true
  }),
  getters: {
    getTheme(): string {
      return this.theme
    },
    getShowSideBar(): boolean {
      return this.showSideBar
    }
  },
  actions: {
    switchTheme() {
      this.theme = this.theme == "dark" ? "light" : "dark"
    },
    switchSideBar() {
      this.showSideBar = this.showSideBar ? false : true
    }
  }
})
