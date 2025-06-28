import { defineStore } from "pinia";
import { computed } from "vue";
import { useStorage } from "@vueuse/core";

export const useEpicsStore = defineStore("epics", () => {
  // Reactive state persisted to localStorage
  const epics = useStorage("jira-epics", []);

  // Actions
  const addEpic = epic => {
    const newEpic = {
      id: Date.now().toString(),
      title: epic.title,
      description: epic.description,
      epicCode: epic.epicCode,
      createdAt: new Date().toISOString()
    };
    epics.value.push(newEpic);
  };

  const updateEpic = (id, updatedEpic) => {
    const index = epics.value.findIndex(epic => epic.id === id);
    if (index !== -1) {
      epics.value[index] = {
        ...epics.value[index],
        ...updatedEpic,
        updatedAt: new Date().toISOString()
      };
    }
  };

  const removeEpic = id => {
    const index = epics.value.findIndex(epic => epic.id === id);
    if (index !== -1) {
      epics.value.splice(index, 1);
    }
  };

  const getEpicById = id => {
    return epics.value.find(epic => epic.id === id);
  };

  const getEpicByCode = code => {
    return epics.value.find(epic => epic.epicCode === code);
  };

  const clearAllEpics = () => {
    epics.value = [];
  };

  // Getters
  const epicCount = computed(() => epics.value.length);

  return {
    // State
    epics,
    // Getters
    epicCount,
    // Actions
    addEpic,
    updateEpic,
    removeEpic,
    getEpicById,
    getEpicByCode,
    clearAllEpics
  };
});
