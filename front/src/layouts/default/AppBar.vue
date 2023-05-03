<template>
  <v-app-bar dense elevation="1">
    <v-app-bar-title>
      <v-btn icon="mdi-menu" @click="appStore.switchSideBar()"></v-btn>
      Equity
    </v-app-bar-title>
    <!-- <v-img contain height="500" src="@/assets/narwal-nobg.png" /> -->
    <v-spacer></v-spacer>
    <v-btn :icon="appStore.getTheme === 'light' ? 'mdi-weather-sunny' : 'mdi-weather-night'"
      @click="appStore.switchTheme()"></v-btn>
  </v-app-bar>
</template>

<script lang="ts" setup>
import { useAppStore } from '@/store/app';

const appStore = useAppStore()
</script>
