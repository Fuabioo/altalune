<template>
    <div id="app">
        <div class="app-layout">
            <!-- Sidebar -->
            <aside
                class="sidebar"
                :class="{ collapsed: uiStore.sidebarCollapsed }"
            >
                <div class="sidebar-header">
                    <h2 v-show="!uiStore.sidebarCollapsed">Altalune</h2>
                    <button
                        class="sidebar-toggle"
                        @click="uiStore.toggleSidebar"
                        :title="
                            uiStore.sidebarCollapsed
                                ? 'Expand sidebar'
                                : 'Collapse sidebar'
                        "
                    >
                        <span class="toggle-icon">{{
                            uiStore.sidebarCollapsed ? "▶" : "◀"
                        }}</span>
                    </button>
                </div>

                <nav class="sidebar-nav">
                    <div class="nav-content">
                        <!-- Home -->
                        <router-link
                            to="/"
                            class="nav-item"
                            exact-active-class="active"
                        >
                            <span class="nav-icon">🏠</span>
                            <span class="nav-text">Home</span>
                        </router-link>

                        <!-- Admin Group -->
                        <div class="nav-group">
                            <div class="nav-group-title">Admin</div>
                            <router-link
                                to="/admin/epics"
                                class="nav-item"
                                active-class="active"
                            >
                                <span class="nav-icon">📋</span>
                                <span class="nav-text">Manage Epics</span>
                            </router-link>
                            <router-link
                                to="/admin/setup"
                                class="nav-item"
                                active-class="active"
                            >
                                <span class="nav-icon">⚙️</span>
                                <span class="nav-text">Setup</span>
                            </router-link>
                        </div>

                        <!-- Divider -->
                        <div class="nav-divider"></div>

                        <!-- Dynamic Epic Routes -->
                        <div
                            class="nav-group"
                            v-if="epicsStore.epics.length > 0"
                        >
                            <div class="nav-group-title">Saved Epics</div>
                            <router-link
                                v-for="epic in epicsStore.epics"
                                :key="epic.id"
                                :to="`/epic/${epic.epicCode}`"
                                class="nav-item epic-item"
                                active-class="active"
                            >
                                <span
                                    class="nav-icon"
                                    v-show="uiStore.sidebarCollapsed"
                                    >📊</span
                                >
                                <span class="nav-text">{{
                                    epic.epicCode
                                }}</span>
                                <span class="epic-title">{{ epic.title }}</span>

                                <!-- Tooltip for collapsed state -->
                                <div
                                    v-if="uiStore.sidebarCollapsed"
                                    class="epic-tooltip"
                                >
                                    <div class="tooltip-content">
                                        <div class="tooltip-code">
                                            {{ epic.epicCode }}
                                        </div>
                                        <div class="tooltip-title">
                                            {{ epic.title }}
                                        </div>
                                    </div>
                                </div>
                            </router-link>
                        </div>
                    </div>

                    <!-- About at bottom -->
                    <div class="nav-footer">
                        <div class="nav-divider"></div>
                        <router-link
                            to="/about"
                            class="nav-item"
                            active-class="active"
                        >
                            <span class="nav-icon">ℹ️</span>
                            <span class="nav-text">About</span>
                        </router-link>
                    </div>
                </nav>
            </aside>

            <!-- Main Content -->
            <main class="main-content">
                <router-view />
            </main>
        </div>
    </div>
</template>

<script>
import { useEpicsStore } from "@/stores/epics";
import { useUIStore } from "@/stores/ui";

export default {
    name: "App",
    setup() {
        const epicsStore = useEpicsStore();
        const uiStore = useUIStore();

        return {
            epicsStore,
            uiStore,
            isCollapsed: uiStore.sidebarCollapsed,
            toggleSidebar: uiStore.toggleSidebar,
        };
    },
};
</script>

<style lang="scss" scoped>
.app-layout {
    display: flex;
    height: 100vh;
    background-color: var(--bg-color);
}

.sidebar {
    width: 240px;
    background-color: var(--sidebar-bg);
    border-right: 1px solid var(--border-color);
    display: flex;
    flex-direction: column;
    overflow-y: auto;
    overflow-x: hidden;
    transition: width 0.3s ease;

    &.collapsed {
        width: 56px;

        .nav-item {
            .nav-text,
            .epic-title {
                display: none;
            }

            .nav-icon {
                margin-right: 0;
            }

            justify-content: center;
        }

        .nav-group-title {
            display: none;
        }

        .nav-divider {
            margin: 0.375rem 0.5rem;
        }
    }

    .sidebar-header {
        padding: 1rem 0.75rem;
        border-bottom: 1px solid var(--border-color);
        background-color: var(--sidebar-header-bg);
        display: flex;
        align-items: center;
        justify-content: space-between;

        h2 {
            margin: 0;
            color: var(--text-primary);
            font-size: 1.1rem;
            font-weight: 600;
        }

        .sidebar-toggle {
            background: none;
            border: none;
            color: var(--text-secondary);
            cursor: pointer;
            padding: 0.375rem;
            border-radius: 4px;
            transition: all 0.2s ease;
            display: flex;
            align-items: center;
            justify-content: center;
            min-width: 28px;
            height: 28px;

            &:hover {
                background-color: var(--hover-bg);
                color: var(--text-primary);
            }

            .toggle-icon {
                font-size: 0.875rem;
                line-height: 1;
            }
        }
    }

    &.collapsed .sidebar-header {
        padding: 1rem 0.375rem;
        justify-content: center;

        .sidebar-toggle {
            margin: 0;
        }
    }

    .sidebar-nav {
        flex: 1;
        padding: 0.75rem 0;
        display: flex;
        flex-direction: column;

        .nav-content {
            flex: 1;
        }

        .nav-footer {
            margin-top: auto;
        }
    }

    .nav-item {
        display: flex;
        align-items: center;
        padding: 0.625rem 0.75rem;
        color: var(--text-secondary);
        text-decoration: none;
        transition: all 0.2s ease;
        border-left: 3px solid transparent;
        min-width: 0;
        overflow: hidden;

        &:hover {
            background-color: var(--hover-bg);
            color: var(--text-primary);
        }

        &.active {
            background-color: var(--active-bg);
            color: var(--primary-color);
            border-left-color: var(--primary-color);
        }

        .nav-icon {
            margin-right: 0.625rem;
            font-size: 1rem;
            min-width: 16px;
        }

        .nav-text {
            font-weight: 500;
            font-size: 0.9rem;
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;
            flex: 1;
        }

        &.epic-item {
            flex-direction: row;
            align-items: flex-start;
            padding: 0.625rem 0.75rem;
            min-width: 0;
            flex-wrap: wrap;

            .nav-icon {
                margin-top: 0.125rem;
            }

            .nav-text {
                font-weight: 600;
                font-size: 0.875rem;
                line-height: 1.2;
            }

            .epic-title {
                font-size: 0.75rem;
                color: var(--text-muted);
                margin-top: 0.125rem;
                white-space: nowrap;
                overflow: hidden;
                text-overflow: ellipsis;
                width: 100%;
                max-width: 100%;
                flex-basis: 100%;
                line-height: 1.2;
            }
        }

        &.epic-item:not(.collapsed) {
            .epic-title {
                margin-left: 0;
                width: 100%;
                max-width: 100%;
            }
        }

        &.epic-item {
            position: relative;

            .epic-tooltip {
                position: absolute;
                left: calc(100% + 8px);
                top: 50%;
                transform: translateY(-50%);
                background-color: var(--card-bg);
                border: 1px solid var(--border-color);
                border-radius: 6px;
                padding: 0.5rem 0.75rem;
                box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
                opacity: 0;
                visibility: hidden;
                transition:
                    opacity 0.2s ease,
                    visibility 0.2s ease;
                z-index: 1000;
                min-width: 180px;
                max-width: 280px;
                white-space: nowrap;

                .tooltip-content {
                    .tooltip-code {
                        font-weight: 600;
                        font-size: 0.875rem;
                        color: var(--text-primary);
                        margin-bottom: 0.25rem;
                    }

                    .tooltip-title {
                        font-size: 0.75rem;
                        color: var(--text-secondary);
                        line-height: 1.3;
                        white-space: normal;
                        word-wrap: break-word;
                    }
                }

                /* Arrow pointing to the nav item */
                &::before {
                    content: "";
                    position: absolute;
                    left: -6px;
                    top: 50%;
                    transform: translateY(-50%);
                    width: 0;
                    height: 0;
                    border-top: 6px solid transparent;
                    border-bottom: 6px solid transparent;
                    border-right: 6px solid var(--card-bg);
                }

                &::after {
                    content: "";
                    position: absolute;
                    left: -7px;
                    top: 50%;
                    transform: translateY(-50%);
                    width: 0;
                    height: 0;
                    border-top: 7px solid transparent;
                    border-bottom: 7px solid transparent;
                    border-right: 7px solid var(--border-color);
                }
            }

            &:hover .epic-tooltip {
                opacity: 1;
                visibility: visible;
            }
        }
    }

    .nav-group {
        margin: 0.375rem 0;

        .nav-group-title {
            font-size: 0.75rem;
            font-weight: 600;
            color: var(--text-muted);
            text-transform: uppercase;
            letter-spacing: 0.5px;
            padding: 0.375rem 0.75rem;
            margin-bottom: 0.25rem;
        }
    }

    .nav-divider {
        height: 1px;
        background-color: var(--border-color);
        margin: 0.75rem 0;
    }
}

.main-content {
    flex: 1;
    overflow-y: auto;
    background-color: var(--bg-color);
}

@media (max-width: 768px) {
    .app-layout {
        flex-direction: column;
    }

    .sidebar {
        width: 100% !important;
        height: auto;
        order: 2;

        &.collapsed {
            width: 100% !important;

            .nav-item {
                .nav-text,
                .epic-title {
                    display: block;
                }

                .nav-icon {
                    margin-right: 0.625rem;
                }

                justify-content: flex-start;
            }

            .nav-group-title {
                display: block;
            }

            .nav-divider {
                margin: 0.75rem 0;
            }
        }

        .sidebar-header {
            padding: 1rem 0.75rem !important;
            justify-content: space-between !important;
        }
    }

    .main-content {
        order: 1;
    }
}
</style>
