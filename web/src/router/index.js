import { createRouter, createWebHistory } from "vue-router";
import Dig from "../views/Dig.vue"

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      name: "home",
      component: Dig,
    },
    {
      path:"/dig",
      name:"dig",
      component:Dig
    }
  ],
});

export default router;
