<template>
    <v-row>
        <v-col>
        </v-col>
        <v-col md="2">
            <v-text-field v-model="month" type="month" class="shrink" density="compact" color="blue">
            </v-text-field>
        </v-col>
    </v-row>

    <Doughnut :data="data" :options="options" />
</template>
  
<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js'
import { Doughnut } from 'vue-chartjs'

const month = ref(new Date().toJSON().slice(0, 7))

// Chart config
ChartJS.register(ArcElement, Tooltip, Legend)
const data = reactive({
    labels: ['Incomes', 'Expenses'],
    datasets: [
        {
            backgroundColor: ['#41B883', '#E46651',],
            data: [80, 20]
        }
    ]
})
const options = {
    responsive: true,
    maintainAspectRatio: true,
    aspectRatio: 2
}
</script>


  