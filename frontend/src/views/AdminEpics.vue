<template>
  <div class="admin-epics-page">
    <div class="container">
      <!-- Header -->
      <div class="page-header">
        <h1>Manage Epics</h1>
        <p>Add, edit, and remove epics you want to track and analyze.</p>
      </div>

      <!-- Add/Edit Epic Form -->
      <div class="epic-form-section">
        <div class="card">
          <h2>{{ isEditing ? 'Edit Epic' : 'Add New Epic' }}</h2>

          <form @submit.prevent="saveEpic" class="epic-form">
            <div class="form-group">
              <label for="title">Epic Title *</label>
              <input
                id="title"
                v-model="form.title"
                type="text"
                placeholder="Enter epic title"
                required
                :class="{ 'error': errors.title }"
              />
              <span v-if="errors.title" class="error-message">{{ errors.title }}</span>
            </div>

            <div class="form-group">
              <label for="epicCode">Epic Code *</label>
              <input
                id="epicCode"
                v-model="form.epicCode"
                type="text"
                placeholder="e.g., PROJ-123"
                required
                :class="{ 'error': errors.epicCode }"
              />
              <span v-if="errors.epicCode" class="error-message">{{ errors.epicCode }}</span>
            </div>

            <div class="form-group">
              <label for="description">Description</label>
              <textarea
                id="description"
                v-model="form.description"
                placeholder="Enter epic description (optional)"
                rows="4"
              ></textarea>
            </div>

            <div class="form-actions">
              <button
                type="submit"
                class="btn btn-primary"
                :disabled="isSubmitting"
              >
                <span v-if="isSubmitting" class="spinner"></span>
                {{ isEditing ? 'Update Epic' : 'Add Epic' }}
              </button>

              <button
                v-if="isEditing"
                type="button"
                class="btn btn-secondary"
                @click="cancelEdit"
              >
                Cancel
              </button>
            </div>
          </form>
        </div>
      </div>

      <!-- Epics List -->
      <div class="epics-list-section">
        <div class="section-header">
          <h2>Saved Epics ({{ epicsStore.epicCount }})</h2>
          <button
            v-if="epicsStore.epics.length > 0"
            @click="showDeleteAllModal = true"
            class="btn btn-error btn-sm"
          >
            Clear All
          </button>
        </div>

        <!-- No Epics State -->
        <div v-if="epicsStore.epics.length === 0" class="empty-state">
          <div class="empty-icon">📋</div>
          <h3>No Epics Added Yet</h3>
          <p>Add your first epic using the form above to get started.</p>
        </div>

        <!-- Epics Grid -->
        <div v-else class="epics-grid">
          <div
            v-for="epic in epicsStore.epics"
            :key="epic.id"
            class="epic-item"
            :class="{ 'editing': editingId === epic.id }"
          >
            <div class="epic-header">
              <span class="epic-code">{{ epic.epicCode }}</span>
              <div class="epic-actions">
                <button
                  @click="editEpic(epic)"
                  class="btn-icon"
                  title="Edit Epic"
                >
                  ✏️
                </button>
                <router-link
                  :to="`/epic/${epic.epicCode}`"
                  class="btn-icon"
                  title="View Epic"
                >
                  👁️
                </router-link>
                <button
                  @click="deleteEpic(epic.id)"
                  class="btn-icon btn-delete"
                  title="Delete Epic"
                >
                  🗑️
                </button>
              </div>
            </div>

            <h4 class="epic-title">{{ epic.title }}</h4>

            <p v-if="epic.description" class="epic-description">
              {{ epic.description }}
            </p>

            <div class="epic-meta">
              <span class="meta-item">
                Created: {{ formatDate(epic.createdAt) }}
              </span>
              <span v-if="epic.updatedAt" class="meta-item">
                Updated: {{ formatDate(epic.updatedAt) }}
              </span>
            </div>
          </div>
        </div>
      </div>

      <!-- Delete All Confirmation Modal -->
      <div v-if="showDeleteAllModal" class="modal-overlay" @click="showDeleteAllModal = false">
        <div class="modal" @click.stop>
          <h3>Clear All Epics</h3>
          <p>Are you sure you want to remove all saved epics? This action cannot be undone.</p>
          <div class="modal-actions">
            <button @click="showDeleteAllModal = false" class="btn btn-secondary">
              Cancel
            </button>
            <button @click="clearAllEpics" class="btn btn-error">
              Clear All
            </button>
          </div>
        </div>
      </div>

      <!-- Success Message -->
      <div v-if="successMessage" class="alert alert-success">
        {{ successMessage }}
      </div>
    </div>
  </div>
</template>

<script>
import { ref, reactive, computed, nextTick } from "vue";
import { useEpicsStore } from "@/stores/epics";

export default {
  name: "AdminEpics",
  setup() {
    const epicsStore = useEpicsStore();

    // Form state
    const form = reactive({
      title: "",
      epicCode: "",
      description: ""
    });

    const errors = reactive({});
    const isEditing = ref(false);
    const editingId = ref(null);
    const isSubmitting = ref(false);
    const showDeleteAllModal = ref(false);
    const successMessage = ref("");

    // Validation
    const validateForm = () => {
      const newErrors = {};

      if (!form.title.trim()) {
        newErrors.title = "Title is required";
      }

      if (!form.epicCode.trim()) {
        newErrors.epicCode = "Epic code is required";
      } else if (!/^[A-Z]+-\d+$/i.test(form.epicCode.trim())) {
        newErrors.epicCode = "Epic code must be in format 'PROJECT-123'";
      } else {
        // Check for duplicate epic codes (excluding current epic when editing)
        const existingEpic = epicsStore.getEpicByCode(form.epicCode.trim());
        if (existingEpic && existingEpic.id !== editingId.value) {
          newErrors.epicCode = "Epic code already exists";
        }
      }

      Object.assign(errors, newErrors);
      return Object.keys(newErrors).length === 0;
    };

    // Form actions
    const saveEpic = async () => {
      // Clear previous errors
      Object.keys(errors).forEach(key => {
        delete errors[key];
      });

      if (!validateForm()) {
        return;
      }

      isSubmitting.value = true;

      try {
        // Simulate API delay
        await new Promise(resolve => setTimeout(resolve, 500));

        const epicData = {
          title: form.title.trim(),
          epicCode: form.epicCode.trim().toUpperCase(),
          description: form.description.trim()
        };

        if (isEditing.value) {
          epicsStore.updateEpic(editingId.value, epicData);
          showSuccess("Epic updated successfully!");
        } else {
          epicsStore.addEpic(epicData);
          showSuccess("Epic added successfully!");
        }

        resetForm();
      } catch (error) {
        console.error("Error saving epic:", error);
      } finally {
        isSubmitting.value = false;
      }
    };

    const editEpic = (epic) => {
      form.title = epic.title;
      form.epicCode = epic.epicCode;
      form.description = epic.description || "";

      isEditing.value = true;
      editingId.value = epic.id;

      // Scroll to form
      nextTick(() => {
        document.querySelector('.epic-form-section').scrollIntoView({
          behavior: 'smooth'
        });
      });
    };

    const cancelEdit = () => {
      resetForm();
    };

    const deleteEpic = (id) => {
      if (confirm("Are you sure you want to delete this epic?")) {
        epicsStore.removeEpic(id);

        // Cancel editing if we're editing the deleted epic
        if (editingId.value === id) {
          resetForm();
        }

        showSuccess("Epic deleted successfully!");
      }
    };

    const clearAllEpics = () => {
      epicsStore.clearAllEpics();
      resetForm();
      showDeleteAllModal.value = false;
      showSuccess("All epics cleared successfully!");
    };

    const resetForm = () => {
      form.title = "";
      form.epicCode = "";
      form.description = "";

      isEditing.value = false;
      editingId.value = null;

      // Clear errors
      Object.keys(errors).forEach(key => {
        delete errors[key];
      });
    };

    const showSuccess = (message) => {
      successMessage.value = message;
      setTimeout(() => {
        successMessage.value = "";
      }, 3000);
    };

    const formatDate = (dateString) => {
      const date = new Date(dateString);
      return date.toLocaleDateString() + " " + date.toLocaleTimeString([], {
        hour: '2-digit',
        minute: '2-digit'
      });
    };

    return {
      epicsStore,
      form,
      errors,
      isEditing,
      editingId,
      isSubmitting,
      showDeleteAllModal,
      successMessage,
      saveEpic,
      editEpic,
      cancelEdit,
      deleteEpic,
      clearAllEpics,
      formatDate
    };
  }
};
</script>

<style lang="scss" scoped>
.admin-epics-page {
  padding: var(--spacing-xl) 0;
  min-height: 100vh;
}

.page-header {
  text-align: center;
  margin-bottom: var(--spacing-xxl);

  h1 {
    margin-bottom: var(--spacing-sm);
  }

  p {
    color: var(--text-secondary);
    font-size: var(--font-lg);
  }
}

.epic-form-section {
  margin-bottom: var(--spacing-xxl);

  .card {
    max-width: 600px;
    margin: 0 auto;
  }

  h2 {
    margin-bottom: var(--spacing-lg);
    color: var(--text-primary);
  }
}

.epic-form {
  .form-group {
    margin-bottom: var(--spacing-lg);

    label {
      display: block;
      margin-bottom: var(--spacing-sm);
      font-weight: var(--font-medium);
      color: var(--text-primary);
    }

    input, textarea {
      &.error {
        border-color: var(--error-color);

        &:focus {
          border-color: var(--error-color);
          box-shadow: 0 0 0 2px var(--error-light);
        }
      }
    }

    .error-message {
      display: block;
      color: var(--error-color);
      font-size: var(--font-sm);
      margin-top: var(--spacing-xs);
    }
  }

  .form-actions {
    @include flex-between;
    gap: var(--spacing-md);

    .btn {
      min-width: 120px;
    }
  }
}

.epics-list-section {
  .section-header {
    @include flex-between;
    margin-bottom: var(--spacing-lg);

    h2 {
      margin: 0;
    }
  }
}

.empty-state {
  text-align: center;
  padding: var(--spacing-xxl);
  color: var(--text-secondary);

  .empty-icon {
    font-size: 4rem;
    margin-bottom: var(--spacing-lg);
    opacity: 0.5;
  }

  h3 {
    margin-bottom: var(--spacing-sm);
    color: var(--text-primary);
  }
}

.epics-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: var(--spacing-lg);
}

.epic-item {
  @include card-style;
  padding: var(--spacing-lg);
  transition: var(--transition-normal);

  &.editing {
    border-color: var(--primary-color);
    box-shadow: 0 0 0 2px var(--primary-light);
  }

  &:hover {
    transform: translateY(-2px);
    box-shadow: var(--shadow-md);
  }

  .epic-header {
    @include flex-between;
    margin-bottom: var(--spacing-md);

    .epic-code {
      background-color: var(--primary-light);
      color: var(--primary-color);
      padding: var(--spacing-xs) var(--spacing-sm);
      border-radius: var(--radius-md);
      font-weight: var(--font-medium);
      font-size: var(--font-sm);
    }

    .epic-actions {
      display: flex;
      gap: var(--spacing-xs);
    }
  }

  .epic-title {
    margin-bottom: var(--spacing-sm);
    color: var(--text-primary);
  }

  .epic-description {
    color: var(--text-secondary);
    font-size: var(--font-sm);
    margin-bottom: var(--spacing-md);
    display: -webkit-box;
    -webkit-line-clamp: 3;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }

  .epic-meta {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-xs);

    .meta-item {
      font-size: var(--font-xs);
      color: var(--text-muted);
    }
  }
}

.btn-icon {
  @include button-reset;
  padding: var(--spacing-xs);
  border-radius: var(--radius-sm);
  transition: var(--transition-fast);
  text-decoration: none;

  &:hover {
    background-color: var(--hover-bg);
  }

  &.btn-delete:hover {
    background-color: var(--error-light);
  }
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: var(--z-modal);
}

.modal {
  background-color: var(--card-bg);
  padding: var(--spacing-xl);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-xl);
  max-width: 400px;
  width: 90%;

  h3 {
    margin-bottom: var(--spacing-md);
    color: var(--text-primary);
  }

  p {
    margin-bottom: var(--spacing-lg);
    color: var(--text-secondary);
  }

  .modal-actions {
    display: flex;
    gap: var(--spacing-md);
    justify-content: flex-end;
  }
}

@media (max-width: 768px) {
  .epics-grid {
    grid-template-columns: 1fr;
  }

  .epic-form .form-actions {
    flex-direction: column;

    .btn {
      width: 100%;
    }
  }

  .section-header {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--spacing-md);
  }
}
</style>
