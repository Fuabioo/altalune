<template>
    <div class="update-info">
        <p class="text-muted">Last updated: {{ formatDate(lastUpdated) }}</p>
    </div>
</template>

<script>
export default {
    name: "UpdateInfo",
    props: {
        lastUpdated: {
            type: [String, Date],
            default: () => new Date().toISOString(),
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

.update-info {
    text-align: center;
    padding: var(--spacing-lg) 0;
    border-top: 1px solid var(--border-color);
    margin-top: var(--spacing-xxl);
}

.text-muted {
    margin: 0;
    color: var(--text-muted);
    font-size: var(--font-sm);
    font-style: italic;
}
</style>
