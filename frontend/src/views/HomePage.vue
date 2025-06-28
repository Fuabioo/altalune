<template>
    <div class="home-page">
        <div class="container">
            <!-- Header Section -->
            <div class="hero-section">
                <h1>Welcome to Altalune</h1>
                <p class="hero-subtitle">
                    Analyze and visualize your Jira epics with detailed
                    statistics and beautiful graphs.
                </p>
            </div>

            <!-- Quick Stats -->
            <div class="stats-grid">
                <div class="stat-card">
                    <div class="stat-icon">📊</div>
                    <div class="stat-content">
                        <h3>{{ epicsStore.epicCount }}</h3>
                        <p>Saved Epics</p>
                    </div>
                </div>

                <div class="stat-card">
                    <div class="stat-icon">⚙️</div>
                    <div class="stat-content">
                        <h3>{{ connectionStatus }}</h3>
                        <p>Connection Status</p>
                    </div>
                </div>

                <div class="stat-card">
                    <div class="stat-icon">🕒</div>
                    <div class="stat-content">
                        <h3>{{ lastUpdated }}</h3>
                        <p>Last Updated</p>
                    </div>
                </div>
            </div>

            <!-- Recent Epics -->
            <div class="recent-epics" v-if="epicsStore.epics.length > 0">
                <h2>Recent Epics</h2>
                <div class="epics-grid">
                    <router-link
                        v-for="epic in recentEpics"
                        :key="epic.id"
                        :to="`/epic/${epic.epicCode}`"
                        class="epic-card"
                    >
                        <div class="epic-header">
                            <span class="epic-code">{{ epic.epicCode }}</span>
                            <span class="epic-date">{{
                                formatDate(epic.createdAt)
                            }}</span>
                        </div>
                        <h4 class="epic-title">{{ epic.title }}</h4>
                        <p class="epic-description">{{ epic.description }}</p>
                    </router-link>
                </div>
            </div>

            <!-- Getting Started -->
            <div class="getting-started" v-else>
                <h2>Getting Started</h2>
                <div class="steps-grid">
                    <div class="step-card">
                        <div class="step-number">1</div>
                        <div class="step-content">
                            <h4>Configure Setup</h4>
                            <p>
                                Configure your Jira connection in the setup
                                page.
                            </p>
                            <router-link
                                to="/admin/setup"
                                class="btn btn-secondary btn-sm"
                            >
                                Go to Setup
                            </router-link>
                        </div>
                    </div>

                    <div class="step-card">
                        <div class="step-number">2</div>
                        <div class="step-content">
                            <h4>Add Your First Epic</h4>
                            <p>Add epics you want to analyze and track.</p>
                            <router-link
                                to="/admin/epics"
                                class="btn btn-primary btn-sm"
                            >
                                Manage Epics
                            </router-link>
                        </div>
                    </div>

                    <div class="step-card">
                        <div class="step-number">3</div>
                        <div class="step-content">
                            <h4>Analyze Data</h4>
                            <p>View detailed analytics for your epics.</p>
                            <span
                                class="btn btn-secondary btn-sm"
                                style="opacity: 0.6"
                            >
                                Coming Soon
                            </span>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Theme Showcase -->
            <div class="theme-showcase">
                <h2>Catppuccin Theme</h2>
                <p>
                    This application uses the beautiful Catppuccin color palette
                    that adapts to your system preferences.
                </p>

                <div class="theme-demo">
                    <div class="color-palette">
                        <div class="color-row">
                            <div class="color-item primary">
                                <div class="color-swatch"></div>
                                <span>Primary</span>
                            </div>
                            <div class="color-item success">
                                <div class="color-swatch"></div>
                                <span>Success</span>
                            </div>
                            <div class="color-item warning">
                                <div class="color-swatch"></div>
                                <span>Warning</span>
                            </div>
                            <div class="color-item error">
                                <div class="color-swatch"></div>
                                <span>Error</span>
                            </div>
                        </div>
                    </div>

                    <div class="theme-info">
                        <div class="current-theme">
                            <h4>Current Theme</h4>
                            <div class="theme-badge" :class="currentTheme">
                                {{
                                    currentTheme === "light"
                                        ? "☀️ Latte (Light)"
                                        : "🌙 Mocha (Dark)"
                                }}
                            </div>
                        </div>
                        <p class="theme-description">
                            {{
                                currentTheme === "light"
                                    ? "Using Catppuccin Latte - a soothing pastel theme for daytime use"
                                    : "Using Catppuccin Mocha - a cozy dark theme perfect for nighttime coding"
                            }}
                        </p>
                    </div>
                </div>
            </div>

            <!-- Quick Actions -->
            <div class="quick-actions">
                <h2>Quick Actions</h2>
                <div class="actions-grid">
                    <router-link to="/admin/epics" class="action-card">
                        <div class="action-icon">📋</div>
                        <h4>Manage Epics</h4>
                        <p>Add, edit, or remove tracked epics</p>
                    </router-link>

                    <router-link to="/admin/setup" class="action-card">
                        <div class="action-icon">⚙️</div>
                        <h4>Configuration</h4>
                        <p>Set up Jira connection and preferences</p>
                    </router-link>

                    <router-link to="/about" class="action-card">
                        <div class="action-icon">ℹ️</div>
                        <h4>About</h4>
                        <p>Learn more about this application</p>
                    </router-link>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { computed, ref, onMounted } from "vue";
import { useEpicsStore } from "@/stores/epics";

export default {
    name: "HomePage",
    setup() {
        const epicsStore = useEpicsStore();
        const connectionStatus = ref("Unknown");

        // Detect current theme
        const currentTheme = ref("light");

        const updateTheme = () => {
            if (
                window.matchMedia &&
                window.matchMedia("(prefers-color-scheme: dark)").matches
            ) {
                currentTheme.value = "dark";
            } else {
                currentTheme.value = "light";
            }
        };

        // Get recent epics (last 3)
        const recentEpics = computed(() => {
            return epicsStore.epics
                .slice()
                .sort((a, b) => new Date(b.createdAt) - new Date(a.createdAt))
                .slice(0, 3);
        });

        const lastUpdated = computed(() => {
            if (epicsStore.epics.length === 0) return "Never";

            const latest = epicsStore.epics.reduce((latest, epic) => {
                const epicDate = new Date(epic.updatedAt || epic.createdAt);
                return epicDate > latest ? epicDate : latest;
            }, new Date(0));

            return formatDate(latest.toISOString());
        });

        const formatDate = (dateString) => {
            const date = new Date(dateString);
            const now = new Date();
            const diffTime = Math.abs(now - date);
            const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));

            if (diffDays === 1) return "Today";
            if (diffDays === 2) return "Yesterday";
            if (diffDays <= 7) return `${diffDays - 1} days ago`;

            return date.toLocaleDateString();
        };

        const checkConnection = async () => {
            try {
                const response = await fetch("/api/ping");
                connectionStatus.value = response.ok
                    ? "Connected"
                    : "Disconnected";
            } catch (error) {
                connectionStatus.value = "Disconnected";
            }
        };

        onMounted(() => {
            checkConnection();
            updateTheme();

            // Listen for theme changes
            if (window.matchMedia) {
                const mediaQuery = window.matchMedia(
                    "(prefers-color-scheme: dark)",
                );
                mediaQuery.addListener(updateTheme);
            }
        });

        return {
            epicsStore,
            recentEpics,
            connectionStatus,
            lastUpdated,
            currentTheme,
            formatDate,
        };
    },
};
</script>

<style lang="scss" scoped>
.home-page {
    padding: var(--spacing-xl) 0;
    min-height: 100vh;
}

.hero-section {
    text-align: center;
    margin-bottom: var(--spacing-xxl);

    h1 {
        font-size: var(--font-xxxl);
        margin-bottom: var(--spacing-md);
        background: linear-gradient(
            135deg,
            var(--primary-color),
            var(--primary-hover)
        );
        -webkit-background-clip: text;
        -webkit-text-fill-color: transparent;
        background-clip: text;
    }

    .hero-subtitle {
        font-size: var(--font-lg);
        color: var(--text-secondary);
        max-width: 600px;
        margin: 0 auto;
    }
}

.stats-grid {
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

.recent-epics,
.getting-started,
.theme-showcase,
.quick-actions {
    margin-bottom: var(--spacing-xxl);

    h2 {
        margin-bottom: var(--spacing-lg);
        color: var(--text-primary);
    }
}

.epics-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: var(--spacing-lg);
}

.epic-card {
    @include card-style;
    padding: var(--spacing-lg);
    text-decoration: none;
    color: inherit;
    transition: var(--transition-normal);

    &:hover {
        transform: translateY(-2px);
        box-shadow: var(--shadow-md);
        text-decoration: none;
    }

    .epic-header {
        @include flex-between;
        margin-bottom: var(--spacing-sm);

        .epic-code {
            background-color: var(--primary-light);
            color: var(--primary-color);
            padding: var(--spacing-xs) var(--spacing-sm);
            border-radius: var(--radius-md);
            font-weight: var(--font-medium);
            font-size: var(--font-sm);
        }

        .epic-date {
            color: var(--text-muted);
            font-size: var(--font-sm);
        }
    }

    .epic-title {
        margin-bottom: var(--spacing-sm);
        color: var(--text-primary);
    }

    .epic-description {
        color: var(--text-secondary);
        font-size: var(--font-sm);
        margin: 0;
        display: -webkit-box;
        -webkit-line-clamp: 2;
        -webkit-box-orient: vertical;
        overflow: hidden;
    }
}

.steps-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: var(--spacing-lg);
}

.step-card {
    @include card-style;
    padding: var(--spacing-lg);
    text-align: center;

    .step-number {
        width: 40px;
        height: 40px;
        border-radius: 50%;
        background-color: var(--primary-color);
        color: var(--text-inverse);
        @include flex-center;
        font-weight: var(--font-bold);
        margin: 0 auto var(--spacing-md);
    }

    .step-content {
        h4 {
            margin-bottom: var(--spacing-sm);
            color: var(--text-primary);
        }

        p {
            margin-bottom: var(--spacing-md);
            color: var(--text-secondary);
            font-size: var(--font-sm);
        }
    }
}

.actions-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: var(--spacing-lg);
}

.action-card {
    @include card-style;
    padding: var(--spacing-lg);
    text-align: center;
    text-decoration: none;
    color: inherit;
    transition: var(--transition-normal);

    &:hover {
        transform: translateY(-2px);
        box-shadow: var(--shadow-md);
        text-decoration: none;
    }

    .action-icon {
        font-size: 2rem;
        margin-bottom: var(--spacing-md);
    }

    h4 {
        margin-bottom: var(--spacing-sm);
        color: var(--text-primary);
    }

    p {
        color: var(--text-secondary);
        font-size: var(--font-sm);
        margin: 0;
    }
}

.theme-showcase {
    .theme-demo {
        @include card-style;
        padding: var(--spacing-lg);
        margin-top: var(--spacing-lg);
    }

    .color-palette {
        margin-bottom: var(--spacing-lg);

        .color-row {
            display: flex;
            gap: var(--spacing-lg);
            justify-content: center;
            flex-wrap: wrap;
        }

        .color-item {
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: var(--spacing-sm);

            .color-swatch {
                width: 40px;
                height: 40px;
                border-radius: var(--radius-lg);
                border: 2px solid var(--border-color);
                transition: var(--transition-normal);
            }

            span {
                font-size: var(--font-sm);
                color: var(--text-secondary);
                font-weight: var(--font-medium);
            }

            &.primary .color-swatch {
                background-color: var(--primary-color);
            }

            &.success .color-swatch {
                background-color: var(--success-color);
            }

            &.warning .color-swatch {
                background-color: var(--warning-color);
            }

            &.error .color-swatch {
                background-color: var(--error-color);
            }
        }
    }

    .theme-info {
        text-align: center;

        .current-theme {
            margin-bottom: var(--spacing-md);

            h4 {
                margin-bottom: var(--spacing-sm);
                color: var(--text-primary);
            }

            .theme-badge {
                display: inline-flex;
                align-items: center;
                gap: var(--spacing-xs);
                padding: var(--spacing-sm) var(--spacing-md);
                border-radius: var(--radius-lg);
                font-weight: var(--font-medium);
                font-size: var(--font-sm);

                &.light {
                    background-color: var(--warning-light);
                    color: var(--warning-color);
                    border: 1px solid var(--warning-color);
                }

                &.dark {
                    background-color: var(--primary-light);
                    color: var(--primary-color);
                    border: 1px solid var(--primary-color);
                }
            }
        }

        .theme-description {
            color: var(--text-secondary);
            font-style: italic;
            margin: 0;
        }
    }
}

@media (max-width: 768px) {
    .hero-section {
        h1 {
            font-size: var(--font-xxl);
        }

        .hero-subtitle {
            font-size: var(--font-md);
        }
    }

    .stats-grid,
    .epics-grid,
    .steps-grid,
    .actions-grid {
        grid-template-columns: 1fr;
    }

    .theme-showcase {
        .color-palette .color-row {
            gap: var(--spacing-md);
        }

        .color-item .color-swatch {
            width: 32px;
            height: 32px;
        }
    }
}
</style>
