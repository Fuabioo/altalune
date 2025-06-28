<template>
    <div class="stats-overview">
        <div class="stat-card" v-for="stat in stats" :key="stat.label">
            <div class="stat-icon">{{ stat.icon }}</div>
            <div class="stat-content">
                <h3>{{ stat.value }}</h3>
                <p>{{ stat.label }}</p>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "StatsOverview",
    props: {
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
    computed: {
        stats() {
            return [
                {
                    icon: "📊",
                    value: this.totalIssues,
                    label: "Total Issues",
                },
                {
                    icon: "✅",
                    value: this.completedIssues,
                    label: "Completed",
                },
                {
                    icon: "⏳",
                    value: this.inProgressIssues,
                    label: "In Progress",
                },
                {
                    icon: "📈",
                    value: `${Math.round(this.completionPercentage)}%`,
                    label: "Complete",
                },
                {
                    icon: "⚡",
                    value: `${Math.round(this.progressPercentage)}%`,
                    label: "In Progress",
                },
            ];
        },
    },
};
</script>

<style lang="scss" scoped>
@use "@/styles/variables.scss" as *;

.stats-overview {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
    gap: var(--spacing-lg);
    margin-bottom: var(--spacing-xxl);
}

.stat-card {
    @include card-style;
    display: flex;
    align-items: center;
    padding: var(--spacing-lg);
    transition:
        transform var(--transition-normal),
        box-shadow var(--transition-normal);
}

.stat-card:hover {
    transform: translateY(-2px);
    box-shadow: var(--shadow-lg);
}

.stat-icon {
    font-size: 2rem;
    margin-right: var(--spacing-md);
    flex-shrink: 0;
}

.stat-content {
    flex: 1;
}

.stat-content h3 {
    margin: 0 0 var(--spacing-xs) 0;
    font-size: var(--font-xxxl);
    font-weight: var(--font-bold);
    color: var(--text-primary);
    line-height: 1;
}

.stat-content p {
    margin: 0;
    font-size: var(--font-md);
    color: var(--text-secondary);
    font-weight: var(--font-medium);
}

@media (max-width: 768px) {
    .stats-overview {
        grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
        gap: var(--spacing-md);
    }

    .stat-card {
        padding: var(--spacing-md);
    }

    .stat-icon {
        font-size: 1.5rem;
        margin-right: var(--spacing-sm);
    }

    .stat-content h3 {
        font-size: var(--font-xxl);
    }

    .stat-content p {
        font-size: var(--font-sm);
    }
}
</style>
