<!-- TODO Add v-if="loaded" to all the chart and set loaded to true only when data are fetched from API -->

<template>
  <v-container>
    <v-row justify="end">
      <v-col>
        <v-text-field label="Starting month" type="month" density="compact"></v-text-field>
      </v-col>
      <v-icon class="mt-6">mdi-arrow-right-bold</v-icon>
      <v-col>
        <v-text-field label="Ending month" type="month" density="compact"></v-text-field>
      </v-col>
    </v-row>
    <v-col class="text-right">
      <v-btn-toggle v-model="chart" variant="outlined" density="compact">
        <v-btn value="bar">
          <v-icon>mdi-chart-bar</v-icon>
        </v-btn>

        <v-btn value="line">
          <v-icon>mdi-chart-line</v-icon>
        </v-btn>

      </v-btn-toggle>
    </v-col>
  </v-container>


  <Line v-if="chart == 'line'" :data="data" :options="options" />
  <Bar v-else :data="data" :options="options" />
</template>
    
<script lang="ts" setup>
import { computed, ref } from 'vue'
import { useAppStore } from '@/store/app'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, LineElement, CategoryScale, LinearScale, PointElement, } from 'chart.js'
import { Bar, Line } from 'vue-chartjs'

const appStore = useAppStore()

ChartJS.register(CategoryScale, LinearScale, BarElement, LineElement, PointElement, Title, Tooltip, Legend)

const data = ({
  labels: ['January', 'February', 'March', 'April', 'May'],
  datasets: [
    {
      label: 'Total expenses',
      backgroundColor: 'rgba(66,165,245,.5)',
      data: [40, 20, 50, 80, 10],
      // backgroundColor: "rgba(66,165,245,.2)",
      borderColor: "rgba(66,165,245)",
      borderWidth: 3,
      lineTension: 0.3,
    }
  ]

})

const chart = ref("bar")

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
  