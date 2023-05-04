<template>
  <v-row>
    <v-col>
    </v-col>
    <v-col md="2">
      <v-text-field v-model="month" type="month" class="shrink" density="compact" color="blue">
      </v-text-field>
    </v-col>
  </v-row>

  <Bar :data="data" :options="options" />
</template>
  
<script lang="ts" setup>
import { reactive, computed, ref } from 'vue'
import { useAppStore } from '@/store/app'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale } from 'chart.js'
import { Bar } from 'vue-chartjs'

const appStore = useAppStore()

const month = ref(new Date().toJSON().slice(0, 7))

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend)

const data = reactive({
  labels: [
    'User1',
    'User2'
  ],
  datasets: [
    {
      label: 'Expenses by users',
      backgroundColor: [
        'rgba(255, 99, 132, 0.2)',
        'rgba(54, 162, 235, 0.2)'
      ],
      data: [40, 20]
    }
  ]
})

const gridColor = computed(() => {
  return appStore.getTheme == "dark" ? "rgba(255, 255, 255, 0.2)" : "rgba(0, 0, 0, 0.2)"
})

const options = computed(() => {
  return {
    responsive: true,
    maintainAspectRatio: true,
    aspectRatio: 2,
    legend: {
      display: false
    },
    scales: {
      y: {
        grid: {
          color: gridColor.value
        }
      },
      x: {
        grid: {
          color: gridColor.value
        }
      }
    }
  }
})


</script>
