<template>
    <div class="error-state">
        <div class="error-card">
            <div class="error-icon">❌</div>
            <h3>{{ title }}</h3>
            <p>{{ message }}</p>
            <div class="error-details" v-if="details">
                <details>
                    <summary>Error Details</summary>
                    <pre><code>{{ formatDetails(details) }}</code></pre>
                </details>
            </div>
            <button @click="onRetry" class="btn btn-primary" v-if="showRetry">
                Try Again
            </button>
        </div>
    </div>
</template>

<script>
export default {
    name: "ErrorState",
    props: {
        title: {
            type: String,
            default: "Failed to Load Epic Data",
        },
        message: {
            type: String,
            default: "An error occurred while fetching data",
        },
        details: {
            type: [Object, String, Error],
            default: null,
        },
        showRetry: {
            type: Boolean,
            default: true,
        },
    },
    emits: ["retry"],
    methods: {
        onRetry() {
            this.$emit("retry");
        },
        formatDetails(details) {
            if (typeof details === "string") {
                return details;
            }
            if (details instanceof Error) {
                return details.stack || details.message;
            }
            return JSON.stringify(details, null, 2);
        },
    },
};
</script>

<style lang="scss" scoped>
@use "@/styles/variables.scss" as *;

.error-state {
    @include flex-center;
    min-height: 300px;
}

.error-card {
    @include card-style;
    text-align: center;
    padding: var(--spacing-xxl);
    max-width: 500px;
}

.error-icon {
    font-size: 3rem;
    margin-bottom: var(--spacing-lg);
}

.error-card h3 {
    margin: 0 0 var(--spacing-sm) 0;
    color: var(--error-color);
    font-size: var(--font-xl);
    font-weight: var(--font-semibold);
}

.error-card p {
    margin: 0 0 var(--spacing-lg) 0;
    color: var(--text-secondary);
    font-size: var(--font-md);
}

.error-details {
    margin: var(--spacing-lg) 0;
    text-align: left;
}

.error-details details {
    background: var(--border-light);
    border: 1px solid var(--border-color);
    border-radius: var(--radius-md);
    padding: var(--spacing-sm);
}

.error-details summary {
    cursor: pointer;
    font-weight: var(--font-semibold);
    color: var(--text-primary);
    padding: var(--spacing-xs) 0;
    transition: var(--transition-normal);
}

.error-details summary:hover {
    color: var(--primary-color);
}

.error-details pre {
    margin: var(--spacing-sm) 0 0 0;
    padding: var(--spacing-sm);
    background: var(--card-bg);
    border: 1px solid var(--border-color);
    border-radius: var(--radius-sm);
    overflow-x: auto;
    font-size: var(--font-sm);
    color: var(--error-color);
}

.btn {
    @include button-reset;
    padding: var(--spacing-sm) var(--spacing-md);
    border-radius: var(--radius-md);
    cursor: pointer;
    font-size: var(--font-md);
    font-weight: var(--font-medium);
    transition: var(--transition-normal);
    border: 1px solid transparent;
    min-height: 38px;
}

.btn-primary {
    background-color: var(--primary-color);
    color: var(--text-inverse);
}

.btn-primary:hover {
    background-color: var(--primary-hover);
}

.btn-primary:disabled {
    opacity: 0.6;
    cursor: not-allowed;
}
</style>
