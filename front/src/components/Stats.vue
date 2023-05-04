<template>
  <v-card>
    <v-tabs v-model="tab" bg-color="transparent" grow>
      <v-tab value="total_exp" prepend-icon="mdi-chart-bar-stacked">Total expenses</v-tab>
      <v-tab value="exp_vs_inc" prepend-icon="mdi-chart-donut">Expenses VS Incomes</v-tab>
      <v-tab value="exp_cat" prepend-icon="mdi-chart-pie">Expenses Categories</v-tab>
      <v-tab value="usr_exp" prepend-icon="mdi-chart-bar">Users Expenses</v-tab>
      <v-tab value="cum_exp" prepend-icon="mdi-chart-areaspline-variant">Cumulative Expenses</v-tab>
    </v-tabs>

    <v-card-text>
      <v-window v-model="tab">
        <v-window-item value="total_exp">
          <TotalExpenses />
        </v-window-item>

        <v-window-item value="exp_vs_inc">
          <ExpensesIncomes />
        </v-window-item>

        <v-window-item value="exp_cat">
          <ExpensesCategories />
        </v-window-item>

        <v-window-item value="usr_exp">
          <UsersExpenses />
        </v-window-item>

        <v-window-item value="cum_exp">
          <CumulativeExpenses />
        </v-window-item>
      </v-window>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import CumulativeExpenses from './charts/CumulativeExpenses.vue';
import ExpensesCategories from './charts/ExpensesCategories.vue';
import ExpensesIncomes from './charts/ExpensesIncomes.vue';
import TotalExpenses from './charts/TotalExpenses.vue';
import UsersExpenses from './charts/UsersExpenses.vue';

const tab = ref(null)
</script>