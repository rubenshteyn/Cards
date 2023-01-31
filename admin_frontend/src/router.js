import { createWebHistory, createRouter } from 'vue-router'
import loginPage from "@/components/loginPage"
import mainPage from "@/components/mainPage"
import usersCards from "@/components/usersCards"
import transactionLogModal from "@/components/modals/transactionLogModal"
import tableSessions from "@/components/tableSessions"
import tableSuppliers from "@/components/tableSuppliers"
import tableTransactions from "@/components/tableTransactions"
import tableUsers from "@/components/tableUsers"
import tableLogs from "@/components/tableLogs"
import taskList from "@/components/taskList"

const routes = [
    { path: '/', component: mainPage },
    { path: '/users', component: tableUsers},
    { path: '/login', component: loginPage },
    { path: '/cards', component: usersCards},
    { path: '/log', component: transactionLogModal},
    { path: '/sessions', component: tableSessions},
    { path: '/suppliers', component: tableSuppliers},
    { path: '/transactions', component: tableTransactions},
    { path: '/logs', component: tableLogs},
    { path: '/tasks', component: taskList},
    { path: '/access/:token', beforeEnter() {
            window.location.href = process.env.VUE_APP_URL + "/access";
        }}
]
const router = createRouter({
    history: createWebHistory(),
    base: process.env.VUE_APP_URL,
    routes,
})

export default router
