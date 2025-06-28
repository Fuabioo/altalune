<template>
    <div class="epic-header">
        <div class="epic-info">
            <div class="epic-code-badge">{{ epicCode }}</div>
            <h1>{{ title || "Loading..." }}</h1>
            <p v-if="description" class="epic-description">
                {{ description }}
            </p>
        </div>
        <div class="epic-actions">
            <button
                @click="onRefresh"
                class="btn btn-secondary"
                :disabled="isLoading"
            >
                <span v-if="isLoading" class="spinner"></span>
                🔄 Refresh
            </button>
            <button @click="onManageEpics" class="btn btn-primary">
                📋 Manage Epics
            </button>
        </div>
    </div>
</template>

<script>
export default {
    name: "EpicHeader",
    props: {
        epicCode: {
            type: String,
            required: true,
        },
        title: {
            type: String,
            default: "",
        },
        description: {
            type: String,
            default: "",
        },
        isLoading: {
            type: Boolean,
            default: false,
        },
    },
    emits: ["refresh", "manage-epics"],
    methods: {
        onRefresh() {
            this.$emit("refresh");
        },
        onManageEpics() {
            this.$emit("manage-epics");
        },
    },
};
</script>

<style lang="scss" scoped>
@use "@/styles/variables.scss" as *;

.epic-header {
    @include flex-between;
    margin-bottom: var(--spacing-xxl);
    padding: var(--spacing-lg);
    background: var(--card-bg);
    border-radius: var(--radius-lg);
    box-shadow: var(--shadow-md);
    border: 1px solid var(--border-color);
}

.epic-info {
    flex: 1;
}

.epic-code-badge {
    display: inline-block;
    background: var(--primary-light);
    color: var(--primary-color);
    padding: var(--spacing-xs) var(--spacing-md);
    border-radius: var(--radius-xl);
    font-size: var(--font-xs);
    font-weight: var(--font-semibold);
    margin-bottom: var(--spacing-md);
    text-transform: uppercase;
    letter-spacing: 0.5px;
}

.epic-info h1 {
    margin: 0 0 var(--spacing-md) 0;
    color: var(--text-primary);
    font-size: var(--font-xxxl);
    font-weight: var(--font-semibold);
    line-height: 1.2;
}

.epic-description {
    margin: 0;
    color: var(--text-secondary);
    font-size: var(--font-lg);
    line-height: 1.5;
    max-width: 600px;
}

.epic-actions {
    display: flex;
    gap: var(--spacing-md);
    flex-shrink: 0;
}

.btn {
    @include button-reset;
    display: flex;
    align-items: center;
    gap: var(--spacing-sm);
    padding: var(--spacing-sm) var(--spacing-md);
    border-radius: var(--radius-md);
    cursor: pointer;
    font-size: var(--font-md);
    font-weight: var(--font-medium);
    transition: var(--transition-normal);
    text-decoration: none;
    border: 1px solid transparent;
    min-height: 38px;
}

.btn-secondary {
    background-color: var(--card-bg);
    color: var(--text-primary);
    border-color: var(--border-color);
}

.btn-secondary:hover:not(:disabled) {
    background-color: var(--hover-bg);
    border-color: var(--primary-color);
}

.btn-secondary:disabled {
    opacity: 0.6;
    cursor: not-allowed;
}

.btn-primary {
    background-color: var(--primary-color);
    color: var(--text-inverse);
}

.btn-primary:hover {
    background-color: var(--primary-hover);
}

.spinner {
    width: 14px;
    height: 14px;
    border: 2px solid transparent;
    border-top: 2px solid currentColor;
    border-radius: 50%;
    animation: spin 1s linear infinite;
}

@keyframes spin {
    0% {
        transform: rotate(0deg);
    }
    100% {
        transform: rotate(360deg);
    }
}

@media (max-width: 768px) {
    .epic-header {
        flex-direction: column;
        gap: var(--spacing-lg);
    }

    .epic-actions {
        width: 100%;
        justify-content: stretch;
    }

    .epic-actions .btn {
        flex: 1;
        justify-content: center;
    }
}
</style>
