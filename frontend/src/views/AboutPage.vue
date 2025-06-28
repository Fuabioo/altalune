<template>
    <div class="about-page">
        <div class="container">
            <!-- Header -->
            <div class="page-header">
                <h1>About Altalune</h1>
                <p>
                    A powerful tool for analyzing and visualizing Jira epic data
                    with beautiful charts and insights.
                </p>
            </div>

            <!-- Application Overview -->
            <div class="overview-section">
                <div class="card">
                    <h2>What is Altalune?</h2>
                    <p>
                        Altalune is a web application designed to help teams
                        track and analyze their Jira epics with comprehensive
                        statistics, visual charts, and detailed insights. It
                        provides an intuitive dashboard to monitor epic
                        progress, issue distributions, and team performance
                        metrics.
                    </p>

                    <div class="feature-grid">
                        <div class="feature-item">
                            <div class="feature-icon">📊</div>
                            <h4>Visual Analytics</h4>
                            <p>
                                Beautiful charts and graphs to visualize epic
                                progress and issue distributions.
                            </p>
                        </div>

                        <div class="feature-item">
                            <div class="feature-icon">⚡</div>
                            <h4>Real-time Data</h4>
                            <p>
                                Connect directly to your Jira instance for
                                up-to-date information.
                            </p>
                        </div>

                        <div class="feature-item">
                            <div class="feature-icon">🎯</div>
                            <h4>Epic Tracking</h4>
                            <p>
                                Track multiple epics and compare their progress
                                across projects.
                            </p>
                        </div>

                        <div class="feature-item">
                            <div class="feature-icon">🎨</div>
                            <h4>Catppuccin Themes</h4>
                            <p>
                                Beautiful Latte (light) and Mocha (dark) themes
                                that adapt to your system preferences.
                            </p>
                        </div>

                        <div class="feature-item">
                            <div class="feature-icon">📱</div>
                            <h4>Responsive Design</h4>
                            <p>
                                Works seamlessly on desktop, tablet, and mobile
                                devices.
                            </p>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Theme Information -->
            <div class="theme-section">
                <div class="card">
                    <h2>🎨 Catppuccin Color Palette</h2>
                    <p>
                        This application features the beloved Catppuccin color
                        scheme, providing a soothing and aesthetically pleasing
                        experience.
                    </p>

                    <div class="theme-showcase">
                        <div class="theme-variant">
                            <h3>☀️ Latte Theme (Light)</h3>
                            <p>
                                A warm, pastel theme perfect for daytime use
                                with excellent readability.
                            </p>
                            <div class="color-preview latte">
                                <span
                                    class="color-dot"
                                    style="background-color: #1e66f5"
                                ></span>
                                <span
                                    class="color-dot"
                                    style="background-color: #40a02b"
                                ></span>
                                <span
                                    class="color-dot"
                                    style="background-color: #df8e1d"
                                ></span>
                                <span
                                    class="color-dot"
                                    style="background-color: #d20f39"
                                ></span>
                                <span
                                    class="color-dot"
                                    style="background-color: #8839ef"
                                ></span>
                            </div>
                        </div>

                        <div class="theme-variant">
                            <h3>🌙 Mocha Theme (Dark)</h3>
                            <p>
                                A cozy, dark theme that's easy on the eyes for
                                extended coding sessions.
                            </p>
                            <div class="color-preview mocha">
                                <span
                                    class="color-dot"
                                    style="background-color: #89b4fa"
                                ></span>
                                <span
                                    class="color-dot"
                                    style="background-color: #a6e3a1"
                                ></span>
                                <span
                                    class="color-dot"
                                    style="background-color: #f9e2af"
                                ></span>
                                <span
                                    class="color-dot"
                                    style="background-color: #f38ba8"
                                ></span>
                                <span
                                    class="color-dot"
                                    style="background-color: #cba6f7"
                                ></span>
                            </div>
                        </div>
                    </div>

                    <div class="theme-note">
                        <p>
                            <strong>Automatic Theme Detection:</strong> The
                            application automatically switches between themes
                            based on your system's dark/light mode preference.
                        </p>
                    </div>
                </div>
            </div>

            <!-- Technology Stack -->
            <div class="tech-section">
                <div class="card">
                    <h2>Built With</h2>
                    <p>
                        This application is built using modern web technologies
                        for optimal performance and user experience:
                    </p>

                    <div class="tech-stack">
                        <div class="tech-category">
                            <h3>Frontend</h3>
                            <div class="tech-items">
                                <div class="tech-item">
                                    <span class="tech-icon">⚡</span>
                                    <div class="tech-info">
                                        <strong>Vite</strong>
                                        <p>
                                            Lightning-fast build tool and
                                            development server
                                        </p>
                                    </div>
                                </div>

                                <div class="tech-item">
                                    <span class="tech-icon">💚</span>
                                    <div class="tech-info">
                                        <strong>Vue 3</strong>
                                        <p>
                                            Progressive JavaScript framework
                                            with Composition API
                                        </p>
                                    </div>
                                </div>

                                <div class="tech-item">
                                    <span class="tech-icon">🎨</span>
                                    <div class="tech-info">
                                        <strong>SCSS</strong>
                                        <p>
                                            Enhanced CSS with variables, mixins,
                                            and nesting
                                        </p>
                                    </div>
                                </div>

                                <div class="tech-item">
                                    <span class="tech-icon">🧭</span>
                                    <div class="tech-info">
                                        <strong>Vue Router</strong>
                                        <p>
                                            Official router for Vue.js
                                            applications
                                        </p>
                                    </div>
                                </div>

                                <div class="tech-item">
                                    <span class="tech-icon">🍍</span>
                                    <div class="tech-info">
                                        <strong>Pinia</strong>
                                        <p>
                                            Modern state management for Vue.js
                                        </p>
                                    </div>
                                </div>

                                <div class="tech-item">
                                    <span class="tech-icon">🔧</span>
                                    <div class="tech-info">
                                        <strong>VueUse</strong>
                                        <p>
                                            Collection of Vue composition
                                            utilities
                                        </p>
                                    </div>
                                </div>
                            </div>
                        </div>

                        <div class="tech-category">
                            <h3>Backend</h3>
                            <div class="tech-items">
                                <div class="tech-item">
                                    <span class="tech-icon">🐹</span>
                                    <div class="tech-info">
                                        <strong>Go</strong>
                                        <p>Fast and efficient backend server</p>
                                    </div>
                                </div>

                                <div class="tech-item">
                                    <span class="tech-icon">🔗</span>
                                    <div class="tech-info">
                                        <strong>Jira REST API</strong>
                                        <p>
                                            Direct integration with Atlassian
                                            Jira
                                        </p>
                                    </div>
                                </div>
                            </div>
                        </div>

                        <div class="tech-category">
                            <h3>Development</h3>
                            <div class="tech-items">
                                <div class="tech-item">
                                    <span class="tech-icon">🥟</span>
                                    <div class="tech-info">
                                        <strong>Bun</strong>
                                        <p>
                                            Fast JavaScript runtime and package
                                            manager
                                        </p>
                                    </div>
                                </div>

                                <div class="tech-item">
                                    <span class="tech-icon">📦</span>
                                    <div class="tech-info">
                                        <strong>Component Architecture</strong>
                                        <p>
                                            Modular and reusable component
                                            design
                                        </p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Features Detail -->
            <div class="features-section">
                <div class="card">
                    <h2>Key Features</h2>

                    <div class="features-list">
                        <div class="feature-detail">
                            <h3>📊 Epic Analytics</h3>
                            <ul>
                                <li>
                                    Progress tracking with visual progress bars
                                </li>
                                <li>Issue status distribution charts</li>
                                <li>Issue type breakdown and analysis</li>
                                <li>Completion percentage calculations</li>
                            </ul>
                        </div>

                        <div class="feature-detail">
                            <h3>⚙️ Epic Management</h3>
                            <ul>
                                <li>Add, edit, and remove tracked epics</li>
                                <li>Persistent storage using localStorage</li>
                                <li>Epic validation and duplicate detection</li>
                                <li>Bulk operations for multiple epics</li>
                            </ul>
                        </div>

                        <div class="feature-detail">
                            <h3>🔌 Jira Integration</h3>
                            <ul>
                                <li>Secure API token authentication</li>
                                <li>Real-time data fetching from Jira</li>
                                <li>Connection testing and troubleshooting</li>
                                <li>Support for Jira Cloud and Server</li>
                            </ul>
                        </div>

                        <div class="feature-detail">
                            <h3>🎨 User Experience</h3>
                            <ul>
                                <li>Responsive design for all devices</li>
                                <li>Catppuccin Latte/Mocha theme support</li>
                                <li>Automatic dark/light mode detection</li>
                                <li>Intuitive navigation and layout</li>
                                <li>Smooth transitions and animations</li>
                                <li>Loading states and error handling</li>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Getting Started -->
            <div class="getting-started-section">
                <div class="card">
                    <h2>Getting Started</h2>
                    <p>
                        Follow these simple steps to start analyzing your Jira
                        epics:
                    </p>

                    <div class="steps-timeline">
                        <div class="step">
                            <div class="step-number">1</div>
                            <div class="step-content">
                                <h4>Configure Connection</h4>
                                <p>
                                    Set up your Jira API credentials in the
                                    <router-link to="/admin/setup"
                                        >Setup page</router-link
                                    >.
                                </p>
                            </div>
                        </div>

                        <div class="step">
                            <div class="step-number">2</div>
                            <div class="step-content">
                                <h4>Add Epics</h4>
                                <p>
                                    Use the
                                    <router-link to="/admin/epics"
                                        >Manage Epics</router-link
                                    >
                                    page to add epics you want to track.
                                </p>
                            </div>
                        </div>

                        <div class="step">
                            <div class="step-number">3</div>
                            <div class="step-content">
                                <h4>View Analytics</h4>
                                <p>
                                    Navigate to any epic page to see detailed
                                    analytics and visualizations.
                                </p>
                            </div>
                        </div>

                        <div class="step">
                            <div class="step-number">4</div>
                            <div class="step-content">
                                <h4>Track Progress</h4>
                                <p>
                                    Monitor your epics' progress and share
                                    insights with your team.
                                </p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Version & Credits -->
            <div class="credits-section">
                <div class="card">
                    <div class="credits-content">
                        <div class="version-info">
                            <h3>Version Information</h3>
                            <p><strong>Version:</strong> 1.0.0</p>
                            <p>
                                <strong>Release Date:</strong> {{ currentDate }}
                            </p>
                            <p><strong>License:</strong> MIT License</p>
                        </div>

                        <div class="support-info">
                            <h3>Support & Documentation</h3>
                            <p>
                                For support, feature requests, or bug reports:
                            </p>
                            <div class="support-links">
                                <a href="#" class="support-link"
                                    >📖 Documentation</a
                                >
                                <a href="#" class="support-link"
                                    >🐛 Report Issues</a
                                >
                                <a href="#" class="support-link"
                                    >💡 Feature Requests</a
                                >
                                <a href="#" class="support-link"
                                    >💬 Community</a
                                >
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { computed } from "vue";

export default {
    name: "AboutPage",
    setup() {
        const currentDate = computed(() => {
            return new Date().toLocaleDateString("en-US", {
                year: "numeric",
                month: "long",
                day: "numeric",
            });
        });

        return {
            currentDate,
        };
    },
};
</script>

<style lang="scss" scoped>
.about-page {
    padding: var(--spacing-xl) 0;
    min-height: 100vh;
}

.page-header {
    text-align: center;
    margin-bottom: var(--spacing-xxl);

    h1 {
        margin-bottom: var(--spacing-sm);
        background: linear-gradient(
            135deg,
            var(--primary-color),
            var(--primary-hover)
        );
        -webkit-background-clip: text;
        -webkit-text-fill-color: transparent;
        background-clip: text;
    }

    p {
        color: var(--text-secondary);
        font-size: var(--font-lg);
        max-width: 600px;
        margin: 0 auto;
    }
}

.overview-section,
.tech-section,
.features-section,
.getting-started-section,
.credits-section {
    margin-bottom: var(--spacing-xxl);

    .card {
        max-width: 900px;
        margin: 0 auto;
    }

    h2 {
        margin-bottom: var(--spacing-lg);
        color: var(--text-primary);
    }
}

.feature-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: var(--spacing-lg);
    margin-top: var(--spacing-lg);

    .feature-item {
        text-align: center;
        padding: var(--spacing-lg);

        .feature-icon {
            font-size: 2.5rem;
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
}

.tech-stack {
    margin-top: var(--spacing-lg);

    .tech-category {
        margin-bottom: var(--spacing-xl);

        h3 {
            color: var(--primary-color);
            margin-bottom: var(--spacing-md);
            padding-bottom: var(--spacing-xs);
            border-bottom: 2px solid var(--primary-light);
        }

        .tech-items {
            display: grid;
            gap: var(--spacing-md);

            .tech-item {
                display: flex;
                align-items: flex-start;
                gap: var(--spacing-md);
                padding: var(--spacing-md);
                background-color: var(--border-light);
                border-radius: var(--radius-md);

                .tech-icon {
                    font-size: 1.5rem;
                    flex-shrink: 0;
                }

                .tech-info {
                    strong {
                        color: var(--text-primary);
                        display: block;
                        margin-bottom: var(--spacing-xs);
                    }

                    p {
                        color: var(--text-secondary);
                        font-size: var(--font-sm);
                        margin: 0;
                    }
                }
            }
        }
    }
}

.features-list {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: var(--spacing-lg);
    margin-top: var(--spacing-lg);

    .feature-detail {
        h3 {
            margin-bottom: var(--spacing-md);
            color: var(--text-primary);
        }

        ul {
            list-style: none;
            padding: 0;

            li {
                position: relative;
                padding-left: var(--spacing-lg);
                margin-bottom: var(--spacing-sm);
                color: var(--text-secondary);

                &::before {
                    content: "✓";
                    position: absolute;
                    left: 0;
                    color: var(--success-color);
                    font-weight: var(--font-bold);
                }
            }
        }
    }
}

.steps-timeline {
    margin-top: var(--spacing-lg);

    .step {
        display: flex;
        gap: var(--spacing-lg);
        margin-bottom: var(--spacing-xl);
        position: relative;

        &:not(:last-child)::after {
            content: "";
            position: absolute;
            left: 19px;
            top: 40px;
            bottom: -24px;
            width: 2px;
            background-color: var(--border-color);
        }

        .step-number {
            width: 40px;
            height: 40px;
            border-radius: 50%;
            background-color: var(--primary-color);
            color: var(--text-inverse);
            @include flex-center;
            font-weight: var(--font-bold);
            flex-shrink: 0;
            position: relative;
            z-index: 1;
        }

        .step-content {
            flex: 1;

            h4 {
                margin-bottom: var(--spacing-sm);
                color: var(--text-primary);
            }

            p {
                color: var(--text-secondary);
                margin: 0;

                a {
                    color: var(--primary-color);
                    text-decoration: none;

                    &:hover {
                        text-decoration: underline;
                    }
                }
            }
        }
    }
}

.credits-content {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: var(--spacing-xl);

    .version-info,
    .support-info {
        h3 {
            margin-bottom: var(--spacing-md);
            color: var(--text-primary);
        }

        p {
            margin-bottom: var(--spacing-sm);
            color: var(--text-secondary);

            strong {
                color: var(--text-primary);
            }
        }
    }

    .support-links {
        display: flex;
        flex-direction: column;
        gap: var(--spacing-sm);
        margin-top: var(--spacing-md);

        .support-link {
            display: inline-flex;
            align-items: center;
            gap: var(--spacing-xs);
            color: var(--primary-color);
            text-decoration: none;
            font-size: var(--font-sm);

            &:hover {
                text-decoration: underline;
            }
        }
    }
}

.theme-section {
    .theme-showcase {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
        gap: var(--spacing-lg);
        margin: var(--spacing-lg) 0;
    }

    .theme-variant {
        padding: var(--spacing-md);
        background-color: var(--border-light);
        border-radius: var(--radius-md);
        text-align: center;

        h3 {
            margin-bottom: var(--spacing-sm);
            color: var(--text-primary);
        }

        p {
            margin-bottom: var(--spacing-md);
            color: var(--text-secondary);
            font-size: var(--font-sm);
        }
    }

    .color-preview {
        display: flex;
        justify-content: center;
        gap: var(--spacing-sm);

        .color-dot {
            width: 24px;
            height: 24px;
            border-radius: 50%;
            border: 2px solid var(--border-color);
            transition: transform 0.2s ease;

            &:hover {
                transform: scale(1.2);
            }
        }
    }

    .theme-note {
        margin-top: var(--spacing-lg);
        padding: var(--spacing-md);
        background-color: var(--info-light);
        border: 1px solid var(--info-color);
        border-radius: var(--radius-md);
        text-align: center;

        p {
            margin: 0;
            color: var(--info-color);
            font-size: var(--font-sm);
        }
    }
}

@media (max-width: 768px) {
    .feature-grid {
        grid-template-columns: 1fr;
    }

    .tech-stack .tech-category .tech-items {
        grid-template-columns: 1fr;
    }

    .features-list {
        grid-template-columns: 1fr;
    }

    .credits-content {
        grid-template-columns: 1fr;
    }

    .support-links {
        flex-direction: column;
    }

    .theme-section .theme-showcase {
        grid-template-columns: 1fr;
    }
}
</style>
