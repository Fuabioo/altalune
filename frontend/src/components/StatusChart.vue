<template>
    <div class="chart-card">
        <h3>{{ title }}</h3>
        <div class="chart-placeholder">
            <div class="chart-bar-container">
                <div
                    v-for="status in statusData"
                    :key="status.name"
                    class="chart-bar"
                >
                    <span class="bar-label">{{ status.name }}</span>
                    <div class="bar-track">
                        <div
                            class="bar-fill"
                            :style="{
                                width: (status.count / maxCount) * 100 + '%',
                                backgroundColor: status.color,
                            }"
                        ></div>
                    </div>
                    <span class="bar-count">{{ status.count }}</span>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "StatusChart",
    props: {
        title: {
            type: String,
            default: "Status Distribution",
        },
        statusCounts: {
            type: Object,
            default: () => ({}),
        },
    },
    computed: {
        statusData() {
            // Use CSS variables for consistent theming
            const colors = {
                new: "var(--ctp-overlay1)",
                indeterminate: "var(--ctp-blue)",
                done: "var(--ctp-green)",
            };

            return Object.entries(this.statusCounts).map(
                ([name, { classification, count }]) => ({
                    name,
                    classification,
                    count,
                    color: colors[classification] || "var(--text-muted)",
                }),
            );
        },
        maxCount() {
            return Math.max(...this.statusData.map((s) => s.count), 1);
        },
    },
};
</script>

<style lang="scss" scoped>
@use "@/styles/variables.scss" as *;

.chart-card {
    @include card-style;
    padding: var(--spacing-lg);
}

.chart-card h3 {
    margin: 0 0 var(--spacing-lg) 0;
    color: var(--text-primary);
    font-size: var(--font-lg);
    font-weight: var(--font-semibold);
}

.chart-placeholder {
    width: 100%;
    height: auto;
    padding: var(--spacing-md);
}

.chart-bar-container {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-lg);
    width: 100%;
}

.chart-bar {
    display: flex;
    align-items: center;
    gap: var(--spacing-md);
    width: 100%;
}

.bar-label {
    font-size: var(--font-sm);
    color: var(--text-secondary);
    font-weight: var(--font-medium);
    min-width: 100px;
    text-align: left;
    flex-shrink: 0;
}

.bar-track {
    flex: 1;
    height: 24px;
    background-color: var(--surface-variant);
    border-radius: var(--radius-sm);
    position: relative;
    overflow: hidden;
}

.bar-fill {
    height: 100%;
    min-width: 2px;
    border-radius: var(--radius-sm);
    transition: width 0.3s ease;
}

.bar-count {
    font-size: var(--font-md);
    font-weight: var(--font-semibold);
    color: var(--text-primary);
    min-width: 40px;
    text-align: right;
    flex-shrink: 0;
}

@media (max-width: 768px) {
    .chart-bar-container {
        gap: var(--spacing-md);
    }

    .bar-label {
        min-width: 80px;
        font-size: var(--font-xs);
    }

    .bar-track {
        height: 20px;
    }

    .bar-count {
        font-size: var(--font-sm);
        min-width: 35px;
    }
}

@media (max-width: 480px) {
    .chart-placeholder {
        padding: var(--spacing-sm);
    }

    .chart-bar-container {
        gap: var(--spacing-sm);
    }

    .chart-bar {
        gap: var(--spacing-sm);
    }

    .bar-label {
        min-width: 70px;
        font-size: var(--font-xs);
    }

    .bar-track {
        height: 18px;
    }

    .bar-count {
        font-size: var(--font-sm);
        min-width: 30px;
    }
}
</style>
