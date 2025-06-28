<template>
    <div class="issues-section">
        <div class="card">
            <h2>{{ title }} ({{ issues.length }})</h2>
            <div v-if="issues.length" class="issues-table">
                <div class="table-header">
                    <span>Key</span>
                    <span>Summary</span>
                    <span>Status</span>
                    <span>Assignee</span>
                </div>
                <div v-for="issue in issues" :key="issue.key" class="table-row">
                    <span class="issue-key">{{ issue.key }}</span>
                    <span class="issue-summary">{{
                        issue.fields?.summary
                    }}</span>
                    <span class="issue-status">
                        <span
                            class="status-badge"
                            :class="
                                getStatusClass(
                                    issue.fields?.statusCategory?.key,
                                )
                            "
                        >
                            {{ issue.fields?.status?.name }}
                        </span>
                    </span>
                    <span class="issue-assignee">
                        {{
                            issue.fields?.assignee?.displayName || "Unassigned"
                        }}
                    </span>
                </div>
            </div>
            <div v-else class="no-issues">
                <p>{{ emptyMessage }}</p>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "IssuesTable",
    props: {
        title: {
            type: String,
            default: "Epic Issues",
        },
        issues: {
            type: Array,
            default: () => [],
        },
        emptyMessage: {
            type: String,
            default: "No issues found for this epic.",
        },
    },
    methods: {
        getStatusClass(classification) {
            const statusClasses = {
                new: "status-todo",
                indeterminate: "status-progress",
                done: "status-done",
            };
            return statusClasses[classification] || "status-default";
        },
    },
};
</script>

<style lang="scss" scoped>
@use "@/styles/variables.scss" as *;

.issues-section {
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

.issues-table {
    width: 100%;
    border-radius: var(--radius-md);
    overflow: hidden;
    border: 1px solid var(--border-color);
}

.table-header {
    display: grid;
    grid-template-columns: 120px 1fr 150px 120px 150px;
    background: var(--border-light);
    padding: var(--spacing-md) var(--spacing-md);
    font-weight: var(--font-semibold);
    font-size: var(--font-sm);
    color: var(--text-secondary);
    text-transform: uppercase;
    letter-spacing: 0.5px;
    border-bottom: 1px solid var(--border-color);
}

.table-row {
    display: grid;
    grid-template-columns: 120px 1fr 150px 120px 150px;
    padding: var(--spacing-md) var(--spacing-md);
    border-bottom: 1px solid var(--border-light);
    align-items: center;
    transition: var(--transition-normal);
}

.table-row:hover {
    background-color: var(--hover-bg);
}

.issue-key {
    font-weight: var(--font-semibold);
    color: var(--primary-color);
    font-size: var(--font-sm);
}

.issue-summary {
    color: var(--text-primary);
    font-size: var(--font-md);
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}

.issue-type {
    display: flex;
    align-items: center;
    gap: var(--spacing-xs);
    font-size: var(--font-sm);
    color: var(--text-secondary);
}

.type-icon {
    width: 16px;
    height: 16px;
    flex-shrink: 0;
}

.status-badge {
    display: inline-block;
    padding: var(--spacing-xs) var(--spacing-sm);
    border-radius: var(--radius-xl);
    font-size: var(--font-xs);
    font-weight: var(--font-semibold);
    text-transform: uppercase;
    letter-spacing: 0.3px;
}

.status-badge.status-todo {
    background-color: var(--ctp-overlay1);
    color: var(--text-inverse);
}

.status-badge.status-progress {
    background-color: var(--ctp-blue);
    color: var(--text-inverse);
}

.status-badge.status-done {
    background-color: var(--ctp-green);
    color: var(--text-inverse);
}

.status-badge.status-blocked {
    background-color: var(--ctp-red);
    color: var(--text-inverse);
}

.status-badge.status-review {
    background-color: var(--ctp-yellow);
    color: var(--ctp-base);
}

.status-badge.status-default {
    background-color: var(--text-muted);
    color: var(--text-inverse);
}

.issue-assignee {
    font-size: var(--font-sm);
    color: var(--text-secondary);
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}

.no-issues {
    text-align: center;
    padding: var(--spacing-xxl) var(--spacing-lg);
    color: var(--text-secondary);
}

.no-issues p {
    margin: 0;
    font-size: var(--font-md);
}

@media (max-width: 768px) {
    .table-header,
    .table-row {
        grid-template-columns: 1fr;
        gap: var(--spacing-sm);
    }

    .table-header {
        display: none;
    }

    .table-row {
        padding: var(--spacing-md);
        border-radius: var(--radius-md);
        margin-bottom: var(--spacing-sm);
        background: var(--border-light);
        border: 1px solid var(--border-color);
    }

    .table-row span {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: var(--spacing-xs) 0;
    }

    .table-row span::before {
        content: attr(data-label);
        font-weight: var(--font-semibold);
        color: var(--text-secondary);
        font-size: var(--font-sm);
        text-transform: uppercase;
        letter-spacing: 0.5px;
    }

    .issue-key::before {
        content: "Key: ";
    }

    .issue-summary::before {
        content: "Summary: ";
    }

    .issue-type::before {
        content: "Type: ";
    }

    .issue-status::before {
        content: "Status: ";
    }

    .issue-assignee::before {
        content: "Assignee: ";
    }
}
</style>
