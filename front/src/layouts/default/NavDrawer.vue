<template>
    <v-navigation-drawer v-model="appStore.showSideBar">
        <template v-slot:prepend>
            <v-list-item lines="two" :title="userStore.getUserName" subtitle="Logged in">
                <template v-slot:prepend>
                    <v-avatar color="grey-darken-3">
                        <span class="text-h6">{{ userStore.getShortUserName }}</span>
                    </v-avatar>
                </template>
            </v-list-item>
        </template>

        <v-divider></v-divider>

        <v-list density="compact" nav v-for="item in navItems">
            <v-list-item @click="router.push({ path: item.path })" :prepend-icon="item.icon" :title="item.title"
                :active="isOnPage(item.path)" :value="item.title"></v-list-item>
        </v-list>
    </v-navigation-drawer>
</template>

<script lang="ts" setup>
import { onBeforeMount } from 'vue';
// import { useUserStore } from '@/store/user';
import { useAppStore } from '@/store/app';
import { useUserStore } from '@/store/user';
import router from "@/router";

const appStore = useAppStore()
const userStore = useUserStore()

// const userStore = useUserStore()

// onBeforeMount(async () => {
//     if (userStore.userName == '') {
//         await userStore.getUsername()
//     }
// })
const navItems = [
    { title: "Expenses", path: "/expenses", icon: "mdi-currency-eur" },
    { title: "Balance", path: "/balance", icon: "mdi-scale-balance" },
    { title: "Stats", path: "/stats", icon: "mdi-chart-line" },
]

function isOnPage(page: string): boolean {
    return router.currentRoute.value.path == page
}
</script>