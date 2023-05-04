<template>
  <v-row>
    <v-col>
    </v-col>
    <v-col md="2">
      <v-text-field v-model="month" type="month" class="shrink" density="compact" color="blue">
      </v-text-field>
    </v-col>
  </v-row>

  <Pie :data="data" :options="options" />
</template>
  
<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js'
import { Pie } from 'vue-chartjs'

const month = ref(new Date().toJSON().slice(0, 7))

ChartJS.register(ArcElement, Tooltip, Legend)

const data = reactive({
  labels: ['VueJs', 'EmberJs', 'ReactJs', 'AngularJs'],
  datasets: [
    {
      backgroundColor: ['#41B883', '#E46651', '#00D8FF', '#DD1B16'],
      data: [40, 20, 80, 10]
    }
  ]
})

const options = {
  responsive: true,
  maintainAspectRatio: true,
  aspectRatio: 2
}
</script>