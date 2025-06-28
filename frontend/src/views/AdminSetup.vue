<template>
    <div class="admin-setup-page">
        <div class="container">
            <!-- Header -->
            <div class="page-header">
                <h1>Setup Configuration</h1>
                <p>
                    Configure your Jira connection and test the API
                    connectivity.
                </p>
            </div>

            <!-- Configuration Instructions -->
            <div class="setup-section">
                <div class="card">
                    <h2>Environment Configuration</h2>
                    <p>
                        To connect to your Jira instance, you need to configure
                        the following environment variables on your backend
                        server:
                    </p>

                    <div class="env-variables">
                        <div class="env-var">
                            <label>JIRA_EPIC_WORKSPACE</label>
                            <code>your-company.atlassian.net</code>
                            <p>The base URL of your Jira instance</p>
                        </div>

                        <div class="env-var">
                            <label>JIRA_EPIC_EMAIL</label>
                            <code>your-email@company.com</code>
                            <p>Your Jira account email address</p>
                        </div>

                        <div class="env-var">
                            <label>JIRA_EPIC_API_TOKEN</label>
                            <code>your-api-token-here</code>
                            <p>
                                Your Jira API token (generate from Account
                                Settings)
                            </p>
                        </div>
                    </div>

                    <div class="setup-steps">
                        <h3>Setup Steps:</h3>
                        <ol>
                            <li>
                                <strong>Generate API Token:</strong>
                                <ul>
                                    <li>
                                        Go to
                                        <a
                                            href="https://id.atlassian.com/manage-profile/security/api-tokens"
                                            target="_blank"
                                            rel="noopener"
                                            >Atlassian Account Settings</a
                                        >
                                    </li>
                                    <li>Click "Create API token"</li>
                                    <li>
                                        Give it a label (e.g., "Epic Analyzer")
                                    </li>
                                    <li>Copy the generated token</li>
                                </ul>
                            </li>
                            <li>
                                <strong>Create .env file:</strong>
                                <pre><code>JIRA_EPIC_WORKSPACE=your-company.atlassian.net
JIRA_EPIC_EMAIL=your-email@company.com
JIRA_EPIC_API_TOKEN=your-api-token-here</code></pre>
                            </li>
                            <li>
                                <strong>Restart your backend server</strong> to
                                load the new environment variables
                            </li>
                            <li>
                                <strong>Test the connection</strong> using the
                                button below
                            </li>
                        </ol>
                    </div>

                    <div class="warning-box">
                        <div class="warning-icon">⚠️</div>
                        <div class="warning-content">
                            <strong>Security Note:</strong>
                            <p>
                                Never commit your API token to version control.
                                Keep your .env file in .gitignore and use
                                environment variables in production.
                            </p>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Connection Test -->
            <div class="test-section">
                <div class="card">
                    <h2>Test Connection</h2>
                    <p>
                        Click the button below to test your Jira API connection:
                    </p>

                    <div class="test-container">
                        <button
                            @click="testConnection"
                            class="btn btn-primary btn-lg"
                            :disabled="isTesting"
                            :class="{
                                'btn-success': connectionStatus === 'success',
                                'btn-error': connectionStatus === 'error',
                            }"
                        >
                            <span v-if="isTesting" class="spinner"></span>
                            <span v-else-if="connectionStatus === 'success'"
                                >✅</span
                            >
                            <span v-else-if="connectionStatus === 'error'"
                                >❌</span
                            >
                            <span v-else>🔌</span>
                            {{ testButtonText }}
                        </button>

                        <div
                            v-if="testResult"
                            class="test-result"
                            :class="connectionStatus"
                        >
                            <div class="result-header">
                                <span class="result-icon">
                                    {{
                                        connectionStatus === "success"
                                            ? "✅"
                                            : "❌"
                                    }}
                                </span>
                                <span class="result-title">
                                    {{
                                        connectionStatus === "success"
                                            ? "Connection Successful!"
                                            : "Connection Failed"
                                    }}
                                </span>
                            </div>

                            <div class="result-details">
                                <p>
                                    <strong>Status:</strong>
                                    {{ testResult.status }}
                                </p>
                                <p>
                                    <strong>Response Time:</strong>
                                    {{ testResult.responseTime }}ms
                                </p>
                                <p v-if="testResult.message">
                                    <strong>Message:</strong>
                                    {{ testResult.message }}
                                </p>

                                <div
                                    v-if="testResult.details"
                                    class="result-data"
                                >
                                    <h4>Connection Details:</h4>
                                    <pre><code>{{ JSON.stringify(testResult.details, null, 2) }}</code></pre>
                                </div>

                                <div
                                    v-if="testResult.error"
                                    class="error-details"
                                >
                                    <h4>Error Details:</h4>
                                    <pre><code>{{ testResult.error }}</code></pre>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Troubleshooting -->
            <div class="troubleshooting-section">
                <div class="card">
                    <h2>Troubleshooting</h2>

                    <div class="accordion">
                        <div class="accordion-item">
                            <button
                                @click="toggleAccordion('common-issues')"
                                class="accordion-header"
                                :class="{
                                    active: activeAccordion === 'common-issues',
                                }"
                            >
                                <span>Common Issues</span>
                                <span class="accordion-icon">{{
                                    activeAccordion === "common-issues"
                                        ? "▼"
                                        : "▶"
                                }}</span>
                            </button>
                            <div
                                v-if="activeAccordion === 'common-issues'"
                                class="accordion-content"
                            >
                                <div class="issue">
                                    <h4>❌ 401 Unauthorized</h4>
                                    <p>
                                        Check your email and API token. Make
                                        sure the API token is valid and not
                                        expired.
                                    </p>
                                </div>
                                <div class="issue">
                                    <h4>❌ 404 Not Found</h4>
                                    <p>
                                        Verify your JIRA_EPIC_WORKSPACE is
                                        correct. It should include the full
                                        domain (e.g., company.atlassian.net).
                                    </p>
                                </div>
                                <div class="issue">
                                    <h4>❌ CORS Error</h4>
                                    <p>
                                        This usually means your backend server
                                        is not running or not properly
                                        configured for proxy requests.
                                    </p>
                                </div>
                                <div class="issue">
                                    <h4>❌ Network Error</h4>
                                    <p>
                                        Check if your backend server is running
                                        and accessible at the expected port.
                                    </p>
                                </div>
                            </div>
                        </div>

                        <div class="accordion-item">
                            <button
                                @click="toggleAccordion('permissions')"
                                class="accordion-header"
                                :class="{
                                    active: activeAccordion === 'permissions',
                                }"
                            >
                                <span>Required Permissions</span>
                                <span class="accordion-icon">{{
                                    activeAccordion === "permissions"
                                        ? "▼"
                                        : "▶"
                                }}</span>
                            </button>
                            <div
                                v-if="activeAccordion === 'permissions'"
                                class="accordion-content"
                            >
                                <p>
                                    Your Jira account needs the following
                                    permissions:
                                </p>
                                <ul>
                                    <li>Browse Projects</li>
                                    <li>View Issues</li>
                                    <li>
                                        View Development Tools (for epic data)
                                    </li>
                                </ul>
                                <p>
                                    Contact your Jira administrator if you don't
                                    have these permissions.
                                </p>
                            </div>
                        </div>

                        <div class="accordion-item">
                            <button
                                @click="toggleAccordion('api-limits')"
                                class="accordion-header"
                                :class="{
                                    active: activeAccordion === 'api-limits',
                                }"
                            >
                                <span>API Rate Limits</span>
                                <span class="accordion-icon">{{
                                    activeAccordion === "api-limits"
                                        ? "▼"
                                        : "▶"
                                }}</span>
                            </button>
                            <div
                                v-if="activeAccordion === 'api-limits'"
                                class="accordion-content"
                            >
                                <p>
                                    Atlassian has rate limits for API requests:
                                </p>
                                <ul>
                                    <li>
                                        <strong>Cloud:</strong> 10 requests per
                                        second per IP
                                    </li>
                                    <li>
                                        <strong>Server/DC:</strong> Varies by
                                        configuration
                                    </li>
                                </ul>
                                <p>
                                    If you encounter rate limit errors, wait a
                                    moment and try again.
                                </p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { ref, computed } from "vue";

export default {
    name: "AdminSetup",
    setup() {
        const isTesting = ref(false);
        const connectionStatus = ref(null); // 'success', 'error', null
        const testResult = ref(null);
        const activeAccordion = ref(null);

        const testButtonText = computed(() => {
            if (isTesting.value) return "Testing Connection...";
            if (connectionStatus.value === "success")
                return "Connection Successful";
            if (connectionStatus.value === "error")
                return "Connection Failed - Retry";
            return "Test Connection";
        });

        const testConnection = async () => {
            isTesting.value = true;
            connectionStatus.value = null;
            testResult.value = null;

            const startTime = Date.now();

            try {
                const response = await fetch("/api/ping", {
                    method: "GET",
                    headers: {
                        "Content-Type": "application/json",
                    },
                });

                const responseTime = Date.now() - startTime;

                if (response.ok) {
                    let data = null;
                    try {
                        data = await response.json();
                    } catch (e) {
                        // Response might not be JSON
                        data = await response.text();
                    }

                    connectionStatus.value = "success";
                    testResult.value = {
                        status: `${response.status} ${response.statusText}`,
                        responseTime,
                        message: "Successfully connected to Jira API!",
                        details: data,
                    };
                } else {
                    connectionStatus.value = "error";
                    let errorData = null;
                    try {
                        errorData = await response.json();
                    } catch (e) {
                        errorData = await response.text();
                    }

                    testResult.value = {
                        status: `${response.status} ${response.statusText}`,
                        responseTime,
                        message: "Failed to connect to Jira API",
                        error: errorData,
                    };
                }
            } catch (error) {
                const responseTime = Date.now() - startTime;
                connectionStatus.value = "error";
                testResult.value = {
                    status: "Network Error",
                    responseTime,
                    message: "Could not reach the backend server",
                    error: error.message,
                };
            } finally {
                isTesting.value = false;
            }
        };

        const toggleAccordion = (section) => {
            activeAccordion.value =
                activeAccordion.value === section ? null : section;
        };

        return {
            isTesting,
            connectionStatus,
            testResult,
            testButtonText,
            activeAccordion,
            testConnection,
            toggleAccordion,
        };
    },
};
</script>

<style lang="scss" scoped>
.admin-setup-page {
    padding: var(--spacing-xl) 0;
    min-height: 100vh;
}

.page-header {
    text-align: center;
    margin-bottom: var(--spacing-xxl);

    h1 {
        margin-bottom: var(--spacing-sm);
    }

    p {
        color: var(--text-secondary);
        font-size: var(--font-lg);
    }
}

.setup-section,
.test-section,
.troubleshooting-section {
    margin-bottom: var(--spacing-xxl);

    .card {
        max-width: 800px;
        margin: 0 auto;
    }

    h2 {
        margin-bottom: var(--spacing-lg);
        color: var(--text-primary);
    }
}

.env-variables {
    margin: var(--spacing-lg) 0;

    .env-var {
        margin-bottom: var(--spacing-lg);
        padding: var(--spacing-md);
        background-color: var(--border-light);
        border-radius: var(--radius-md);

        label {
            display: block;
            font-weight: var(--font-semibold);
            color: var(--text-primary);
            margin-bottom: var(--spacing-xs);
        }

        code {
            display: block;
            background-color: var(--card-bg);
            padding: var(--spacing-sm);
            border-radius: var(--radius-sm);
            margin-bottom: var(--spacing-xs);
            font-family: "Monaco", "Menlo", "Ubuntu Mono", monospace;
        }

        p {
            margin: 0;
            font-size: var(--font-sm);
            color: var(--text-secondary);
        }
    }
}

.setup-steps {
    margin: var(--spacing-lg) 0;

    h3 {
        margin-bottom: var(--spacing-md);
        color: var(--text-primary);
    }

    ol {
        padding-left: var(--spacing-lg);

        li {
            margin-bottom: var(--spacing-md);
            line-height: 1.6;

            strong {
                color: var(--text-primary);
            }

            ul {
                margin-top: var(--spacing-sm);
                margin-bottom: var(--spacing-sm);
            }
        }
    }

    pre {
        margin-top: var(--spacing-sm);
        background-color: var(--border-light);
        padding: var(--spacing-md);
        border-radius: var(--radius-md);
        overflow-x: auto;

        code {
            background: none;
            padding: 0;
        }
    }

    a {
        color: var(--primary-color);

        &:hover {
            text-decoration: underline;
        }
    }
}

.warning-box {
    display: flex;
    gap: var(--spacing-md);
    padding: var(--spacing-md);
    background-color: var(--warning-light);
    border: 1px solid var(--warning-color);
    border-radius: var(--radius-md);
    margin-top: var(--spacing-lg);

    .warning-icon {
        font-size: var(--font-lg);
        flex-shrink: 0;
    }

    .warning-content {
        strong {
            color: var(--warning-color);
            display: block;
            margin-bottom: var(--spacing-xs);
        }

        p {
            margin: 0;
            color: var(--text-primary);
            font-size: var(--font-sm);
        }
    }
}

.test-container {
    text-align: center;

    .btn {
        margin-bottom: var(--spacing-lg);
        min-width: 200px;

        .spinner {
            margin-right: var(--spacing-sm);
        }
    }
}

.test-result {
    @include card-style;
    padding: var(--spacing-lg);
    text-align: left;
    max-width: 600px;
    margin: 0 auto;

    &.success {
        border-color: var(--success-color);
        background-color: var(--success-light);
    }

    &.error {
        border-color: var(--error-color);
        background-color: var(--error-light);
    }

    .result-header {
        display: flex;
        align-items: center;
        gap: var(--spacing-sm);
        margin-bottom: var(--spacing-md);

        .result-icon {
            font-size: var(--font-lg);
        }

        .result-title {
            font-weight: var(--font-semibold);
            font-size: var(--font-lg);
        }
    }

    .result-details {
        p {
            margin-bottom: var(--spacing-sm);

            strong {
                color: var(--text-primary);
            }
        }

        .result-data,
        .error-details {
            margin-top: var(--spacing-md);

            h4 {
                margin-bottom: var(--spacing-sm);
                color: var(--text-primary);
            }

            pre {
                background-color: var(--card-bg);
                padding: var(--spacing-md);
                border-radius: var(--radius-sm);
                overflow-x: auto;
                max-height: 200px;
                border: 1px solid var(--border-color);

                code {
                    background: none;
                    padding: 0;
                    font-size: var(--font-sm);
                }
            }
        }
    }
}

.accordion {
    .accordion-item {
        border: 1px solid var(--border-color);
        border-radius: var(--radius-md);
        margin-bottom: var(--spacing-sm);
        overflow: hidden;

        .accordion-header {
            @include button-reset;
            width: 100%;
            padding: var(--spacing-md);
            background-color: var(--hover-bg);
            @include flex-between;
            font-weight: var(--font-medium);
            transition: var(--transition-normal);

            &:hover {
                background-color: var(--border-light);
            }

            &.active {
                background-color: var(--primary-light);
                color: var(--primary-color);
            }

            .accordion-icon {
                transition: var(--transition-normal);
            }
        }

        .accordion-content {
            padding: var(--spacing-md);
            border-top: 1px solid var(--border-color);

            .issue {
                margin-bottom: var(--spacing-md);

                &:last-child {
                    margin-bottom: 0;
                }

                h4 {
                    margin-bottom: var(--spacing-xs);
                    color: var(--text-primary);
                }

                p {
                    margin: 0;
                    color: var(--text-secondary);
                    font-size: var(--font-sm);
                }
            }

            ul {
                margin: var(--spacing-sm) 0;
                padding-left: var(--spacing-lg);

                li {
                    margin-bottom: var(--spacing-xs);
                    color: var(--text-secondary);
                }
            }
        }
    }
}

@media (max-width: 768px) {
    .warning-box {
        flex-direction: column;
        text-align: center;
    }

    .test-container .btn {
        width: 100%;
    }

    .env-variables .env-var code {
        word-break: break-all;
    }
}
</style>
