<template>
    <div class="charts-section">
        <div class="chart-grid">
            <StatusChart
                :title="statusChartTitle"
                :statusCounts="statusCounts"
            />
        </div>
    </div>
</template>

<script>
import StatusChart from "./StatusChart.vue";

export default {
    name: "ChartsSection",
    components: {
        StatusChart,
    },
    props: {
        statusCounts: {
            type: Object,
            default: () => ({}),
        },
        typeCounts: {
            type: Object,
            default: () => ({}),
        },
        statusChartTitle: {
            type: String,
            default: "Status Distribution",
        },
        typesChartTitle: {
            type: String,
            default: "Issue Types",
        },
    },
};
</script>

<style lang="scss" scoped>
@use "@/styles/variables.scss" as *;

.charts-section {
    margin-bottom: var(--spacing-xxl);
}

.chart-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
    gap: var(--spacing-xl);
}

@media (max-width: 768px) {
    .chart-grid {
        grid-template-columns: 1fr;
        gap: var(--spacing-lg);
    }
}

@media (max-width: 480px) {
    .chart-grid {
        gap: var(--spacing-md);
    }
}
</style>
