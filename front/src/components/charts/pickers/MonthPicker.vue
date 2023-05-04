<template>
  <v-row>
    <v-col>
    </v-col>
    <v-col md="2">
      <v-text-field v-model="month" type="month" class="shrink" density="compact" color="blue">
      </v-text-field>
    </v-col>
  </v-row>
</template>

<script lang="ts" setup>
import { ref } from 'vue';

const month = ref(new Date().toJSON().slice(0, 7))
</script>