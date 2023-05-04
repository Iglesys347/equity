<template>
  <v-row>
    <v-col>
    </v-col>
    <v-col md="2">
      <v-text-field v-model="month" type="month" class="shrink" density="compact" color="blue">
      </v-text-field>
    </v-col>
  </v-row>

  <Line :data="data" :options="options" />
</template>
  
<script lang="ts" setup>
import { reactive, ref } from 'vue'
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler
} from 'chart.js'
import { Line } from 'vue-chartjs'

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler
)

const month = ref(new Date().toJSON().slice(0, 7))

const data = reactive({
  labels: ['01', '05', '10', '15', '20', '25', '30'],
  datasets: [
    {
      label: 'Cumulative expenses',
      // backgroundColor: '#f87979',
      data: [10, 15, 50, 55, 60, 70, 120],
      fill: true,
      // backgroundColor: "rgba(66,165,245)",
      backgroundColor: "rgba(66,165,245,.2)",
      borderColor: "rgba(66,165,245)",
      borderWidth: 3,
      lineTension: 0.2,
    },
    {
      label: 'Cumulative Incomes',
      // backgroundColor: '#f87979',
      data: [200, 200, 200, 200, 200, 200, 200],
      fill: false,
      pointRadius: 0,
      // backgroundColor: "rgba(66,165,245)",
      backgroundColor: "rgba(66,245,165,.2)",
      borderColor: "rgba(66,245,165)",
      borderWidth: 3,
      lineTension: 0.2,
    }
  ]
})

const options = {
  responsive: true,
  maintainAspectRatio: true,
  aspectRatio: 2
}

</script>
  