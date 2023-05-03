<template>
  <v-app :theme="appStore.getTheme">
    <default-bar />
    <default-nav />
    <default-view />
  </v-app>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

import DefaultBar from './AppBar.vue'
import DefaultNav from './NavDrawer.vue'
import DefaultView from './View.vue'

import { useAppStore } from '@/store/app';

const appStore = useAppStore()
</script>
