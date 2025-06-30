<template>
    <div class="epic-page">
        <div class="container">
            <!-- Epic Header -->
            <EpicHeader
                :epicCode="epicCode"
                :title="epicData?.title"
                :description="epicData?.description"
                :isLoading="isLoading"
                @refresh="refreshData"
                @manage-epics="handleManageEpics"
            />

            <!-- Loading State -->
            <LoadingState v-if="isLoading && !apiData" />

            <!-- Error State -->
            <ErrorState
                v-else-if="error"
                :message="error.message"
                :details="error.details"
                @retry="refreshData"
            />

            <!-- Epic Analytics -->
            <div v-else-if="apiData" class="epic-analytics">
                <!-- Quick Stats -->
                <StatsOverview
                    :totalIssues="apiData.stats?.total || 0"
                    :completedIssues="apiData.stats?.done || 0"
                    :inProgressIssues="apiData.stats?.inProgress || 0"
                    :completionPercentage="apiData.stats?.percentage || 0"
                    :progressPercentage="apiData.stats?.progressPer || 0"
                />

                <!-- Progress Bar -->
                <ProgressBar
                    :totalIssues="apiData.stats?.total || 0"
                    :completedIssues="apiData.stats?.done || 0"
                    :inProgressIssues="apiData.stats?.inProgress || 0"
                    :completionPercentage="apiData.stats?.percentage || 0"
                    :progressPercentage="apiData.stats?.progressPer || 0"
                />

                <!-- Dependency Graph -->
                <DependencyGraph
                    v-if="apiData.graph"
                    :jiraBaseUrl="apiData.jiraBaseUrl"
                    :graphData="apiData.graph"
                />

                <!-- Charts Section -->
                <ChartsSection
                    :statusCounts="apiData.statusCounts || {}"
                    :typeCounts="apiData.typeCounts || {}"
                />

                <!-- Issues List -->
                <IssuesTable :issues="apiData.issues || []" />

                <!-- Last Updated -->
                <UpdateInfo />
            </div>
        </div>
    </div>
</template>

<script>
import { ref, computed, onMounted, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useEpicsStore } from "@/stores/epics";

// Import components
import EpicHeader from "@/components/EpicHeader.vue";
import LoadingState from "@/components/LoadingState.vue";
import ErrorState from "@/components/ErrorState.vue";
import StatsOverview from "@/components/StatsOverview.vue";
import ProgressBar from "@/components/ProgressBar.vue";
import ChartsSection from "@/components/ChartsSection.vue";
import DependencyGraph from "@/components/DependencyGraph.vue";
import IssuesTable from "@/components/IssuesTable.vue";
import UpdateInfo from "@/components/UpdateInfo.vue";

export default {
    name: "EpicPage",
    components: {
        EpicHeader,
        LoadingState,
        ErrorState,
        StatsOverview,
        ProgressBar,
        ChartsSection,
        DependencyGraph,
        IssuesTable,
        UpdateInfo,
    },
    props: {
        epicCode: {
            type: String,
            required: true,
        },
    },
    setup(props) {
        const route = useRoute();
        const router = useRouter();
        const epicsStore = useEpicsStore();

        const isLoading = ref(false);
        const apiData = ref(null);
        const error = ref(null);

        // Get epic data from store
        const epicData = computed(() => {
            return epicsStore.getEpicByCode(props.epicCode);
        });

        // Computed values for charts and stats
        const statusDistribution = computed(() => {
            if (!apiData.value?.statusCounts) return [];

            const colors = {
                new: "#ddd",
                indeterminate: "#0052cc",
                done: "#36b37e",
            };

            return Object.entries(apiData.value.statusCounts).map(
                ([name, { classification, count }]) => ({
                    name,
                    count,
                    classification,
                    color: colors[classification] || "#6b7280",
                }),
            );
        });

        const maxStatusCount = computed(() => {
            return Math.max(...statusDistribution.value.map((s) => s.count), 1);
        });

        // Methods
        const fetchEpicData = async () => {
            isLoading.value = true;
            error.value = null;

            try {
                const response = await fetch(`/api/epic/${props.epicCode}`);

                if (!response.ok) {
                    throw new Error(
                        `HTTP ${response.status}: ${response.statusText}`,
                    );
                }

                const data = await response.json();
                apiData.value = data;
                // if the jira base url doesn't have a https:// prefix, add it
                if (!apiData.value.jiraBaseUrl.startsWith("https://")) {
                    apiData.value.jiraBaseUrl = `https://${apiData.value.jiraBaseUrl}`;
                }
            } catch (err) {
                console.error("Error fetching epic data:", err);
                error.value = {
                    message: err.message,
                    details: err,
                };
            } finally {
                isLoading.value = false;
            }
        };

        const refreshData = () => {
            fetchEpicData();
        };

        const handleManageEpics = () => {
            router.push("/admin/epics");
        };

        // Watch for route changes
        watch(
            () => props.epicCode,
            () => {
                fetchEpicData();
            },
        );

        // Load data on mount
        onMounted(() => {
            fetchEpicData();
        });

        return {
            isLoading,
            apiData,
            error,
            epicData,
            refreshData,
            handleManageEpics,
        };
    },
};
</script>

<style lang="scss" scoped>
.epic-page {
    padding: var(--spacing-xl) 0;
    min-height: 100vh;
}

.epic-header {
    @include flex-between;
    margin-bottom: var(--spacing-xxl);
    gap: var(--spacing-lg);

    .epic-info {
        flex: 1;

        .epic-code-badge {
            display: inline-block;
            background-color: var(--primary-light);
            color: var(--primary-color);
            padding: var(--spacing-xs) var(--spacing-md);
            border-radius: var(--radius-lg);
            font-weight: var(--font-semibold);
            font-size: var(--font-sm);
            margin-bottom: var(--spacing-md);
        }

        h1 {
            margin-bottom: var(--spacing-sm);
            color: var(--text-primary);
        }

        .epic-description {
            color: var(--text-secondary);
            font-size: var(--font-lg);
            margin: 0;
        }
    }

    .epic-actions {
        display: flex;
        gap: var(--spacing-sm);
        flex-shrink: 0;
    }
}

.loading-state,
.error-state {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 400px;
}

.loading-card,
.error-card {
    @include card-style;
    padding: var(--spacing-xxl);
    text-align: center;
    max-width: 400px;

    .spinner-large {
        width: 40px;
        height: 40px;
        border: 4px solid var(--border-color);
        border-radius: 50%;
        border-top-color: var(--primary-color);
        animation: spin 1s ease-in-out infinite;
        margin: 0 auto var(--spacing-lg);
    }

    .error-icon {
        font-size: 3rem;
        margin-bottom: var(--spacing-lg);
    }

    h3 {
        margin-bottom: var(--spacing-md);
        color: var(--text-primary);
    }

    p {
        color: var(--text-secondary);
        margin-bottom: var(--spacing-lg);
    }

    .error-details {
        margin: var(--spacing-lg) 0;
        text-align: left;

        details {
            summary {
                cursor: pointer;
                font-weight: var(--font-medium);
                margin-bottom: var(--spacing-sm);
            }

            pre {
                background-color: var(--border-light);
                padding: var(--spacing-md);
                border-radius: var(--radius-md);
                overflow-x: auto;
                font-size: var(--font-sm);
            }
        }
    }
}

.stats-overview {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: var(--spacing-lg);
    margin-bottom: var(--spacing-xxl);
}

.stat-card {
    @include card-style;
    padding: var(--spacing-lg);
    text-align: center;
    transition: var(--transition-normal);

    &:hover {
        transform: translateY(-2px);
        box-shadow: var(--shadow-md);
    }

    .stat-icon {
        font-size: 2rem;
        margin-bottom: var(--spacing-sm);
    }

    .stat-content {
        h3 {
            font-size: var(--font-xl);
            color: var(--primary-color);
            margin-bottom: var(--spacing-xs);
        }

        p {
            color: var(--text-secondary);
            font-size: var(--font-sm);
            margin: 0;
        }
    }
}

.progress-section,
.charts-section,
.details-section,
.issues-section,
.dependency-graph {
    margin-bottom: var(--spacing-xxl);
}

.progress-bar {
    width: 100%;
    height: 20px;
    background-color: var(--border-light);
    border-radius: var(--radius-lg);
    overflow: hidden;
    margin: var(--spacing-md) 0;

    .progress-fill {
        height: 100%;
        background: linear-gradient(
            90deg,
            var(--primary-color),
            var(--success-color)
        );
        transition: width 0.5s ease;
    }
}

.progress-labels {
    @include flex-between;
    font-size: var(--font-sm);
    color: var(--text-secondary);
}

.chart-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: var(--spacing-lg);
}

.chart-card {
    h3 {
        margin-bottom: var(--spacing-lg);
        color: var(--text-primary);
    }
}

.chart-bar-container {
    display: flex;
    align-items: end;
    gap: var(--spacing-md);
    height: 200px;
    padding: var(--spacing-md) 0;

    .chart-bar {
        flex: 1;
        display: flex;
        flex-direction: column;
        align-items: center;
        height: 100%;

        .bar-fill {
            width: 100%;
            min-height: 4px;
            border-radius: var(--radius-sm) var(--radius-sm) 0 0;
            transition: height 0.5s ease;
        }

        .bar-label {
            font-size: var(--font-xs);
            color: var(--text-secondary);
            margin-top: var(--spacing-xs);
            text-align: center;
        }

        .bar-count {
            font-size: var(--font-sm);
            font-weight: var(--font-medium);
            color: var(--text-primary);
            margin-top: var(--spacing-xs);
        }
    }
}

.pie-chart-list {
    .pie-item {
        display: flex;
        align-items: center;
        gap: var(--spacing-sm);
        margin-bottom: var(--spacing-sm);

        .pie-color {
            width: 16px;
            height: 16px;
            border-radius: 50%;
        }

        .pie-label {
            flex: 1;
            color: var(--text-primary);
        }

        .pie-count {
            font-weight: var(--font-medium);
            color: var(--text-primary);
        }
    }
}

.details-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: var(--spacing-lg);

    .detail-item {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-xs);

        label {
            font-size: var(--font-sm);
            color: var(--text-secondary);
            font-weight: var(--font-medium);
        }

        value {
            color: var(--text-primary);
            font-weight: var(--font-medium);
        }
    }
}

.issues-table {
    .table-header {
        display: grid;
        grid-template-columns: 100px 1fr 120px 120px 150px;
        gap: var(--spacing-md);
        padding: var(--spacing-md);
        background-color: var(--border-light);
        border-radius: var(--radius-md) var(--radius-md) 0 0;
        font-weight: var(--font-semibold);
        color: var(--text-primary);
        font-size: var(--font-sm);
    }

    .table-row {
        display: grid;
        grid-template-columns: 100px 1fr 120px 120px 150px;
        gap: var(--spacing-md);
        padding: var(--spacing-md);
        border-bottom: 1px solid var(--border-color);
        align-items: center;

        &:hover {
            background-color: var(--hover-bg);
        }

        .issue-key {
            font-weight: var(--font-medium);
            color: var(--primary-color);
        }

        .issue-summary {
            @include truncate-text;
        }

        .issue-type {
            display: flex;
            align-items: center;
            gap: var(--spacing-xs);
            font-size: var(--font-sm);

            .type-icon {
                width: 16px;
                height: 16px;
            }
        }

        .status-badge {
            display: inline-block;
            padding: var(--spacing-xs) var(--spacing-sm);
            border-radius: var(--radius-md);
            font-size: var(--font-xs);
            font-weight: var(--font-medium);
            text-transform: uppercase;

            &.status-todo {
                background-color: var(--border-light);
                color: var(--text-secondary);
            }

            &.status-progress {
                background-color: var(--info-light);
                color: var(--info-color);
            }

            &.status-done {
                background-color: var(--success-light);
                color: var(--success-color);
            }

            &.status-blocked {
                background-color: var(--error-light);
                color: var(--error-color);
            }

            &.status-review {
                background-color: var(--warning-light);
                color: var(--warning-color);
            }

            &.status-default {
                background-color: var(--border-light);
                color: var(--text-secondary);
            }
        }

        .issue-assignee {
            @include truncate-text;
            font-size: var(--font-sm);
        }
    }
}

.no-issues {
    text-align: center;
    padding: var(--spacing-xl);
    color: var(--text-secondary);
}

.update-info {
    text-align: center;
    margin-top: var(--spacing-xl);
}

@keyframes spin {
    to {
        transform: rotate(360deg);
    }
}

@media (max-width: 768px) {
    .epic-header {
        flex-direction: column;
        align-items: flex-start;

        .epic-actions {
            width: 100%;

            .btn {
                flex: 1;
            }
        }
    }

    .stats-overview {
        grid-template-columns: repeat(2, 1fr);
    }

    .chart-grid {
        grid-template-columns: 1fr;
    }

    .details-grid {
        grid-template-columns: 1fr;
    }

    .issues-table {
        .table-header,
        .table-row {
            grid-template-columns: 1fr;
            gap: var(--spacing-xs);
        }

        .table-header {
            display: none;
        }

        .table-row {
            border: 1px solid var(--border-color);
            border-radius: var(--radius-md);
            margin-bottom: var(--spacing-sm);

            span {
                padding: var(--spacing-xs) 0;

                &::before {
                    content: attr(data-label) ": ";
                    font-weight: var(--font-semibold);
                    color: var(--text-secondary);
                }
            }
        }
    }
}
</style>
