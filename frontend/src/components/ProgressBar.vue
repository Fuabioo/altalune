<template>
    <div class="progress-section">
        <div class="card">
            <div class="progress-container">
                <div class="progress-header">
                    <span class="progress-title">{{ title }}</span>
                    <span class="progress-summary">
                        {{ completedIssues }}/{{ totalIssues }} completed
                        <span v-if="inProgressIssues > 0">
                            • {{ inProgressIssues }} in progress</span
                        >
                    </span>
                </div>

                <div class="progress-bar-container">
                    <div class="progress-bar">
                        <!-- Completed section -->
                        <div
                            class="progress-fill completion-fill"
                            :style="{ width: completionPercentage + '%' }"
                        ></div>
                        <!-- In progress section -->
                        <div
                            class="progress-fill progress-fill-blue"
                            :style="{
                                width: progressPercentage + '%',
                                left: completionPercentage + '%',
                            }"
                        ></div>
                    </div>
                    <div class="progress-percentage">
                        {{
                            Math.round(
                                completionPercentage + progressPercentage,
                            )
                        }}%
                    </div>
                </div>
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
    margin-bottom: var(--spacing-lg);
}

.card {
    @include card-style;
    padding: var(--spacing-md);
}

.progress-container {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-xs);
}

.progress-header {
    @include flex-between;
    align-items: baseline;
}

.progress-title {
    font-size: var(--font-lg);
    color: var(--text-primary);
    font-weight: var(--font-semibold);
    margin: 0;
}

.progress-summary {
    font-size: var(--font-sm);
    color: var(--text-secondary);
    font-weight: var(--font-medium);
}

.progress-bar-container {
    display: flex;
    align-items: center;
    gap: var(--spacing-sm);
}

.progress-bar {
    flex: 1;
    height: 6px;
    background-color: var(--border-light);
    border-radius: var(--radius-sm);
    overflow: hidden;
    position: relative;
}

.progress-fill {
    height: 100%;
    position: absolute;
    top: 0;
    border-radius: var(--radius-sm);
    transition:
        width 0.3s ease,
        left 0.3s ease;
}

.completion-fill {
    background: var(--ctp-green);
    left: 0;
}

.progress-fill-blue {
    background: var(--ctp-blue);
}

.progress-percentage {
    font-size: var(--font-sm);
    font-weight: var(--font-semibold);
    color: var(--text-primary);
    min-width: 3ch;
    text-align: right;
}
</style>
