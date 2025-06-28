<template>
    <div class="progress-section">
        <div class="card">
            <h2>{{ title }}</h2>

            <!-- Completion Progress -->
            <div class="progress-item">
                <div class="progress-header">
                    <span class="progress-label">Completed Tasks</span>
                    <span class="progress-value"
                        >{{ Math.round(completionPercentage) }}%</span
                    >
                </div>
                <div class="progress-bar">
                    <div
                        class="progress-fill completion-fill"
                        :style="{ width: completionPercentage + '%' }"
                    ></div>
                </div>
            </div>

            <!-- In Progress -->
            <div class="progress-item">
                <div class="progress-header">
                    <span class="progress-label">In Progress Tasks</span>
                    <span class="progress-value"
                        >{{ Math.round(progressPercentage) }}%</span
                    >
                </div>
                <div class="progress-bar">
                    <div
                        class="progress-fill progress-fill-blue"
                        :style="{ width: progressPercentage + '%' }"
                    ></div>
                </div>
            </div>

            <div class="summary-labels">
                <span
                    >{{ completedIssues }} completed • {{ inProgressIssues }} in
                    progress • {{ totalIssues }} total</span
                >
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "ProgressBar",
    props: {
        title: {
            type: String,
            default: "Epic Progress",
        },
        totalIssues: {
            type: Number,
            default: 0,
        },
        completedIssues: {
            type: Number,
            default: 0,
        },
        inProgressIssues: {
            type: Number,
            default: 0,
        },
        completionPercentage: {
            type: Number,
            default: 0,
        },
        progressPercentage: {
            type: Number,
            default: 0,
        },
    },
};
</script>

<style lang="scss" scoped>
@use "@/styles/variables.scss" as *;

.progress-section {
    margin-bottom: var(--spacing-xxl);
}

.card {
    @include card-style;
    padding: var(--spacing-lg);
}

.card h2 {
    margin: 0 0 var(--spacing-lg) 0;
    color: var(--text-primary);
    font-size: var(--font-xxl);
    font-weight: var(--font-semibold);
}

.progress-item {
    margin-bottom: var(--spacing-lg);
}

.progress-item:last-of-type {
    margin-bottom: var(--spacing-md);
}

.progress-header {
    @include flex-between;
    margin-bottom: var(--spacing-xs);
}

.progress-label {
    font-size: var(--font-md);
    color: var(--text-primary);
    font-weight: var(--font-medium);
}

.progress-value {
    font-size: var(--font-md);
    font-weight: var(--font-semibold);
    color: var(--text-primary);
}

.progress-bar {
    width: 100%;
    height: 8px;
    background-color: var(--border-light);
    border-radius: var(--radius-md);
    overflow: hidden;
}

.progress-fill {
    height: 100%;
    border-radius: var(--radius-md);
    transition: width 0.3s ease;
}

.completion-fill {
    background: var(--ctp-green);
}

.progress-fill-blue {
    background: var(--ctp-blue);
}

.summary-labels {
    text-align: center;
    font-size: var(--font-sm);
    color: var(--text-secondary);
    font-weight: var(--font-medium);
}
</style>
