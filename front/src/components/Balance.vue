<template>
    <v-container>
        <v-row>
            <v-col v-for="user in usr">
                <v-card>
                    <template v-slot:title>
                        <v-avatar color="surface-variant" class="mx-4"></v-avatar>
                        {{ user.name }}
                    </template>

                    <v-card-text>
                        <v-table>
                            <tbody>
                                <tr>
                                    <td>User expenses</td>
                                    <td>{{ user.expenses }}€</td>
                                </tr>
                                <tr>
                                    <td>Total expenses</td>
                                    <td>{{ totalExpenses }}€</td>
                                </tr>
                                <tr>
                                    <td>User ratio</td>
                                    <td>{{ user.ratio }}</td>
                                </tr>
                                <tr>
                                    <td>
                                        Balance
                                        <v-tooltip location="bottom"
                                            text="balance = total expenses * user ratio - user expenses">
                                            <template v-slot:activator="{ props }">
                                                <!-- <v-btn v-bind="props">Tooltip</v-btn> -->
                                                <v-icon icon="mdi-information" v-bind="props" size="x-small"></v-icon>
                                            </template>
                                        </v-tooltip>
                                    </td>
                                    <td>{{ balance(user.expenses, totalExpenses, user.ratio) }}</td>
                                </tr>
                            </tbody>
                        </v-table>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>
        <v-divider class="border-opacity-100 mt-5"></v-divider>
        <v-row class="mt-5">
            <v-col>
                <v-card title="Refunds summary" prepend-icon="mdi-hand-coin">
                    <template v-slot:title>
                        Refunds summary
                    </template>

                    <v-card-text>
                        <v-container>
                            <v-row align="center" justify="space-between">
                                <v-col md="2" align="center">
                                    <v-card color="red" title="user1" class="text-center"></v-card>
                                </v-col>
                                <v-col align="center" md="8">
                                    <v-btn prepend-icon="mdi-minus" append-icon="mdi-arrow-right" size="x-large"
                                        variant="text" style="pointer-events: none">
                                        50€
                                    </v-btn>
                                </v-col>
                                <v-col md="2" align="center">
                                    <v-card color="green" title="user2" class="text-center">
                                    </v-card>
                                </v-col>
                            </v-row>
                        </v-container>
                    </v-card-text>
                </v-card>
            </v-col>

        </v-row>
    </v-container>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue';

const usr = reactive([{ name: "user1", expenses: 200, ratio: 0.2 }, { name: "user2", expenses: 200, ratio: 0.6 }, { name: "user3", expenses: 200, ratio: 0.2 }])
const totalExpenses = ref(usr.reduce((accumulator, currentValue) => accumulator + currentValue.expenses, 0))

function balance(usrExpenses: number, totalExpenses: number, userRatio: number): number {
    return totalExpenses * userRatio - usrExpenses
}
</script>