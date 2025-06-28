import { createRouter, createWebHashHistory } from "vue-router";
import { useEpicsStore } from "@/stores/epics";

// Views
import HomePage from "@/views/HomePage.vue";
import AdminEpics from "@/views/AdminEpics.vue";
import AdminSetup from "@/views/AdminSetup.vue";
import EpicPage from "@/views/EpicPage.vue";
import AboutPage from "@/views/AboutPage.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: HomePage
  },
  {
    path: "/admin/epics",
    name: "AdminEpics",
    component: AdminEpics
  },
  {
    path: "/admin/setup",
    name: "AdminSetup",
    component: AdminSetup
  },
  {
    path: "/epic/:epicCode",
    name: "Epic",
    component: EpicPage,
    props: true,
    beforeEnter: (to, from, next) => {
      // Check if epic exists in store
      const epicsStore = useEpicsStore();
      const epic = epicsStore.getEpicByCode(to.params.epicCode);

      if (epic) {
        next();
      } else {
        // Redirect to admin epics if epic not found
        next({ name: "AdminEpics" });
      }
    }
  },
  {
    path: "/about",
    name: "About",
    component: AboutPage
  },
  {
    path: "/:pathMatch(.*)*",
    redirect: "/"
  }
];

const router = createRouter({
  history: createWebHashHistory(),
  routes
});

export default router;
