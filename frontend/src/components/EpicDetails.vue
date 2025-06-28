<template>
    <div class="details-section">
        <div class="card">
            <h2>{{ title }}</h2>
            <div class="details-grid">
                <div
                    class="detail-item"
                    v-for="detail in detailItems"
                    :key="detail.label"
                >
                    <label>{{ detail.label }}</label>
                    <div class="value">{{ detail.value }}</div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "EpicDetails",
    props: {
        title: {
            type: String,
            default: "Detailed Statistics",
        },
        epicKey: {
            type: String,
            default: "",
        },
        createdDate: {
            type: String,
            default: "",
        },
        updatedDate: {
            type: String,
            default: "",
        },
        reporter: {
            type: String,
            default: "Unknown",
        },
        assignee: {
            type: String,
            default: "Unassigned",
        },
        priority: {
            type: String,
            default: "None",
        },
    },
    computed: {
        detailItems() {
            return [
                {
                    label: "Epic Key",
                    value: this.epicKey,
                },
                {
                    label: "Created Date",
                    value: this.formatDate(this.createdDate),
                },
                {
                    label: "Last Updated",
                    value: this.formatDate(this.updatedDate),
                },
                {
                    label: "Reporter",
                    value: this.reporter,
                },
                {
                    label: "Assignee",
                    value: this.assignee,
                },
                {
                    label: "Priority",
                    value: this.priority,
                },
            ];
        },
    },
    methods: {
        formatDate(dateString) {
            if (!dateString) return "Unknown";

            const date = new Date(dateString);
            if (isNaN(date.getTime())) return "Invalid Date";

            return (
                date.toLocaleDateString() +
                " " +
                date.toLocaleTimeString([], {
                    hour: "2-digit",
                    minute: "2-digit",
                })
            );
        },
    },
};
</script>

<style lang="scss" scoped>
@use "@/styles/variables.scss" as *;

.details-section {
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

.details-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: var(--spacing-lg);
}

.detail-item {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-xs);
}

.detail-item label {
    font-size: var(--font-sm);
    font-weight: var(--font-semibold);
    color: var(--text-secondary);
    text-transform: uppercase;
    letter-spacing: 0.5px;
}

.detail-item .value {
    font-size: var(--font-md);
    font-weight: var(--font-medium);
    color: var(--text-primary);
    word-break: break-word;
}

@media (max-width: 768px) {
    .details-grid {
        grid-template-columns: 1fr;
        gap: var(--spacing-md);
    }

    .detail-item {
        padding: var(--spacing-md);
        background: var(--border-light);
        border-radius: var(--radius-md);
    }
}
</style>
