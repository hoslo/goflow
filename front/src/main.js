import { createApp } from 'vue'
import ElementPlus from 'element-plus';
import 'element-plus/lib/theme-chalk/index.css';
import { createRouter, createWebHashHistory } from "vue-router";

import axios from 'axios'
import VueAxios from 'vue-axios'

const routes = [
    { path: '/', component: Login, name: "login"},
    { path: '/dag_list', component: DashboardDagList, name: "dag_list"},
    { path: '/httpoperator_list', component: DashboardHttpOperatorList, name: "httpoperator_list"},
    { path: '/dag/:id', component: DashboardDag, name: "dag"},
    { path: '/task/:id', component: DashboardTask, name: "task"},
]
const router = createRouter({
    // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
    history: createWebHashHistory(),
    routes, // `routes: routes` 的缩写
})

import App from './App.vue'
import Login from "@/page/Login";
import DashboardDagList from "@/page/DashboardDagList";
import DashboardDag from "@/page/DashboardDag";
import DashboardHttpOperatorList from "@/page/DashboardHttpOperatorList";
import DashboardTask from "@/page/DashboardTask";

const app = createApp(App)
app.use(ElementPlus)
app.use(VueAxios, axios)
app.use(router)
app.mount('#app')
