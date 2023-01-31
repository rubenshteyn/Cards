import { createWebHistory, createRouter } from 'vue-router'
import startPage from "@/components/startPage"
import aboutPage from "@/components/aboutPage"
import loginPage from "@/components/loginPage"
import forgotPasswordPage from "@/components/forgotPasswordPage"
import registrationPage from "@/components/registrationPage"
import profilePage from "@/components/profilePage"
import verifyEmail from "@/components/verifyEmail"
import resetPassword from "@/components/resetPassword"
import accessPage from "@/components/accessPage"
import tableSessions from "@/components/tableSessions"
import usersCards from "@/components/usersCards"
import tableTransactions from "@/components/tableTransactions"
import userAccounts from "@/components/userAccounts"
import dashBoard from "@/components/dashBoard"
import cardRates from "@/components/cardRates"
import taskList from "@/components/taskList"

const routes = [
    { path: '/', component: startPage },
    { path: '/dashboard', component: dashBoard },
    { path: '/about', component: aboutPage },
    { path: '/login', component: loginPage },
    { path: '/forgot-password', component: forgotPasswordPage },
    { path: '/registration', component: registrationPage },
    { path: '/profile', component: profilePage },
    { path: '/verify-email/:token', component: verifyEmail },
    { path: '/reset-password/:token', component: resetPassword },
    { path: '/access/:token', component: accessPage },
    { path: '/sessions', component: tableSessions},
    { path: '/cards', component: usersCards},
    { path: '/transactions', component: tableTransactions},
    { path: '/accounts', component: userAccounts},
    { path: '/rates', component: cardRates},
    { path: '/tasks', component: taskList},
    { path: '/admin', beforeEnter() {
            window.location.href = process.env.VUE_APP_REDIRECT_URL;
        }}
]
const router = createRouter({
    history: createWebHistory(),
    base: process.env.VUE_APP_URL,
    routes,
})

export default router
