<template>
    <div class="loading-state">
        <div class="loading-card">
            <div class="spinner-large"></div>
            <h3>{{ title }}</h3>
            <p>{{ message }}</p>
        </div>
    </div>
</template>

<script>
export default {
    name: "LoadingState",
    props: {
        title: {
            type: String,
            default: "Loading Epic Data...",
        },
        message: {
            type: String,
            default: "Fetching analytics from Jira API",
        },
    },
};
</script>

<style lang="scss" scoped>
@use "@/styles/variables.scss" as *;

.loading-state {
    @include flex-center;
    min-height: 300px;
}

.loading-card {
    @include card-style;
    text-align: center;
    padding: var(--spacing-xxl);
    max-width: 400px;
}

.spinner-large {
    width: 48px;
    height: 48px;
    border: 4px solid var(--border-color);
    border-top: 4px solid var(--primary-color);
    border-radius: 50%;
    animation: spin 1s linear infinite;
    margin: 0 auto var(--spacing-lg);
}

.loading-card h3 {
    margin: 0 0 var(--spacing-sm) 0;
    color: var(--text-primary);
    font-size: var(--font-xl);
    font-weight: var(--font-semibold);
}

.loading-card p {
    margin: 0;
    color: var(--text-secondary);
    font-size: var(--font-md);
}

@keyframes spin {
    0% {
        transform: rotate(0deg);
    }
    100% {
        transform: rotate(360deg);
    }
}
</style>
