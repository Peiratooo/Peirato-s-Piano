
import {createRouter, createWebHashHistory, createWebHistory} from "vue-router"

const routes = [
    // {
    //     path: "",
    //     name: "sandtable",
    //     component: () => import("../components/Sandtable.vue"),
    // },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

router.beforeEach((to,from,next)=>{
    next()
})

export default router