<template>
    <v-container>
        <v-row>
            <v-col>
            </v-col>

            <!-- <v-col md="4">
                <v-text-field v-model="filterValue" density="compact" variant="solo" label="Search email or domains"
                    append-inner-icon="mdi-magnify" hide-details
                    @update:model-value="page = 1; getResults()"></v-text-field>
            </v-col> -->
        </v-row>
    </v-container>

    <v-table fixed-header>
        <thead>
            <tr>
                <th class="text-left">
                    Name
                </th>
                <th class="text-left">
                    Date
                </th>
                <th class="text-left">
                    Category
                </th>
                <th class="text-left">
                    Amount
                </th>
                <th class="text-left">
                    User ID
                </th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="item in expenses" :key="item.id">
                <td>{{ item.name }}</td>
                <td>{{ item.date }}</td>
                <td>{{ item.category }}</td>
                <td>{{ item.amount }}</td>
                <td>{{ item.user_id }}</td>
            </tr>
        </tbody>
    </v-table>

</template>

<script lang="ts" setup>
import { ref, onMounted, reactive } from 'vue';
import { getExpenses } from '@/api/expenses';
import { Ref } from 'vue';

interface Expense {
    id: number,
    name: string,
    date: Date,
    category: string,
    amount: number,
    description: string,
    user_id: number
}

// const expenses = ref<Expense[]>()
// const expenses: Ref<Expense[]> = ref([]);
const expenses: Expense[] = reactive([])
// const expenses: Expense[] | Ref<never[]> = ref([]);

function getAllExpenses() {
    // expenses = []
    getExpenses().then((response) => {
        console.log(response)
        for (const elt of response.data) {
            expenses.push(elt)
        }
    }).catch((err) => {
        console.log("Error retrieving expenses")
        console.log(err)
    });
}

onMounted(() => {
    getAllExpenses()
})

</script>