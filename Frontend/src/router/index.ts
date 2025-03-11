import {createRouter, createWebHistory} from 'vue-router'
import Login from '../views/Login.vue'
import HomeFeed from '../views/HomeFeed.vue'
import User from "../views/user.vue";
import PaymentManagement from "../views/PaymentManagement.vue";
import MealManagement from "../views/MealManagement.vue";
import DebtOverview from "../views/DebtOverview.vue";
import Dashboard from "../views/Dashboard.vue";
import Profile from "../views/Profile.vue";

const routes = [{
    path: '/',
    redirect: '/login' // Leere URL wird zu /login umgeleitet
},
    {
        path: '/login',
        name: 'login',
        component: Login
    },
    {
        path: '/homefeed',
        name: 'HomeFeed',
        component: HomeFeed
    },
    {
        path: '/user',
        name: 'User',
        component: User
    },
    {
        path: '/payments',
        name: 'PaymentManagement',
        component: PaymentManagement
    },
    {
        path: '/meal',
        name: 'MealManagement',
        component: MealManagement
    },   {
        path: '/debts',
        name: 'DebtOverview',
        component: DebtOverview
    }, {
        path: '/dashboard',
        name: 'Dashboard',
        component: Dashboard
    },{
        path: '/profile',
        name: 'profile',
        component: Profile
    },

]

const router = createRouter({
    history: createWebHistory(), // Optional: process.env.BASE_URL, falls du Umgebungsvariablen nutzt
    routes
})

export default router
