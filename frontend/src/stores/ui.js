import { defineStore } from "pinia";
import { useStorage } from "@vueuse/core";

export const useUIStore = defineStore("ui", () => {
  // Reactive state persisted to localStorage
  const sidebarCollapsed = useStorage("sidebar-collapsed", false);

  // Actions
  const toggleSidebar = () => {
    sidebarCollapsed.value = !sidebarCollapsed.value;
  };

  const setSidebarCollapsed = collapsed => {
    sidebarCollapsed.value = collapsed;
  };

  const collapseSidebar = () => {
    sidebarCollapsed.value = true;
  };

  const expandSidebar = () => {
    sidebarCollapsed.value = false;
  };

  return {
    // State
    sidebarCollapsed,
    // Actions
    toggleSidebar,
    setSidebarCollapsed,
    collapseSidebar,
    expandSidebar
  };
});
