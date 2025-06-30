<template>
    <div class="dependency-graph">
        <div class="graph-header">
            <div class="header-top">
                <h3>
                    Ticket Dependencies
                    <span v-if="phaseCount > 0" class="phase-counter"
                        >({{ phaseCount }} phases)</span
                    >
                </h3>
                <div class="graph-controls">
                    <button @click="resetZoom" class="btn btn-secondary">
                        Reset View
                    </button>
                    <!-- these toggles are for the development only, not production ready just don't break them -->
                    <!-- <button @click="toggleLayout" class="btn btn-secondary">
                        {{
                            layoutType === "phase"
                                ? "Force Layout"
                                : layoutType === "force"
                                  ? "Tree Layout"
                                  : "Phase Layout"
                        }}
                    </button> -->
                    <!-- <button @click="debugPhases" class="btn btn-secondary">
                        Debug Phases
                    </button> -->
                    <button @click="toggleFullscreen" class="btn btn-secondary">
                        {{ isFullscreen ? "Exit Fullscreen" : "Fullscreen" }}
                    </button>
                    <button
                        v-if="availableStatuses.size > 0"
                        @click="toggleFilters"
                        class="btn btn-secondary filter-toggle"
                        :class="{ active: showFilters }"
                        :aria-expanded="showFilters"
                        aria-controls="filter-row"
                    >
                        <span class="filter-icon">▼</span>
                        Filters
                        <span
                            v-if="statusFilters.size > 0"
                            class="filter-count"
                        >
                            {{ statusFilters.size }}
                        </span>
                    </button>
                    <button
                        v-if="assignees && assignees.length > 0"
                        @click="toggleAssigneeRow"
                        class="btn btn-secondary filter-toggle"
                        :class="{ active: showAssigneeRow }"
                        :aria-expanded="showAssigneeRow"
                        aria-controls="assignee-row"
                    >
                        👤 Assignees
                        <span
                            v-if="assigneeFilters.size > 0"
                            class="filter-count"
                        >
                            {{ assigneeFilters.size }}
                        </span>
                    </button>
                </div>
            </div>
            <div
                id="filter-row"
                class="filter-row"
                v-if="availableStatuses.size > 0"
                :class="{ expanded: showFilters }"
            >
                <div class="status-filters">
                    <span class="filter-label">Hide Status:</span>
                    <div class="filter-items">
                        <label
                            v-for="status in Array.from(availableStatuses)"
                            :key="status"
                            class="status-filter-item"
                        >
                            <input
                                type="checkbox"
                                :value="status"
                                :checked="statusFilters.has(status)"
                                @change="toggleStatusFilter(status)"
                            />
                            <span class="status-name">{{ status }}</span>
                        </label>
                    </div>
                </div>
            </div>
            <div
                id="assignee-row"
                class="assignee-row"
                v-if="assignees && assignees.length > 0 && showAssigneeRow"
            >
                <div class="assignee-filters">
                    <span class="filter-label">👤</span>
                    <div class="assignee-items">
                        <label
                            v-for="assignee in assignees"
                            :key="assignee.accountId"
                            class="assignee-filter-item"
                            :class="{
                                active: assigneeFilters.has(assignee.accountId),
                            }"
                            :title="assignee.displayName || 'Unknown'"
                        >
                            <input
                                type="checkbox"
                                :value="assignee.accountId"
                                :checked="
                                    assigneeFilters.has(assignee.accountId)
                                "
                                @change="
                                    toggleAssigneeFilter(assignee.accountId)
                                "
                            />
                            <img
                                v-if="assignee.avatarUrl"
                                :src="assignee.avatarUrl"
                                :alt="assignee.displayName || 'Assignee'"
                                class="assignee-avatar"
                            />
                            <span class="assignee-name">{{
                                (assignee.displayName || "Unknown").split(
                                    " ",
                                )[0]
                            }}</span>
                        </label>
                    </div>
                </div>
            </div>
        </div>
        <div class="graph-legend">
            <div class="legend-item">
                <div class="legend-color epic"></div>
                <span>Epic</span>
            </div>
            <div class="legend-item">
                <div class="legend-color new"></div>
                <span>Backlog</span>
            </div>
            <div class="legend-item">
                <div class="legend-color indeterminate"></div>
                <span>In Progress</span>
            </div>
            <div class="legend-item">
                <div class="legend-color done"></div>
                <span>Done</span>
            </div>
            <div class="legend-item">
                <div class="legend-line blocks"></div>
                <span>Blocks</span>
            </div>
            <div class="legend-item">
                <div class="legend-line epic-link"></div>
                <span>Epic Link</span>
            </div>
            <div class="legend-item">
                <div class="legend-line status-line"></div>
                <span
                    >Connection Lines: Colored by Source Ticket
                    Classification</span
                >
            </div>
            <div v-if="layoutType === 'phases'" class="legend-item">
                <div class="legend-color phase-indicator"></div>
                <span>Phases: Concurrent Work Groups</span>
            </div>
        </div>
        <div
            ref="graphContainer"
            class="graph-container"
            :style="{ height: containerHeight + 'px' }"
        ></div>
        <div
            class="resize-handle"
            @mousedown="startResize"
            :class="{ active: isResizing }"
        ></div>
        <dialog
            ref="nodeDialog"
            class="node-dialog"
            @close="handleDialogClose"
            @click="handleDialogBackdropClick"
            role="dialog"
            :aria-labelledby="
                selectedNode ? `dialog-title-${selectedNode.id}` : undefined
            "
            aria-modal="true"
        >
            <div v-if="selectedNode" class="dialog-content">
                <div class="dialog-header">
                    <strong :id="`dialog-title-${selectedNode.id}`">{{
                        selectedNode.id
                    }}</strong>
                    <button
                        @click="closeDialog"
                        class="close-dialog"
                        aria-label="Close dialog"
                        type="button"
                        style="padding: 0 16px"
                    >
                        &times;
                    </button>
                </div>
                <div class="dialog-body">
                    <div class="dialog-label">{{ selectedNode.label }}</div>
                    <div v-if="selectedNode.status" class="dialog-status">
                        Status:
                        <span
                            :class="`status-${selectedNode.status.toLowerCase().replace(/\s+/g, '-')}`"
                            >{{ selectedNode.status }}</span
                        >
                    </div>
                    <div
                        v-if="
                            selectedNode.points ||
                            selectedNode.storyPoints ||
                            selectedNode.story_points
                        "
                        class="dialog-points"
                    >
                        Story Points:
                        <span class="points-value">
                            {{
                                selectedNode.points ||
                                selectedNode.storyPoints ||
                                selectedNode.story_points
                            }}
                        </span>
                    </div>
                    <div v-if="selectedNode.assignee" class="dialog-assignee">
                        Assignee:
                        <span class="assignee-value">{{
                            selectedNode.assignee
                        }}</span>
                    </div>
                    <div class="dialog-actions">
                        <button
                            @click="copyNodeLink(selectedNode)"
                            class="dialog-btn copy-btn"
                        >
                            📋 Copy Link
                        </button>
                        <button
                            @click="openNodeInNewTab(selectedNode)"
                            class="dialog-btn open-btn"
                        >
                            🔗 Open in New Tab
                        </button>
                    </div>
                </div>
            </div>
        </dialog>
    </div>
</template>

<script>
import { ref, onMounted, onUnmounted, watch, nextTick } from "vue";
import * as d3 from "d3";

export default {
    name: "DependencyGraph",
    props: {
        jiraBaseUrl: String,
        graphData: {
            type: Object,
            required: true,
            validator: (value) => {
                return (
                    value &&
                    value.nodes &&
                    value.edges &&
                    Array.isArray(value.nodes) &&
                    Array.isArray(value.edges)
                );
            },
        },
        assignees: {
            type: Array,
            default: () => [],
        },
    },
    setup(props) {
        const graphContainer = ref(null);
        const nodeDialog = ref(null);
        const selectedNode = ref(null);
        const layoutType = ref("phase");
        const phaseCount = ref(0);
        const debugInfo = ref(null);
        const isFullscreen = ref(false);
        const containerHeight = ref(600);
        const isResizing = ref(false);
        const statusFilters = ref(new Set(["DEFERRED"]));
        const availableStatuses = ref(new Set());
        const showFilters = ref(false);
        const assigneeFilters = ref(new Set());
        const availableAssignees = ref(new Set());
        const showAssigneeRow = ref(false);

        let svg, simulation, nodes, links, nodeElements, linkElements, zoom;
        let isInitialized = false;
        let lastDataHash = null;
        let mutationObserver = null;

        // Function to get line color based on source node classification
        const getLineColorByStatus = (sourceNodeRef) => {
            let sourceNode;
            if (typeof sourceNodeRef === "string") {
                sourceNode = nodes.find((n) => n.id === sourceNodeRef);
            } else if (sourceNodeRef && sourceNodeRef.id) {
                sourceNode = sourceNodeRef;
            } else {
                return "var(--ctp-overlay1)"; // Default gray color
            }

            if (!sourceNode || !sourceNode.classification) {
                return "var(--ctp-overlay1)"; // Default gray color
            }

            const classification = sourceNode.classification.toLowerCase();

            if (classification === "done") {
                return "var(--ctp-green)"; // Green for completed
            } else if (classification === "epic") {
                return "var(--ctp-mauve)"; // Purple for epic
            } else if (classification === "indeterminate") {
                return "var(--ctp-yellow)"; // Yellow for in progress
            } else if (classification === "new") {
                return "var(--ctp-blue)"; // Blue for backlog/new
            } else {
                return "var(--ctp-overlay1)"; // Default gray for unknown classification
            }
        };

        // Function to update link colors based on current node statuses
        const updateLinkColors = () => {
            if (linkElements) {
                linkElements.attr("stroke", (d) =>
                    getLineColorByStatus(d.source.id || d.source),
                );
            }
        };

        // Helper function to calculate connection point on card edge
        const getCardConnectionPoint = (sourceNode, targetNode) => {
            // Resolve string IDs to actual node objects
            const resolvedSourceNode =
                typeof sourceNode === "string"
                    ? nodes.find((n) => n.id === sourceNode)
                    : sourceNode;
            const resolvedTargetNode =
                typeof targetNode === "string"
                    ? nodes.find((n) => n.id === targetNode)
                    : targetNode;

            // Add safety checks
            if (
                !resolvedSourceNode ||
                !resolvedTargetNode ||
                resolvedSourceNode.x === undefined ||
                resolvedSourceNode.y === undefined ||
                resolvedTargetNode.x === undefined ||
                resolvedTargetNode.y === undefined
            ) {
                return {
                    source: resolvedSourceNode || { x: 0, y: 0 },
                    target: resolvedTargetNode || { x: 0, y: 0 },
                    path: "M0,0 L0,0",
                };
            }

            // Determine if nodes are epic cards (larger)
            const isSourceEpic =
                resolvedSourceNode.status &&
                resolvedSourceNode.status === "Epic";
            const isTargetEpic =
                resolvedTargetNode.status &&
                resolvedTargetNode.status === "Epic";

            // Card dimensions
            const sourceWidth = isSourceEpic ? 160 : 140;
            const sourceHeight = isSourceEpic ? 90 : 80;
            const targetWidth = isTargetEpic ? 160 : 140;
            const targetHeight = isTargetEpic ? 90 : 80;

            // Calculate vector from source to target
            const dx = resolvedTargetNode.x - resolvedSourceNode.x;
            const dy = resolvedTargetNode.y - resolvedSourceNode.y;
            const distance = Math.sqrt(dx * dx + dy * dy);

            if (distance === 0) {
                return {
                    source: {
                        x: resolvedSourceNode.x,
                        y: resolvedSourceNode.y,
                    },
                    target: {
                        x: resolvedTargetNode.x,
                        y: resolvedTargetNode.y,
                    },
                    path: "M0,0 L0,0",
                };
            }

            // Normalize vector
            const unitX = dx / distance;
            const unitY = dy / distance;

            // Calculate intersection with source card edge
            const sourceIntersection = {
                x: resolvedSourceNode.x + unitX * (sourceWidth / 2),
                y: resolvedSourceNode.y + unitY * (sourceHeight / 2),
            };

            // Calculate intersection with target card edge
            const targetIntersection = {
                x: resolvedTargetNode.x - unitX * (targetWidth / 2),
                y: resolvedTargetNode.y - unitY * (targetHeight / 2),
            };

            // Create curved path for force layout too
            const curveDistance = distance * 0.2;
            const perpX = -unitY * curveDistance;
            const perpY = unitX * curveDistance;

            const midX = (sourceIntersection.x + targetIntersection.x) / 2;
            const midY = (sourceIntersection.y + targetIntersection.y) / 2;

            // Add slight curve to avoid overlapping
            const controlX = midX + perpX * 0.3;
            const controlY = midY + perpY * 0.3;

            const path = `M${sourceIntersection.x},${sourceIntersection.y} Q${controlX},${controlY} ${targetIntersection.x},${targetIntersection.y}`;

            return {
                source: sourceIntersection,
                target: targetIntersection,
                path: path,
            };
        };

        // Helper function to calculate side connections for phase layout
        const getPhaseConnectionPoint = (sourceNode, targetNode) => {
            // Resolve string IDs to actual node objects
            const resolvedSourceNode =
                typeof sourceNode === "string"
                    ? nodes.find((n) => n.id === sourceNode)
                    : sourceNode;
            const resolvedTargetNode =
                typeof targetNode === "string"
                    ? nodes.find((n) => n.id === targetNode)
                    : targetNode;

            // Add null checks to prevent undefined errors
            if (!resolvedSourceNode || !resolvedTargetNode) {
                console.warn(
                    "Missing node(s) referenced in edge - skipping link:",
                    {
                        sourceId:
                            typeof sourceNode === "string"
                                ? sourceNode
                                : sourceNode?.id,
                        targetId:
                            typeof targetNode === "string"
                                ? targetNode
                                : targetNode?.id,
                        sourceFound: !!resolvedSourceNode,
                        targetFound: !!resolvedTargetNode,
                    },
                );
                return null; // Return null to indicate this link should be skipped
            }

            if (
                resolvedSourceNode.x === undefined ||
                resolvedSourceNode.y === undefined ||
                resolvedTargetNode.x === undefined ||
                resolvedTargetNode.y === undefined
            ) {
                console.warn(
                    "Invalid node positions in getPhaseConnectionPoint:",
                    {
                        sourceNode: resolvedSourceNode,
                        targetNode: resolvedTargetNode,
                    },
                );
                return {
                    source: resolvedSourceNode || { x: 0, y: 0 },
                    target: resolvedTargetNode || { x: 0, y: 0 },
                    path: "M0,0 L0,0",
                };
            }

            const isSourceEpic =
                resolvedSourceNode.status &&
                resolvedSourceNode.status === "Epic";
            const isTargetEpic =
                resolvedTargetNode.status &&
                resolvedTargetNode.status === "Epic";

            const sourceWidth = isSourceEpic ? 160 : 140;
            const targetWidth = isTargetEpic ? 160 : 140;
            const sourceHeight = isSourceEpic ? 90 : 80;
            const targetHeight = isTargetEpic ? 90 : 80;

            // Calculate connection points on card edges
            const sourceX =
                resolvedSourceNode.x +
                (resolvedTargetNode.x > resolvedSourceNode.x
                    ? sourceWidth / 2
                    : -sourceWidth / 2);
            const targetX =
                resolvedTargetNode.x +
                (resolvedTargetNode.x > resolvedSourceNode.x
                    ? -targetWidth / 2
                    : targetWidth / 2);

            const sourceY = resolvedSourceNode.y;
            const targetY = resolvedTargetNode.y;

            // Create right-angle path (orthogonal routing)
            // The path will go: source -> horizontal -> vertical -> target
            const horizontalDistance = Math.abs(targetX - sourceX);
            const verticalDistance = Math.abs(targetY - sourceY);

            // Create a right-angle path with a horizontal segment followed by vertical
            // Add some padding to avoid overlapping with nodes
            const horizontalExtension = Math.min(horizontalDistance * 0.6, 80);
            const midX =
                sourceX +
                (targetX > sourceX
                    ? horizontalExtension
                    : -horizontalExtension);

            let path;
            if (Math.abs(verticalDistance) < 10) {
                // Nearly same height - direct horizontal line
                path = `M${sourceX},${sourceY} L${targetX},${targetY}`;
            } else {
                // Right-angle path: horizontal then vertical
                path = `M${sourceX},${sourceY} L${midX},${sourceY} L${midX},${targetY} L${targetX},${targetY}`;
            }

            return {
                source: {
                    x: sourceX,
                    y: sourceY,
                },
                target: {
                    x: targetX,
                    y: targetY,
                },
                path: path,
            };
        };

        // Safe helper for getting connection points during transitions
        const getSafeConnectionPoint = (d, attr, usePhaseLayout = false) => {
            try {
                // Resolve source and target nodes from IDs if needed
                const sourceNode =
                    typeof d.source === "string"
                        ? nodes.find((n) => n.id === d.source)
                        : d.source;
                const targetNode =
                    typeof d.target === "string"
                        ? nodes.find((n) => n.id === d.target)
                        : d.target;

                // Check if nodes exist
                if (!sourceNode || !targetNode) {
                    console.warn("Missing node(s) in getSafeConnectionPoint:", {
                        sourceId:
                            typeof d.source === "string"
                                ? d.source
                                : d.source?.id,
                        targetId:
                            typeof d.target === "string"
                                ? d.target
                                : d.target?.id,
                        sourceFound: !!sourceNode,
                        targetFound: !!targetNode,
                    });
                    return null; // Return null to indicate missing nodes
                }

                const connections = usePhaseLayout
                    ? getPhaseConnectionPoint(sourceNode, targetNode)
                    : getCardConnectionPoint(sourceNode, targetNode);

                if (connections === null) {
                    return null; // Missing node case
                }

                if (
                    !connections ||
                    !connections.source ||
                    !connections.target ||
                    !connections.path
                ) {
                    return connections?.path || "M0,0 L0,0";
                }

                if (attr === "path") return connections.path;
                if (attr === "source.x") return connections.source.x || 0;
                if (attr === "source.y") return connections.source.y || 0;
                if (attr === "target.x") return connections.target.x || 0;
                if (attr === "target.y") return connections.target.y || 0;
                return connections.path || "M0,0 L0,0";
            } catch (error) {
                console.warn("Error calculating connection point:", error);
                return "M0,0 L0,0";
            }
        };

        // Ensure arrow markers are available in the DOM
        const ensureMarkers = () => {
            if (!svg) return;

            // Check if markers already exist
            const existingDefs = svg.select("defs");
            const blockMarker = existingDefs.select("#arrow-blocks");
            const epicMarker = existingDefs.select("#arrow-epic");

            if (blockMarker.empty() || epicMarker.empty()) {
                // Re-create markers if they don't exist
                const defs = existingDefs.empty()
                    ? svg.append("defs")
                    : existingDefs;

                if (blockMarker.empty()) {
                    defs.append("marker")
                        .attr("id", "arrow-blocks")
                        .attr("viewBox", "0 0 10 10")
                        .attr("refX", 9)
                        .attr("refY", 3)
                        .attr("markerWidth", 6)
                        .attr("markerHeight", 6)
                        .attr("orient", "auto")
                        .append("path")
                        .attr("d", "M0,0 L0,6 L9,3 z")
                        .attr("fill", "var(--ctp-red)");
                }

                if (epicMarker.empty()) {
                    defs.append("marker")
                        .attr("id", "arrow-epic")
                        .attr("viewBox", "0 0 10 10")
                        .attr("refX", 9)
                        .attr("refY", 3)
                        .attr("markerWidth", 6)
                        .attr("markerHeight", 6)
                        .attr("orient", "auto")
                        .append("path")
                        .attr("d", "M0,0 L0,6 L9,3 z")
                        .attr("fill", "var(--ctp-blue)");
                }
            }
        };

        const calculatePhases = () => {
            const nodeMap = new Map();
            const phases = [];

            // Always use original graph data for phase calculation
            // D3 modifies the links array, replacing string IDs with objects
            const originalNodes = props.graphData.nodes;
            const originalLinks = props.graphData.edges;

            // Filter out epic nodes - we only want work tickets in phases
            const workNodes = originalNodes.filter(
                (node) => !(node.status && node.status === "Epic"),
            );

            // Create node map only for work tickets
            workNodes.forEach((node) => {
                nodeMap.set(node.id, {
                    ...node,
                    dependencies: [],
                    dependents: [],
                    phase: -1,
                });
            });

            // Build dependency relationships - ONLY for "Blocks" relationships between work tickets
            // Epic links and epic nodes are ignored for phase calculation
            originalLinks.forEach((link) => {
                const source = nodeMap.get(link.from);
                const target = nodeMap.get(link.to);

                // Only process blocking relationships between work tickets (non-epic nodes)
                if (source && target && link.type === "Blocks") {
                    // Only "Blocks" relationships create phase dependencies
                    // Source blocks target, so target depends on source
                    target.dependencies.push(source.id);
                    source.dependents.push(target.id);
                }
            });

            // Calculate phases using topological sort
            let currentPhase = 0;
            let remainingNodes = new Set(nodeMap.keys());

            while (remainingNodes.size > 0) {
                const currentPhaseNodes = [];

                // Find nodes with no unresolved dependencies
                for (const nodeId of remainingNodes) {
                    const node = nodeMap.get(nodeId);
                    const unresolvedDeps = node.dependencies.filter((depId) =>
                        remainingNodes.has(depId),
                    );

                    if (unresolvedDeps.length === 0) {
                        currentPhaseNodes.push(nodeId);
                        node.phase = currentPhase;
                    }
                }

                if (currentPhaseNodes.length === 0) {
                    // Circular dependency or isolated nodes
                    const isolated = Array.from(remainingNodes)[0];
                    nodeMap.get(isolated).phase = currentPhase;
                    currentPhaseNodes.push(isolated);
                }

                phases[currentPhase] = currentPhaseNodes;
                currentPhaseNodes.forEach((nodeId) =>
                    remainingNodes.delete(nodeId),
                );
                currentPhase++;
            }

            console.log("Phase calculation completed:", {
                totalWorkNodes: workNodes.length,
                totalPhases: phases.length,
                blockingLinks: originalLinks.filter((l) => l.type === "Blocks")
                    .length,
                phases: phases.map((phase, index) => ({
                    phase: index + 1,
                    nodes: phase.length,
                    nodeIds: phase,
                })),
                epicNodesFiltered: originalNodes.length - workNodes.length,
            });

            // Update phase count for display
            phaseCount.value = phases.length;

            return { nodeMap, phases };
        };

        const initializeGraph = () => {
            if (!graphContainer.value || !props.graphData) return;

            // Extract and store available statuses
            const statuses = new Set(
                props.graphData.nodes
                    .map((node) => node.status)
                    .filter(Boolean),
            );
            availableStatuses.value = statuses;

            // Clear existing graph
            if (svg) {
                svg.remove();
            }

            // Force container to update dimensions
            const containerRect = graphContainer.value.getBoundingClientRect();
            const width = Math.max(
                containerRect.width,
                graphContainer.value.clientWidth,
                400,
            );
            const height = Math.max(
                containerRect.height,
                graphContainer.value.clientHeight,
                300,
            );

            console.log("Initializing graph with dimensions:", {
                width,
                height,
                isFullscreen: isFullscreen.value,
            });

            svg = d3
                .select(graphContainer.value)
                .append("svg")
                .attr("width", width)
                .attr("height", height)
                .style("background-color", "transparent");

            // Add zoom and pan behavior
            zoom = d3
                .zoom()
                .scaleExtent([0.1, 3])
                .on("zoom", (event) => {
                    svg.select(".graph-group").attr(
                        "transform",
                        event.transform,
                    );
                });

            svg.call(zoom);

            // Create main group for zooming/panning
            const g = svg.append("g").attr("class", "graph-group");

            // Define arrow markers for different edge types
            const defs = svg.append("defs");

            // Blocks arrow (using Catppuccin red)
            defs.append("marker")
                .attr("id", "arrow-blocks")
                .attr("viewBox", "0 0 10 10")
                .attr("refX", 9)
                .attr("refY", 3)
                .attr("markerWidth", 6)
                .attr("markerHeight", 6)
                .attr("orient", "auto")
                .append("path")
                .attr("d", "M0,0 L0,6 L9,3 z")
                .attr("fill", "var(--ctp-red)");

            // Epic link arrow (using Catppuccin blue)
            defs.append("marker")
                .attr("id", "arrow-epic")
                .attr("viewBox", "0 0 10 10")
                .attr("refX", 9)
                .attr("refY", 3)
                .attr("markerWidth", 6)
                .attr("markerHeight", 6)
                .attr("orient", "auto")
                .append("path")
                .attr("d", "M0,0 L0,6 L9,3 z")
                .attr("fill", "var(--ctp-blue)");

            // Filter nodes by status and assignee
            const filteredNodes = props.graphData.nodes.filter((node) => {
                // Filter by status
                if (statusFilters.value.has(node.status)) {
                    return false;
                }
                // Filter by assignee
                if (assigneeFilters.value.size > 0) {
                    // If node has no assignee but filters are active, hide it
                    if (!node.assigneeId) {
                        return false;
                    }
                    // If node has assignee but it's not in the filter, hide it
                    if (!assigneeFilters.value.has(node.assigneeId)) {
                        return false;
                    }
                }
                return true;
            });

            // Process data with better initial positioning
            nodes = filteredNodes.map((d, i) => ({
                ...d,
                x: width / 2 + ((i % 3) - 1) * 100,
                y: height / 2 + Math.floor(i / 3) * 80,
            }));

            // Create a set of valid node IDs for quick lookup (filtered nodes only)
            const validNodeIds = new Set(nodes.map((n) => n.id));

            // Filter out links that reference non-existent nodes OR epic links OR filtered nodes
            const validEdges = props.graphData.edges.filter((edge) => {
                // Check if source and target exist in filtered nodes
                const sourceExists = validNodeIds.has(edge.from);
                const targetExists = validNodeIds.has(edge.to);

                if (!sourceExists || !targetExists) {
                    // Only log if nodes don't exist in original data (not just filtered out)
                    const sourceInOriginal = props.graphData.nodes.some(
                        (n) => n.id === edge.from,
                    );
                    const targetInOriginal = props.graphData.nodes.some(
                        (n) => n.id === edge.to,
                    );

                    if (!sourceInOriginal || !targetInOriginal) {
                        console.warn(
                            "Filtering out edge with missing node(s):",
                            {
                                from: edge.from,
                                to: edge.to,
                                sourceExists: sourceInOriginal,
                                targetExists: targetInOriginal,
                                edgeType: edge.type,
                            },
                        );
                    }
                    return false;
                }

                // Filter out ALL epic-related links at the data level
                if (edge.type === "epic link") {
                    console.log("Filtering out epic link:", {
                        from: edge.from,
                        to: edge.to,
                        type: edge.type,
                    });
                    return false;
                }

                // Also filter out links involving epic nodes
                const sourceNode = nodes.find((n) => n.id === edge.from);
                const targetNode = nodes.find((n) => n.id === edge.to);

                if (
                    sourceNode?.status === "Epic" ||
                    targetNode?.status === "Epic"
                ) {
                    console.log("Filtering out link involving epic node:", {
                        from: edge.from,
                        to: edge.to,
                        sourceIsEpic: sourceNode?.status === "Epic",
                        targetIsEpic: targetNode?.status === "Epic",
                    });
                    return false;
                }

                return true;
            });

            links = validEdges.map((d) => ({
                ...d,
                source: d.from,
                target: d.to,
            }));

            console.log(
                `Processed ${props.graphData.edges.length} edges, kept ${validEdges.length} valid edges`,
            );

            // Create links
            linkElements = g
                .selectAll(".link")
                .data(links)
                .enter()
                .append("path")
                .attr("class", "link")
                .attr("data-type", (d) => d.type)
                .attr("stroke", (d) =>
                    getLineColorByStatus(d.source.id || d.source),
                )
                .attr("stroke-width", 2.5)
                .attr("stroke-dasharray", (d) =>
                    d.type === "epic link" ? "5,5" : null,
                )
                .attr("fill", "none")
                .attr("stroke-linecap", "round")
                .attr("stroke-linejoin", "round")
                .attr("opacity", 0.8)
                .attr("marker-end", (d) =>
                    d.type === "Blocks"
                        ? "url(#arrow-blocks)"
                        : "url(#arrow-epic)",
                )
                .attr("d", "M-9999,-9999 L-9999,-9999")
                .style("filter", "drop-shadow(0 1px 2px rgba(0,0,0,0.1))")
                .on("mouseover", function (event, d) {
                    // Highlight the connection line
                    d3.select(this)
                        .transition()
                        .duration(150)
                        .attr("stroke-width", 5)
                        .attr("opacity", 1)
                        .style(
                            "filter",
                            "drop-shadow(0 3px 8px rgba(0,0,0,0.3))",
                        )
                        .style(
                            "stroke-dasharray",
                            d.type === "epic link" ? "8,8" : null,
                        );

                    // Highlight connected nodes
                    const sourceNode =
                        typeof d.source === "string"
                            ? nodes.find((n) => n.id === d.source)
                            : d.source;
                    const targetNode =
                        typeof d.target === "string"
                            ? nodes.find((n) => n.id === d.target)
                            : d.target;

                    if (sourceNode && targetNode) {
                        // Add highlight class to connected nodes
                        nodeElements
                            .filter(
                                (nodeData) =>
                                    nodeData.id === sourceNode.id ||
                                    nodeData.id === targetNode.id,
                            )
                            .classed("node-highlighted", true)
                            .selectAll(".card-background")
                            .transition()
                            .duration(150)
                            .attr("stroke-width", 4)
                            .style(
                                "filter",
                                "drop-shadow(0 4px 12px rgba(0,0,0,0.3))",
                            );

                        // Dim other nodes
                        nodeElements
                            .filter(
                                (nodeData) =>
                                    nodeData.id !== sourceNode.id &&
                                    nodeData.id !== targetNode.id,
                            )
                            .transition()
                            .duration(150)
                            .style("opacity", 0.4);

                        // Dim other links
                        linkElements
                            .filter((linkData) => linkData !== d)
                            .transition()
                            .duration(150)
                            .style("opacity", 0.2);
                    }
                })
                .on("mouseout", function (event, d) {
                    // Reset connection line
                    d3.select(this)
                        .transition()
                        .duration(150)
                        .attr("stroke-width", 2.5)
                        .attr("opacity", 0.8)
                        .style(
                            "filter",
                            "drop-shadow(0 1px 2px rgba(0,0,0,0.1))",
                        )
                        .style(
                            "stroke-dasharray",
                            d.type === "epic link" ? "5,5" : null,
                        );

                    // Reset all nodes
                    nodeElements
                        .classed("node-highlighted", false)
                        .transition()
                        .duration(150)
                        .style("opacity", 1)
                        .selectAll(".card-background")
                        .transition()
                        .duration(150)
                        .attr("stroke-width", 3)
                        .style("filter", null);

                    // Reset all links
                    linkElements
                        .transition()
                        .duration(150)
                        .style("opacity", 0.8);
                });

            // Ensure proper attributes are set immediately
            if (linkElements) {
                linkElements.attr("marker-end", (d) =>
                    d.type === "Blocks"
                        ? "url(#arrow-blocks)"
                        : "url(#arrow-epic)",
                );

                // Also ensure stroke attributes are properly set
                linkElements
                    .attr("stroke", (d) =>
                        getLineColorByStatus(d.source.id || d.source),
                    )
                    .attr("stroke-width", 2);
            }

            // Create nodes
            nodeElements = g
                .selectAll(".node")
                .data(nodes)
                .enter()
                .append("g")
                .attr("class", "node")
                .style("cursor", "pointer")
                .call(
                    d3
                        .drag()
                        .on("start", dragstarted)
                        .on("drag", dragged)
                        .on("end", dragended),
                );

            // Create card containers
            const cardElements = nodeElements
                .append("g")
                .attr("class", "node-card");

            // Add card background rectangles
            cardElements
                .append("rect")
                .attr("class", "card-background")
                .attr("width", (d) => (d.status === "Epic" ? 160 : 140))
                .attr("height", (d) => (d.status === "Epic" ? 90 : 80))
                .attr("x", (d) => (d.status === "Epic" ? -80 : -70))
                .attr("y", (d) => (d.status === "Epic" ? -45 : -40))
                .attr("rx", 8)
                .attr("ry", 8)
                .attr("fill", "#ffffff")
                .attr("stroke", (d) => {
                    // Border color by classification
                    const classification =
                        d.classification?.toLowerCase() || "new";
                    const style = getComputedStyle(document.documentElement);

                    if (classification === "epic") {
                        return style.getPropertyValue("--ctp-mauve").trim();
                    } else if (classification === "done") {
                        return style.getPropertyValue("--ctp-green").trim();
                    } else if (classification === "indeterminate") {
                        return style.getPropertyValue("--ctp-blue").trim();
                    } else if (classification === "new") {
                        return style.getPropertyValue("--ctp-overlay1").trim();
                    }

                    return (
                        style.getPropertyValue("--text-muted").trim() ||
                        "#6b7280"
                    );
                })
                .attr("stroke-width", 3);

            // Add status indicator bar at top of card
            cardElements
                .append("rect")
                .attr("class", "status-indicator")
                .attr("width", (d) => (d.status === "Epic" ? 160 : 140))
                .attr("height", 6)
                .attr("x", (d) => (d.status === "Epic" ? -80 : -70))
                .attr("y", (d) => (d.status === "Epic" ? -45 : -40))
                .attr("rx", 6)
                .attr("ry", 6)
                .attr("fill", (d) => {
                    const classification =
                        d.classification?.toLowerCase() || "new";
                    const style = getComputedStyle(document.documentElement);

                    if (classification === "epic") {
                        return style.getPropertyValue("--ctp-mauve").trim();
                    } else if (classification === "done") {
                        return style.getPropertyValue("--ctp-green").trim();
                    } else if (classification === "indeterminate") {
                        return style.getPropertyValue("--ctp-blue").trim();
                    } else if (classification === "new") {
                        return style.getPropertyValue("--ctp-overlay1").trim();
                    }

                    return (
                        style.getPropertyValue("--text-muted").trim() ||
                        "#6b7280"
                    );
                });

            // Add ticket ID (bold, top)
            // Add ticket id
            cardElements
                .append("text")
                .attr("class", "card-id")
                .attr("x", 0)
                .attr("y", (d) => (d.status === "Epic" ? -25 : -20))
                .attr("text-anchor", "middle")
                .attr("font-size", "12px")
                .attr("font-weight", "bold")
                .attr("fill", "#2d3748")
                .text((d) => d.id);

            // Add story points badge if available
            const pointsGroup = cardElements
                .filter((d) => d.points || d.storyPoints || d.story_points)
                .append("g")
                .attr("class", "card-points-badge")
                .attr("transform", (d) =>
                    d.status === "Epic"
                        ? "translate(60, -35)"
                        : "translate(50, -30)",
                );

            // Add circular background for points badge
            pointsGroup
                .append("circle")
                .attr("r", 10)
                .attr("fill", "var(--primary-color)")
                .attr("stroke", "var(--card-bg)")
                .attr("stroke-width", "2")
                .attr("opacity", "0.95");

            // Add points text
            pointsGroup
                .append("text")
                .attr("class", "card-points-text")
                .attr("text-anchor", "middle")
                .attr("dominant-baseline", "central")
                .attr("font-size", "9px")
                .attr("font-weight", "bold")
                .attr("fill", "#ffffff")
                .text((d) => {
                    const points = d.points || d.storyPoints || d.story_points;
                    return points || "";
                });

            // Add ticket title (wrapped text)
            cardElements
                .append("text")
                .attr("class", "card-title")
                .attr("x", 0)
                .attr("y", (d) => (d.status === "Epic" ? -7 : -3))
                .attr("text-anchor", "middle")
                .attr("font-size", "10px")
                .attr("font-weight", "normal")
                .attr("fill", "#4a5568")
                .each(function (d) {
                    const text = d3.select(this);
                    const words = d.label.split(/\s+/);
                    const maxWidth = d.status === "Epic" ? 150 : 130;
                    const lineHeight = 12;
                    const maxLines = 2;

                    text.text(null);

                    let line = [];
                    let lineNumber = 0;
                    let wordIndex = 0;

                    while (wordIndex < words.length && lineNumber < maxLines) {
                        line = [];

                        // Build line word by word until it exceeds maxWidth
                        while (wordIndex < words.length) {
                            const testLine = [...line, words[wordIndex]];

                            // Create temporary text element to measure width
                            const tempText = text
                                .append("tspan")
                                .text(testLine.join(" "));
                            const textLength = tempText
                                .node()
                                .getComputedTextLength();
                            tempText.remove();

                            if (textLength > maxWidth && line.length > 0) {
                                // Current line is full, break to create tspan
                                break;
                            } else {
                                // Add word to current line
                                line.push(words[wordIndex]);
                                wordIndex++;
                            }
                        }

                        // Create tspan for current line
                        if (line.length > 0) {
                            const isLastLine = lineNumber === maxLines - 1;
                            const hasMoreWords = wordIndex < words.length;

                            let lineText = line.join(" ");
                            if (isLastLine && hasMoreWords) {
                                lineText += "...";
                            }

                            text.append("tspan")
                                .attr("x", 0)
                                .attr("dy", lineNumber === 0 ? 0 : lineHeight)
                                .text(lineText);

                            lineNumber++;

                            // If this was the last line and we have more words, stop
                            if (isLastLine && hasMoreWords) {
                                break;
                            }
                        }
                    }
                });

            // Add status text (bottom)
            cardElements
                .append("text")
                .attr("class", "card-status")
                .attr("x", 0)
                .attr("y", (d) => (d.status === "Epic" ? 35 : 30))
                .attr("text-anchor", "middle")
                .attr("font-size", "9px")
                .attr("font-weight", "600")
                .attr("fill", (d) => {
                    const classification =
                        d.classification?.toLowerCase() || "new";
                    const style = getComputedStyle(document.documentElement);

                    if (classification === "epic") {
                        return style.getPropertyValue("--ctp-mauve").trim();
                    } else if (classification === "done") {
                        return style.getPropertyValue("--ctp-green").trim();
                    } else if (classification === "indeterminate") {
                        return style.getPropertyValue("--ctp-blue").trim();
                    } else if (classification === "new") {
                        return style.getPropertyValue("--ctp-overlay1").trim();
                    }

                    return (
                        style.getPropertyValue("--text-muted").trim() ||
                        "#6b7280"
                    );
                })
                .text((d) => d.status);

            // Add assignee text (top-right corner, very subtle)
            cardElements
                .filter((d) => d.assignee)
                .append("text")
                .attr("class", "card-assignee")
                .attr("x", (d) => (d.status === "Epic" ? 70 : 60))
                .attr("y", (d) => (d.status === "Epic" ? -35 : -30))
                .attr("text-anchor", "end")
                .attr("font-size", "6px")
                .attr("font-weight", "300")
                .attr("fill", "#cbd5e1")
                .attr("opacity", "0.6")
                .text((d) => d.assignee.split(" ")[0].substring(0, 8)); // First name, max 8 chars

            // Add click handler for nodes
            nodeElements.on("click", (event, d) => {
                selectedNode.value = d;
                nextTick(() => {
                    if (nodeDialog.value) {
                        nodeDialog.value.showModal();
                        // Focus the close button for keyboard accessibility
                        const closeBtn =
                            nodeDialog.value.querySelector(".close-dialog");
                        if (closeBtn) {
                            closeBtn.focus();
                        }
                    }
                });
            });

            // Add hover effects to nodes
            nodeElements
                .on("mouseover", function (event, d) {
                    const currentNode = d;

                    // Highlight current node
                    d3.select(this)
                        .classed("node-highlighted", true)
                        .selectAll(".card-background")
                        .transition()
                        .duration(150)
                        .attr("stroke-width", 4)
                        .style(
                            "filter",
                            "drop-shadow(0 4px 12px rgba(0,0,0,0.3))",
                        );

                    // Find and highlight connected links and nodes
                    const connectedNodeIds = new Set();

                    linkElements.each(function (linkData) {
                        const sourceId =
                            typeof linkData.source === "string"
                                ? linkData.source
                                : linkData.source.id;
                        const targetId =
                            typeof linkData.target === "string"
                                ? linkData.target
                                : linkData.target.id;

                        // Skip epic-related links entirely in phase layout
                        if (layoutType.value === "phase") {
                            const sourceNode = nodes.find(
                                (n) => n.id === sourceId,
                            );
                            const targetNode = nodes.find(
                                (n) => n.id === targetId,
                            );
                            const isSourceEpic =
                                sourceNode && sourceNode.status === "Epic";
                            const isTargetEpic =
                                targetNode && targetNode.status === "Epic";

                            if (isSourceEpic || isTargetEpic) {
                                return; // Skip this link entirely
                            }
                        }

                        if (
                            sourceId === currentNode.id ||
                            targetId === currentNode.id
                        ) {
                            // Highlight connected link
                            d3.select(this)
                                .classed("link-highlighted", true)
                                .transition()
                                .duration(150)
                                .attr("stroke-width", 4)
                                .attr("opacity", 1)
                                .style(
                                    "filter",
                                    "drop-shadow(0 3px 8px rgba(0,0,0,0.3))",
                                );

                            // Track connected nodes
                            if (sourceId === currentNode.id) {
                                connectedNodeIds.add(targetId);
                            } else {
                                connectedNodeIds.add(sourceId);
                            }
                        } else {
                            // Dim unconnected links
                            d3.select(this)
                                .transition()
                                .duration(150)
                                .style("opacity", 0.2);
                        }
                    });

                    // Highlight connected nodes and dim others
                    nodeElements.each(function (nodeData) {
                        if (nodeData.id !== currentNode.id) {
                            if (connectedNodeIds.has(nodeData.id)) {
                                // Highlight connected nodes
                                d3.select(this)
                                    .classed("node-connected", true)
                                    .selectAll(".card-background")
                                    .transition()
                                    .duration(150)
                                    .attr("stroke-width", 4)
                                    .style(
                                        "filter",
                                        "drop-shadow(0 2px 6px rgba(0,0,0,0.2))",
                                    );
                            } else {
                                // Dim unconnected nodes
                                d3.select(this)
                                    .transition()
                                    .duration(150)
                                    .style("opacity", 0.3);
                            }
                        }
                    });
                })
                .on("mouseout", function (event, d) {
                    // Reset all nodes
                    nodeElements
                        .classed("node-highlighted", false)
                        .classed("node-connected", false)
                        .transition()
                        .duration(150)
                        .style("opacity", 1)
                        .selectAll(".card-background")
                        .transition()
                        .duration(150)
                        .attr("stroke-width", 3)
                        .style("filter", null);

                    // Reset all links
                    linkElements
                        .classed("link-highlighted", false)
                        .transition()
                        .duration(150)
                        .attr("stroke-width", 2.5)
                        .style("opacity", 0.8)
                        .style(
                            "filter",
                            "drop-shadow(0 1px 2px rgba(0,0,0,0.1))",
                        );
                });

            // Always calculate phases for debugging/analysis purposes
            const phaseData = calculatePhases();

            // Initialize simulation
            if (layoutType.value === "force") {
                createForceSimulation();
            } else if (layoutType.value === "tree") {
                createTreeLayout();
            } else {
                createPhaseLayout();
            }

            // Ensure epic tickets are hidden in phase layout from the start
            if (layoutType.value === "phase") {
                nodeElements.style("display", (d) => {
                    const isEpic = d.status === "Epic";
                    return isEpic ? "none" : "block";
                });

                linkElements.each(function (d) {
                    const sourceIsEpic =
                        d.source.status && d.source.status === "Epic";
                    const targetIsEpic =
                        d.target.status && d.target.status === "Epic";

                    if (sourceIsEpic || targetIsEpic) {
                        d3.select(this)
                            .style("display", "none")
                            .style("visibility", "hidden")
                            .style("opacity", 0);
                    } else {
                        d3.select(this)
                            .style("display", "block")
                            .style("visibility", "visible")
                            .style("opacity", 1);
                    }
                });
            }
        };

        const createForceSimulation = () => {
            // Ensure markers are available before working with links
            ensureMarkers();

            // Create a copy of links for force simulation to avoid modifying the global array
            const forceLinks = links.map((d) => ({ ...d }));

            simulation = d3
                .forceSimulation(nodes)
                .force(
                    "link",
                    d3
                        .forceLink(forceLinks)
                        .id((d) => d.id)
                        .distance(180),
                )
                .force("charge", d3.forceManyBody().strength(-500))
                .force(
                    "center",
                    d3.forceCenter(
                        graphContainer.value.clientWidth / 2,
                        graphContainer.value.clientHeight / 2,
                    ),
                )
                .force(
                    "collision",
                    d3
                        .forceCollide()
                        .radius((d) => (d.status === "Epic" ? 90 : 80)),
                );

            simulation.on("tick", () => {
                // Only update links if we're actually in force layout mode
                if (layoutType.value === "force") {
                    linkElements.each(function (d) {
                        // In force simulation, d.source and d.target should already be node objects
                        // but add safety check just in case
                        const sourceNode =
                            typeof d.source === "string"
                                ? nodes.find((n) => n.id === d.source)
                                : d.source;
                        const targetNode =
                            typeof d.target === "string"
                                ? nodes.find((n) => n.id === d.target)
                                : d.target;

                        const connections = getCardConnectionPoint(
                            sourceNode,
                            targetNode,
                        );
                        if (
                            connections &&
                            connections.source &&
                            connections.target &&
                            connections.path
                        ) {
                            d3.select(this).attr("d", connections.path);
                        }
                    });

                    nodeElements.attr(
                        "transform",
                        (d) => `translate(${d.x},${d.y})`,
                    );
                } else {
                    console.warn(
                        "Force simulation tick fired but layout is:",
                        layoutType.value,
                    );
                }
            });

            // Force initial link positioning for first render - immediate
            linkElements.each(function (d) {
                // Resolve source and target nodes from IDs if needed
                const sourceNode =
                    typeof d.source === "string"
                        ? nodes.find((n) => n.id === d.source)
                        : d.source;
                const targetNode =
                    typeof d.target === "string"
                        ? nodes.find((n) => n.id === d.target)
                        : d.target;

                const connections = getCardConnectionPoint(
                    sourceNode,
                    targetNode,
                );
                if (
                    connections &&
                    connections.source &&
                    connections.target &&
                    connections.path
                ) {
                    d3.select(this).attr("d", connections.path);
                }
            });
        };

        const createTreeLayout = () => {
            if (simulation) {
                simulation.stop();
                simulation = null;
            }

            // Stop any running transitions
            linkElements.interrupt();
            nodeElements.interrupt();

            // Ensure markers are available before working with links
            ensureMarkers();

            const root = d3.hierarchy(buildHierarchy());
            const treeLayout = d3
                .tree()
                .size([
                    graphContainer.value.clientWidth - 100,
                    graphContainer.value.clientHeight - 100,
                ]);

            treeLayout(root);

            // Update node positions
            root.descendants().forEach((d) => {
                const node = nodes.find((n) => n.id === d.data.id);
                if (node) {
                    node.x = d.x + 50;
                    node.y = d.y + 50;
                    node.fx = d.x + 50;
                    node.fy = d.y + 50;
                }
            });

            // Update visual elements
            nodeElements
                .transition()
                .duration(1000)
                .attr("transform", (d) => `translate(${d.x},${d.y})`);

            linkElements
                .transition()
                .duration(1000)
                .attrTween("d", function (d) {
                    const currentPath = this.getAttribute("d") || "M0,0 L0,0";
                    const targetPath = getSafeConnectionPoint(d, "path", false);
                    return d3.interpolate(currentPath, targetPath);
                });
        };

        const createPhaseLayout = () => {
            if (simulation) {
                simulation.stop();
                simulation = null;
            }

            // Stop any running transitions
            linkElements.interrupt();
            nodeElements.interrupt();

            // Ensure markers are available before working with links
            ensureMarkers();

            // IMMEDIATELY hide all epic nodes and epic links before any other processing
            nodeElements.style("display", (d) => {
                const isEpic = d.status === "Epic";
                return isEpic ? "none" : "block";
            });

            linkElements.each(function (d) {
                // Check if this is an epic link by type
                if (d.type === "epic link") {
                    d3.select(this)
                        .style("display", "none")
                        .style("visibility", "hidden")
                        .style("opacity", 0);
                    return;
                }

                // Check if this involves epic nodes
                const sourceIsEpic = d.source?.status === "Epic";
                const targetIsEpic = d.target?.status === "Epic";

                if (sourceIsEpic || targetIsEpic) {
                    d3.select(this)
                        .style("display", "none")
                        .style("visibility", "hidden")
                        .style("opacity", 0);
                }
            });

            const { nodeMap, phases } = calculatePhases();

            // Update phase count for display
            phaseCount.value = phases.length;

            const width = graphContainer.value.clientWidth - 100;
            const height = graphContainer.value.clientHeight - 100;

            // Calculate column width and spacing with increased spacing between phases
            const numPhases = phases.length;
            const minColumnWidth = 200; // Minimum width per column
            const phaseGap = 60; // Gap between phases
            const totalGapWidth = (numPhases - 1) * phaseGap;
            const availableWidth = width - totalGapWidth;
            const columnWidth = Math.max(
                minColumnWidth,
                availableWidth / Math.max(numPhases, 1),
            );
            const padding = 50;

            // Calculate maximum height needed for all phases
            let maxPhaseHeight = 0;
            phases.forEach((phaseNodes) => {
                const workNodesInPhase = phaseNodes.filter((nodeId) => {
                    const node = nodes.find((n) => n.id === nodeId);
                    if (!node) return false;
                    const isEpic =
                        node.id.includes("Epic") ||
                        node.label === "Epic" ||
                        node.classification === "epic";
                    return !isEpic;
                });

                const cardHeight = 80;
                const minGap = 8;
                const minSpacingPerCard = cardHeight + minGap;
                const phaseHeight = workNodesInPhase.length * minSpacingPerCard;
                maxPhaseHeight = Math.max(maxPhaseHeight, phaseHeight);
            });

            // Add extra padding for labels and margins
            const totalPhaseHeight = Math.max(maxPhaseHeight + 200, height);

            // Add phase background columns FIRST (behind everything)
            const phaseColumns = svg
                .select(".graph-group")
                .selectAll(".phase-column")
                .data(phases)
                .enter()
                .insert("rect", ":first-child")
                .attr("class", "phase-column")
                .attr("x", (d, i) => i * (columnWidth + phaseGap))
                .attr("y", 0)
                .attr("width", columnWidth - 10)
                .attr("height", totalPhaseHeight)
                .attr("fill", "var(--hover-bg)")
                .attr("opacity", 0.3)
                .attr("rx", 8)
                .style("pointer-events", "none");

            // Add phase labels on top
            const phaseLabels = svg
                .select(".graph-group")
                .selectAll(".phase-label")
                .data(phases)
                .enter()
                .append("text")
                .attr("class", "phase-label")
                .attr(
                    "x",
                    (d, i) => i * (columnWidth + phaseGap) + columnWidth / 2,
                )
                .attr("y", 30)
                .attr("text-anchor", "middle")
                .attr("font-size", "14px")
                .attr("font-weight", "bold")
                .attr("fill", "var(--text-primary)")
                .text((d, i) => `Phase ${i + 1}`);

            // Position nodes in phases
            let totalPositioned = 0;
            let missingNodes = [];

            phases.forEach((phaseNodes, phaseIndex) => {
                const x =
                    phaseIndex * (columnWidth + phaseGap) + columnWidth / 2;

                // Filter out epic nodes first to get accurate count
                const workNodesInPhase = phaseNodes.filter((nodeId) => {
                    const node = nodes.find((n) => n.id === nodeId);
                    if (!node) return false;
                    const isEpic =
                        node.id.includes("Epic") ||
                        node.label === "Epic" ||
                        node.classification === "epic";
                    return !isEpic;
                });

                // Calculate spacing based on actual work nodes count
                // Card height (80px) + minimum gap (8px) = 88px minimum per card
                const cardHeight = 80;
                const minGap = 8;
                const minSpacingPerCard = cardHeight + minGap;
                const availableHeight = totalPhaseHeight - 200; // Leave space for labels and padding

                // Ensure we have at least the minimum spacing, but allow more if there's room
                const idealSpacing =
                    availableHeight / Math.max(workNodesInPhase.length, 1);
                const calculatedSpacing = Math.max(
                    minSpacingPerCard,
                    idealSpacing,
                );

                // Use calculated spacing to allow proper distribution
                const finalSpacing = calculatedSpacing;

                const startY = padding + 80;
                let workNodeIndex = 0; // Track only work nodes for positioning

                phaseNodes.forEach((nodeId) => {
                    const node = nodes.find((n) => n.id === nodeId);
                    if (node) {
                        // Skip epic nodes in phase layout
                        const isEpic = node.status === "Epic";

                        if (!isEpic) {
                            // Center cards in column with slight offset for visual variety
                            const offsetRange = Math.min(20, columnWidth / 8);
                            node.x = x + (Math.random() - 0.5) * offsetRange;
                            node.y = startY + workNodeIndex * finalSpacing;
                            node.fx = node.x;
                            node.fy = node.y;
                            totalPositioned++;
                            workNodeIndex++; // Only increment for work nodes
                        }
                    } else {
                        missingNodes.push(nodeId);
                        console.warn(
                            `Node ${nodeId} not found in nodes array for phase ${phaseIndex + 1}`,
                        );
                    }
                });
            });

            // Position links immediately for phase layout using side connections
            linkElements.each(function (d, i) {
                // FIRST: Check if this is an epic link or involves epic nodes - hide immediately
                if (d.type === "epic link") {
                    d3.select(this)
                        .style("display", "none")
                        .style("visibility", "hidden")
                        .style("opacity", 0);
                    return;
                }

                // Resolve source and target nodes from IDs if needed
                const sourceNode =
                    typeof d.source === "string"
                        ? nodes.find((n) => n.id === d.source)
                        : d.source;
                const targetNode =
                    typeof d.target === "string"
                        ? nodes.find((n) => n.id === d.target)
                        : d.target;

                // Filter out epic relations in phase layout
                if (sourceNode && targetNode) {
                    const isSourceEpic = sourceNode.status === "Epic";
                    const isTargetEpic = targetNode.status === "Epic";

                    // Hide ALL epic-related links in phase layout
                    if (isSourceEpic || isTargetEpic) {
                        d3.select(this)
                            .style("display", "none")
                            .style("visibility", "hidden")
                            .style("opacity", 0);
                        return;
                    }
                }

                const connections = getPhaseConnectionPoint(
                    sourceNode,
                    targetNode,
                );

                if (connections === null) {
                    // Node not found - hide this link completely
                    d3.select(this)
                        .style("display", "none")
                        .style("visibility", "hidden")
                        .style("opacity", 0);
                    return;
                }

                if (
                    connections &&
                    connections.source &&
                    connections.target &&
                    connections.path
                ) {
                    d3.select(this)
                        .attr("d", connections.path)
                        .style("display", "block")
                        .style("visibility", "visible")
                        .style("opacity", 1)
                        .attr("stroke", (d) =>
                            getLineColorByStatus(d.source.id || d.source),
                        )
                        .attr("stroke-width", 2);
                } else {
                    console.warn("Invalid connections for link:", d);
                }
            });

            // Add safeguard to maintain link positions and enforce epic hiding
            const maintainPhaseLinks = () => {
                if (layoutType.value !== "phase") return;

                linkElements.each(function (d, i) {
                    const elem = d3.select(this);

                    // ALWAYS check for epic links first - regardless of path state
                    if (d.type === "epic link") {
                        elem.style("display", "none")
                            .style("visibility", "hidden")
                            .style("opacity", 0);
                        return;
                    }

                    const sourceNode =
                        typeof d.source === "string"
                            ? nodes.find((n) => n.id === d.source)
                            : d.source;
                    const targetNode =
                        typeof d.target === "string"
                            ? nodes.find((n) => n.id === d.target)
                            : d.target;

                    if (sourceNode && targetNode) {
                        const isSourceEpic = sourceNode.status === "Epic";
                        const isTargetEpic = targetNode.status === "Epic";

                        // ALWAYS hide epic-related links
                        if (isSourceEpic || isTargetEpic) {
                            elem.style("display", "none")
                                .style("visibility", "hidden")
                                .style("opacity", 0);
                            return;
                        }
                    }

                    // Check if path was reset or is invalid for non-epic links
                    const currentPath = elem.attr("d");
                    if (
                        !currentPath ||
                        currentPath.includes("-9999") ||
                        currentPath === "M0,0 L0,0" ||
                        currentPath === "M-9999,-9999 L-9999,-9999"
                    ) {
                        if (sourceNode && targetNode) {
                            const connections = getPhaseConnectionPoint(
                                sourceNode,
                                targetNode,
                            );
                            if (connections === null) {
                                // Node not found - hide this link
                                elem.style("display", "none")
                                    .style("visibility", "hidden")
                                    .style("opacity", 0);
                            } else if (
                                connections &&
                                connections.source &&
                                connections.target &&
                                connections.path
                            ) {
                                elem.attr("d", connections.path)
                                    .style("display", "block")
                                    .style("visibility", "visible")
                                    .style("opacity", 1);
                            }
                        }
                    }
                });
            };

            // Monitor and maintain every 50ms for first 5 seconds (more frequent and longer)
            const monitorInterval = setInterval(maintainPhaseLinks, 50);
            setTimeout(() => clearInterval(monitorInterval), 5000);

            // Completely hide epic nodes in phase layout
            nodeElements.style("display", (d) => {
                const isEpic = d.status === "Epic";
                return isEpic ? "none" : "block";
            });

            // Completely hide epic links in phase layout and ensure markers are properly applied
            linkElements.each(function (d) {
                const sourceIsEpic =
                    d.source.status && d.source.status === "Epic";
                const targetIsEpic =
                    d.target.status && d.target.status === "Epic";

                if (sourceIsEpic || targetIsEpic) {
                    d3.select(this)
                        .style("display", "none")
                        .style("visibility", "hidden")
                        .style("opacity", 0);
                } else {
                    d3.select(this)
                        .style("display", "block")
                        .style("visibility", "visible")
                        .style("opacity", 1)
                        .attr(
                            "marker-end",
                            d.type === "Blocks"
                                ? "url(#arrow-blocks)"
                                : "url(#arrow-epic)",
                        );
                }
            });

            // Log positioning results for debugging
            if (missingNodes.length > 0) {
                console.warn("Missing nodes:", missingNodes);
                console.log(
                    "Available nodes:",
                    nodes.map((n) => n.id),
                );
                console.log(
                    "Available node details:",
                    nodes.map((n) => ({
                        id: n.id,
                        label: n.label,
                        classification: n.classification,
                    })),
                );
            }

            // Verify no epic nodes are in phases
            const epicInPhases = phases.flat().filter((nodeId) => {
                const node = props.graphData.nodes.find((n) => n.id === nodeId);
                return node && node.status === "Epic";
            });

            console.log(
                "Epic nodes in phases (should be empty):",
                epicInPhases,
            );

            // Update visual elements with animation
            nodeElements
                .transition()
                .duration(1000)
                .attr("transform", (d) => `translate(${d.x},${d.y})`);

            linkElements
                .transition()
                .duration(1000)
                .attrTween("d", function (d) {
                    const currentPath = this.getAttribute("d") || "M0,0 L0,0";
                    const targetPath = getSafeConnectionPoint(d, "path", true);

                    // If targetPath is null (missing node), hide the link
                    if (targetPath === null || targetPath === "M0,0 L0,0") {
                        d3.select(this)
                            .style("display", "none")
                            .style("visibility", "hidden")
                            .style("opacity", 0);
                        return d3.interpolate(currentPath, currentPath);
                    }

                    return d3.interpolate(currentPath, targetPath);
                })
                .on("end", function () {
                    // Comprehensive epic link cleanup after animation
                    linkElements.each(function (d) {
                        const elem = d3.select(this);

                        // First check: epic link type
                        if (d.type === "epic link") {
                            elem.style("display", "none")
                                .style("visibility", "hidden")
                                .style("opacity", 0);
                            return;
                        }

                        // Second check: nodes involving epics
                        const sourceIsEpic = d.source?.status === "Epic";
                        const targetIsEpic = d.target?.status === "Epic";

                        if (sourceIsEpic || targetIsEpic) {
                            elem.style("display", "none")
                                .style("visibility", "hidden")
                                .style("opacity", 0);
                            return;
                        }

                        // Only process valid work-to-work links
                        const connections = getPhaseConnectionPoint(
                            d.source,
                            d.target,
                        );
                        if (
                            connections &&
                            connections.source &&
                            connections.target
                        ) {
                            elem.attr("x1", connections.source.x || 0)
                                .attr("y1", connections.source.y || 0)
                                .attr("x2", connections.target.x || 0)
                                .attr("y2", connections.target.y || 0)
                                .attr("marker-end", "url(#arrow-blocks)")
                                .style("display", "block")
                                .style("visibility", "visible")
                                .style("opacity", 1);
                        }
                    });

                    // Final cleanup sweep - run after a short delay to catch any stragglers
                    setTimeout(() => {
                        linkElements.each(function (d) {
                            if (
                                d.type === "epic link" ||
                                d.source?.status === "Epic" ||
                                d.target?.status === "Epic"
                            ) {
                                d3.select(this)
                                    .style("display", "none")
                                    .style("visibility", "hidden")
                                    .style("opacity", 0);
                            }
                        });
                    }, 100);
                });
        };

        const buildHierarchy = () => {
            const epicNode = nodes.find((n) => n.status === "Epic");
            if (!epicNode)
                return {
                    id: "root",
                    children: nodes.map((n) => ({ id: n.id })),
                };

            const children = nodes
                .filter((n) => n.id !== epicNode.id)
                .map((n) => ({ id: n.id }));
            return { id: epicNode.id, children };
        };

        const dragstarted = (event, d) => {
            if (!event.active && simulation)
                simulation.alphaTarget(0.3).restart();
            d.fx = d.x;
            d.fy = d.y;
        };

        const dragged = (event, d) => {
            d.fx = event.x;
            d.fy = event.y;
        };

        const dragended = (event, d) => {
            if (!event.active && simulation) simulation.alphaTarget(0);
            if (layoutType.value === "force") {
                d.fx = null;
                d.fy = null;
            }
        };

        const resetZoom = () => {
            svg.transition()
                .duration(750)
                .call(zoom.transform, d3.zoomIdentity);
        };

        const toggleLayout = () => {
            if (layoutType.value === "phase") {
                layoutType.value = "force";
            } else if (layoutType.value === "force") {
                layoutType.value = "tree";
            } else {
                layoutType.value = "phase";
            }

            // Clear phase-specific elements
            if (svg) {
                svg.selectAll(".phase-label").remove();
                svg.selectAll(".phase-column").remove();
            }

            console.log("Switching to layout:", layoutType.value);

            if (layoutType.value === "force") {
                // Remove fixed positions and restore visibility
                nodes.forEach((d) => {
                    d.fx = null;
                    d.fy = null;
                });
                // Restore all nodes and links visibility
                nodeElements
                    .style("opacity", 1)
                    .style("display", "block")
                    .style("visibility", "visible");
                linkElements
                    .style("opacity", 1)
                    .style("display", "block")
                    .style("visibility", "visible");
                createForceSimulation();
            } else if (layoutType.value === "tree") {
                // Restore all nodes and links visibility for tree layout
                nodeElements
                    .style("opacity", 1)
                    .style("display", "block")
                    .style("visibility", "visible");
                linkElements
                    .style("opacity", 1)
                    .style("display", "block")
                    .style("visibility", "visible");
                createTreeLayout();
            } else {
                createPhaseLayout();
            }
        };

        const handleResize = () => {
            if (!graphContainer.value) return;

            console.log("Handling resize, container dimensions:", {
                width: graphContainer.value.clientWidth,
                height: graphContainer.value.clientHeight,
                isFullscreen: isFullscreen.value,
            });

            nextTick(() => {
                // Force container to recognize new dimensions
                if (svg) {
                    const newWidth = graphContainer.value.clientWidth;
                    const newHeight = graphContainer.value.clientHeight;

                    // Update SVG dimensions immediately
                    svg.attr("width", newWidth).attr("height", newHeight);

                    // Update zoom center
                    if (zoom) {
                        zoom.translateExtent([
                            [0, 0],
                            [newWidth, newHeight],
                        ]);
                    }
                }

                // Reinitialize graph with new dimensions
                initializeGraph();
            });
        };

        const debugPhases = () => {
            if (
                !props.graphData ||
                !props.graphData.nodes ||
                !props.graphData.edges
            ) {
                alert("No graph data available");
                return;
            }

            const { nodeMap, phases } = calculatePhases();

            let debugMessage = `Phase Analysis:\n\n`;
            const workNodes = props.graphData.nodes.filter(
                (node) => !(node.status && node.status === "Epic"),
            );

            debugMessage += `Total nodes: ${props.graphData.nodes.length}\n`;
            debugMessage += `Work nodes (non-epic): ${workNodes.length}\n`;
            debugMessage += `Epic nodes (excluded): ${props.graphData.nodes.length - workNodes.length}\n`;
            debugMessage += `Total edges: ${props.graphData.edges.length}\n`;
            debugMessage += `Blocking relationships: ${props.graphData.edges.filter((l) => l.type === "Blocks").length}\n\n`;

            debugMessage += `Phase breakdown:\n`;
            phases.forEach((phase, i) => {
                debugMessage += `Phase ${i + 1}: ${phase.length} nodes\n`;
                debugMessage += `  - ${phase.join(", ")}\n`;
            });

            debugMessage += `\nBlocking relationships:\n`;
            const blockingLinks = props.graphData.edges.filter(
                (l) => l.type === "Blocks",
            );
            if (blockingLinks.length === 0) {
                debugMessage += "  - No blocking relationships found\n";
            } else {
                blockingLinks.forEach((link) => {
                    debugMessage += `  - ${link.from} blocks ${link.to}\n`;
                });
            }

            debugMessage += `\nDependencies:\n`;
            Array.from(nodeMap.values()).forEach((node) => {
                if (node.dependencies.length > 0) {
                    debugMessage += `  - ${node.id} depends on: ${node.dependencies.join(", ")}\n`;
                }
            });

            alert(debugMessage);
        };

        const toggleFullscreen = async () => {
            try {
                if (!isFullscreen.value) {
                    console.log("Entering fullscreen...");
                    // Enter fullscreen
                    if (graphContainer.value.requestFullscreen) {
                        await graphContainer.value.requestFullscreen();
                    } else if (graphContainer.value.webkitRequestFullscreen) {
                        await graphContainer.value.webkitRequestFullscreen();
                    } else if (graphContainer.value.mozRequestFullScreen) {
                        await graphContainer.value.mozRequestFullScreen();
                    } else if (graphContainer.value.msRequestFullscreen) {
                        await graphContainer.value.msRequestFullscreen();
                    }
                } else {
                    console.log("Exiting fullscreen...");
                    // Exit fullscreen
                    if (document.exitFullscreen) {
                        await document.exitFullscreen();
                    } else if (document.webkitExitFullscreen) {
                        await document.webkitExitFullscreen();
                    } else if (document.mozCancelFullScreen) {
                        await document.mozCancelFullScreen();
                    } else if (document.msExitFullscreen) {
                        await document.msExitFullscreen();
                    }
                }
            } catch (error) {
                console.warn("Fullscreen toggle failed:", error);
            }
        };

        const toggleStatusFilter = (status) => {
            if (statusFilters.value.has(status)) {
                statusFilters.value.delete(status);
            } else {
                statusFilters.value.add(status);
            }
            // Trigger reactivity
            statusFilters.value = new Set(statusFilters.value);
            // Reinitialize graph with new filters
            initializeGraph();
        };

        const toggleAssigneeFilter = (assigneeId) => {
            if (assigneeFilters.value.has(assigneeId)) {
                assigneeFilters.value.delete(assigneeId);
            } else {
                assigneeFilters.value.add(assigneeId);
            }
            // Trigger reactivity
            assigneeFilters.value = new Set(assigneeFilters.value);
            // Reinitialize graph with new filters
            initializeGraph();
        };

        const toggleAssigneeRow = () => {
            showAssigneeRow.value = !showAssigneeRow.value;
        };

        const startResize = (event) => {
            isResizing.value = true;
            const startY = event.clientY;
            const startHeight = containerHeight.value;

            const handleMouseMove = (moveEvent) => {
                const deltaY = moveEvent.clientY - startY;
                const newHeight = Math.max(300, startHeight + deltaY);
                containerHeight.value = newHeight;
            };

            const handleMouseUp = () => {
                isResizing.value = false;
                document.removeEventListener("mousemove", handleMouseMove);
                document.removeEventListener("mouseup", handleMouseUp);
                // Resize graph after resize is complete
                setTimeout(() => {
                    handleResize();
                }, 100);
            };

            document.addEventListener("mousemove", handleMouseMove);
            document.addEventListener("mouseup", handleMouseUp);
        };

        const closeDialog = () => {
            if (nodeDialog.value) {
                nodeDialog.value.close();
            }
            selectedNode.value = null;
            // Return focus to the graph container after closing
            if (graphContainer.value) {
                graphContainer.value.focus();
            }
        };

        const handleDialogClose = () => {
            selectedNode.value = null;
        };

        const handleDialogBackdropClick = (event) => {
            // Close dialog when clicking on the backdrop (outside the dialog content)
            if (event.target === nodeDialog.value) {
                closeDialog();
            }
        };

        const handleKeydown = (event) => {
            if (
                event.key === "Escape" &&
                selectedNode.value &&
                nodeDialog.value?.open
            ) {
                closeDialog();
            }
        };

        onMounted(() => {
            document.addEventListener("keydown", handleKeydown);
        });

        onUnmounted(() => {
            document.removeEventListener("keydown", handleKeydown);
        });

        const toggleFilters = () => {
            showFilters.value = !showFilters.value;
        };

        const copyNodeLink = async (node) => {
            try {
                // Generate Jira URL based on node ID and backend configuration
                const jiraBaseUrl =
                    props.jiraBaseUrl || "https://example.atlassian.net";
                const nodeUrl = `${jiraBaseUrl}/browse/${node.id}`;
                await navigator.clipboard.writeText(nodeUrl);

                // Show temporary feedback
                const copyBtn = document.querySelector(".copy-btn");
                if (copyBtn) {
                    const originalText = copyBtn.textContent;
                    copyBtn.textContent = "✅ Copied!";
                    copyBtn.style.backgroundColor = "var(--success-color)";
                    setTimeout(() => {
                        copyBtn.textContent = originalText;
                        copyBtn.style.backgroundColor = "";
                    }, 2000);
                }
            } catch (error) {
                console.warn("Failed to copy link:", error);
                // Fallback for older browsers
                const jiraBaseUrl =
                    props.jiraBaseUrl || "https://example.atlassian.net";
                const textArea = document.createElement("textarea");
                textArea.value = `${jiraBaseUrl}/browse/${node.id}`;
                document.body.appendChild(textArea);
                textArea.focus();
                textArea.select();
                try {
                    document.execCommand("copy");
                } catch (fallbackError) {
                    console.warn("Fallback copy failed:", fallbackError);
                }
                document.body.removeChild(textArea);
            }
        };

        const openNodeInNewTab = (node) => {
            // Generate Jira URL based on node ID and backend configuration
            const jiraBaseUrl =
                props.jiraBaseUrl || "https://example.atlassian.net";
            const nodeUrl = `${jiraBaseUrl}/browse/${node.id}`;
            window.open(nodeUrl, "_blank", "noopener,noreferrer");
        };

        const handleFullscreenChange = () => {
            isFullscreen.value = !!(
                document.fullscreenElement ||
                document.webkitFullscreenElement ||
                document.mozFullScreenElement ||
                document.msFullscreenElement
            );

            // Wait for fullscreen transition to complete before resizing
            // Multiple delays ensure proper rendering
            setTimeout(() => {
                console.log("Fullscreen change detected, starting resize...");
                handleResize();
            }, 100);

            setTimeout(() => {
                console.log("Second resize attempt...");
                handleResize();
            }, 300);

            setTimeout(() => {
                console.log("Final resize and redraw...");
                // Force complete redraw for fullscreen
                initializeGraph();
            }, 500);
        };

        onMounted(() => {
            initializeGraph();
            window.addEventListener("resize", handleResize);

            // Add fullscreen event listeners
            document.addEventListener(
                "fullscreenchange",
                handleFullscreenChange,
            );
            document.addEventListener(
                "webkitfullscreenchange",
                handleFullscreenChange,
            );
            document.addEventListener(
                "mozfullscreenchange",
                handleFullscreenChange,
            );
            document.addEventListener(
                "MSFullscreenChange",
                handleFullscreenChange,
            );
        });

        onUnmounted(() => {
            if (mutationObserver) {
                mutationObserver.disconnect();
            }
            if (simulation) {
                simulation.stop();
            }

            // Remove fullscreen event listeners
            document.removeEventListener(
                "fullscreenchange",
                handleFullscreenChange,
            );
            document.removeEventListener(
                "webkitfullscreenchange",
                handleFullscreenChange,
            );
            document.removeEventListener(
                "mozfullscreenchange",
                handleFullscreenChange,
            );
            document.removeEventListener(
                "MSFullscreenChange",
                handleFullscreenChange,
            );
        });

        watch(
            () => props.graphData,
            (newData, oldData) => {
                console.log("Props watcher fired:", {
                    newData: !!newData,
                    oldData: !!oldData,
                    same: newData === oldData,
                });
                // Only reinitialize if the data actually changed (not just D3 modifications)
                if (newData !== oldData) {
                    console.log("Data changed, reinitializing graph");
                    initializeGraph();
                }
            },
            { deep: false },
        );

        return {
            graphContainer,
            nodeDialog,
            selectedNode,
            layoutType,
            phaseCount,
            debugInfo,
            isFullscreen,
            containerHeight,
            isResizing,
            statusFilters,
            availableStatuses,
            showFilters,
            assigneeFilters,
            availableAssignees,
            showAssigneeRow,
            resetZoom,
            toggleLayout,
            debugPhases,
            toggleFullscreen,
            toggleStatusFilter,
            toggleAssigneeFilter,
            toggleFilters,
            toggleAssigneeRow,
            startResize,
            closeDialog,
            handleDialogClose,
            handleDialogBackdropClick,
            copyNodeLink,
            openNodeInNewTab,
        };
    },
};
</script>

<style scoped>
.dependency-graph {
    width: 100%;
    background: var(--card-bg);
    border-radius: var(--radius-lg);
    box-shadow: var(--shadow-md);
    border: 1px solid var(--border-color);
    overflow: hidden;
}

.graph-header {
    border-bottom: 1px solid var(--border-color);
    background: var(--bg-color);
}

.header-top {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: var(--spacing-md) var(--spacing-lg);
}

.filter-row {
    max-height: 0;
    overflow: hidden;
    transition:
        max-height 0.3s ease,
        padding 0.3s ease;
    background: var(--hover-bg);
    border-top: 1px solid var(--border-color);
}

.filter-row.expanded {
    max-height: 200px;
    padding: var(--spacing-md) var(--spacing-lg);
}

.graph-header h3 {
    margin: 0;
    font-size: var(--font-xl);
    font-weight: var(--font-semibold);
    color: var(--text-primary);
}

.card-id {
    font-size: var(--font-sm);
    font-weight: var(--font-semibold);
    color: var(--text-primary);
}

.card-points-badge {
    pointer-events: none;
    transition: opacity 0.2s ease;
}

.card-points-badge:hover {
    opacity: 1;
}

.card-points-text {
    font-family: "Arial", sans-serif;
    text-shadow: 0 1px 2px rgba(0, 0, 0, 0.5);
    letter-spacing: -0.5px;
}

.phase-counter {
    font-size: var(--font-sm);
    font-weight: var(--font-normal);
    color: var(--text-secondary);
    margin-left: var(--spacing-sm);
}

.graph-controls {
    display: flex;
    gap: var(--spacing-sm);
    align-items: center;
}

.btn {
    padding: var(--spacing-sm) var(--spacing-md);
    border: 1px solid var(--border-color);
    border-radius: var(--radius-sm);
    background: var(--card-bg);
    color: var(--text-primary);
    font-size: var(--font-sm);
    cursor: pointer;
    transition: var(--transition-fast);
    white-space: nowrap;
    display: flex;
    align-items: center;
    gap: var(--spacing-xs);
}

.filter-toggle .filter-icon {
    transition: transform 0.3s ease;
    font-size: 0.75em;
    display: inline-block;
}

.filter-toggle.active .filter-icon {
    transform: rotate(180deg);
}

.filter-toggle.active {
    background: var(--primary-light);
    color: var(--primary-color);
    border-color: var(--primary-color);
}

.filter-count {
    background: var(--primary-color);
    color: white;
    border-radius: 50%;
    min-width: 1.2em;
    height: 1.2em;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    font-size: 0.75em;
    font-weight: var(--font-semibold);
}

.btn:hover {
    background: var(--hover-bg);
    border-color: var(--primary-color);
}

.btn-secondary {
    background: var(--primary-color);
    color: var(--text-inverse);
    border-color: var(--primary-color);
}

.btn-secondary:hover {
    background: var(--primary-hover);
    border-color: var(--primary-hover);
}

.graph-legend {
    display: flex;
    gap: var(--spacing-md);
    padding: var(--spacing-sm) var(--spacing-lg);
    background: var(--bg-color);
    border-bottom: 1px solid var(--border-color);
    font-size: var(--font-sm);
    color: var(--text-secondary);
}

.legend-item {
    display: flex;
    align-items: center;
    gap: var(--spacing-sm);
}

.legend-color {
    width: 16px;
    height: 16px;
    border-radius: 50%;
    border: 2px solid var(--bg-color);
}

.legend-color.epic {
    background: var(--ctp-mauve);
}

.legend-color.new {
    background: var(--ctp-blue);
}

.legend-color.indeterminate {
    background: var(--ctp-yellow);
}

.legend-color.done {
    background: var(--ctp-green);
}

.legend-color.phase-indicator {
    background: var(--primary-color);
    border-radius: 2px;
    width: 20px;
    height: 12px;
}

.legend-line {
    width: 24px;
    height: 2px;
}

.legend-line.blocks {
    background: var(--error-color);
}

.legend-line.epic-link {
    background: var(--primary-color);
    background-image: repeating-linear-gradient(
        to right,
        var(--primary-color),
        var(--primary-color) 5px,
        transparent 5px,
        transparent 10px
    );
}

.legend-line.status-line {
    background: linear-gradient(
        to right,
        var(--ctp-green) 0%,
        var(--ctp-green) 33%,
        var(--ctp-yellow) 33%,
        var(--ctp-yellow) 66%,
        var(--ctp-blue) 66%,
        var(--ctp-blue) 100%
    );
}

.graph-container {
    height: 500px;
    width: 100%;
    position: relative;
    overflow: hidden;
    background: var(--bg-color);
}

.node-dialog {
    border: 1px solid var(--border-color);
    border-radius: var(--radius-md);
    padding: 0;
    background: var(--card-bg);
    color: var(--text-primary);
    box-shadow: var(--shadow-lg);
    max-width: 350px;
    min-width: 280px;
    width: auto;
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    margin: 0;
    opacity: 0;
    scale: 0.9;
    transition:
        opacity 0.2s ease,
        scale 0.2s ease;
}

.node-dialog::backdrop {
    background: rgba(0, 0, 0, 0.5);
    backdrop-filter: blur(2px);
    opacity: 0;
    transition:
        opacity 0.2s ease,
        backdrop-filter 0.2s ease;
}

.node-dialog[open] {
    display: flex;
    flex-direction: column;
    opacity: 1;
    scale: 1;
}

.node-dialog[open]::backdrop {
    opacity: 1;
}

.dialog-content {
    padding: var(--spacing-sm);
    font-size: var(--font-sm);
}

.dialog-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: var(--spacing-sm);
    padding-bottom: var(--spacing-sm);
    border-bottom: 1px solid var(--border-color);
}

.close-dialog {
    background: none;
    border: none;
    color: var(--text-secondary);
    font-size: var(--font-xl);
    cursor: pointer;
    padding: 0;
    margin-left: var(--spacing-sm);
}

.close-dialog:hover {
    color: var(--text-primary);
}

.dialog-body {
    line-height: 1.4;
    color: var(--text-secondary);
}

.dialog-label {
    margin-bottom: var(--spacing-sm);
    color: var(--text-primary);
    font-weight: var(--font-medium);
}

.dialog-status {
    font-size: var(--font-xs);
    color: var(--text-muted);
}

.dialog-status span {
    font-weight: var(--font-semibold);
    padding: var(--spacing-xs) var(--spacing-sm);
    border-radius: var(--radius-sm);
    background: var(--hover-bg);
    color: var(--text-primary);
}

.dialog-points {
    font-size: var(--font-xs);
    color: var(--text-muted);
    margin-top: var(--spacing-xs);
}

.points-value {
    font-weight: var(--font-semibold);
    padding: var(--spacing-xs) var(--spacing-sm);
    border-radius: var(--radius-sm);
    background: var(--primary-light);
    color: var(--primary-color);
    margin-left: var(--spacing-xs);
}

.dialog-actions {
    display: flex;
    gap: var(--spacing-sm);
    margin-top: var(--spacing-md);
    padding-top: var(--spacing-sm);
    border-top: 1px solid var(--border-color);
}

.dialog-btn {
    display: flex;
    align-items: center;
    gap: var(--spacing-xs);
    padding: var(--spacing-xs) var(--spacing-sm);
    border: 1px solid var(--border-color);
    border-radius: var(--radius-sm);
    background: var(--card-bg);
    color: var(--text-primary);
    font-size: var(--font-xs);
    font-weight: var(--font-medium);
    cursor: pointer;
    transition: var(--transition-fast);
    white-space: nowrap;
}

.dialog-btn:hover {
    background: var(--hover-bg);
    border-color: var(--primary-color);
    transform: translateY(-1px);
    box-shadow: var(--shadow-sm);
}

.copy-btn:hover {
    background: var(--primary-light);
    color: var(--primary-color);
}

.open-btn:hover {
    background: var(--info-light);
    color: var(--info-color);
}

.dialog-btn:active {
    transform: translateY(0);
    box-shadow: none;
}

@media (max-width: 768px) {
    .header-top {
        flex-direction: column;
        gap: var(--spacing-md);
        align-items: stretch;
    }

    .graph-controls {
        justify-content: center;
        flex-wrap: wrap;
    }

    .status-filters {
        align-items: flex-start;
    }

    .filter-items {
        justify-content: flex-start;
        flex-direction: column;
        align-items: flex-start;
        gap: var(--spacing-sm);
    }

    .filter-label {
        font-weight: var(--font-semibold);
        margin-bottom: var(--spacing-xs);
        align-self: flex-start;
    }

    .status-filter-item {
        font-size: var(--font-sm);
        padding: var(--spacing-xs);
    }

    .filter-row.expanded {
        max-height: 400px;
    }

    .filter-count {
        font-size: 0.7em;
        min-width: 1em;
        height: 1em;
    }

    .graph-legend {
        flex-wrap: wrap;
        gap: var(--spacing-sm);
    }

    .graph-container {
        height: 400px;
    }

    .node-dialog {
        max-width: calc(100vw - 2rem);
        min-width: auto;
        width: calc(100vw - 2rem);
    }

    .dialog-actions {
        flex-direction: column;
        gap: var(--spacing-xs);
    }

    .dialog-btn {
        width: 100%;
        justify-content: center;
    }

    .assignee-row {
        margin-top: var(--spacing-xs);
        padding: var(--spacing-xs);
    }

    .assignee-filters {
        flex-direction: row;
        align-items: center;
        gap: var(--spacing-xs);
    }

    .assignee-items {
        width: auto;
        justify-content: flex-start;
        gap: 1px;
    }

    .assignee-filter-item {
        font-size: 9px;
        padding: 1px 2px;
    }

    .assignee-name {
        max-width: 30px;
        overflow: hidden;
        text-overflow: ellipsis;
    }

    .assignee-avatar {
        width: 12px;
        height: 12px;
    }
}

/* Phase layout specific styles */
.phase-label {
    font-family: inherit;
    user-select: none;
    text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.1);
}

.phase-column {
    transition: opacity 0.3s ease;
}

.phase-column:hover {
    opacity: 0.5 !important;
}

/* Card-style node styles */
.node-card {
    cursor: pointer;
    transition: all 0.2s ease;
}

.node-card:hover {
    transform: scale(1.05);
}

.card-background {
    filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1));
    transition: all 0.2s ease;
}

.node-card:hover .card-background {
    filter: drop-shadow(0 4px 8px rgba(0, 0, 0, 0.15));
}

.card-id {
    font-family: inherit;
    font-feature-settings: "tnum" 1;
    user-select: none;
}

.card-title {
    font-family: inherit;
    user-select: none;
}

.card-status {
    font-family: inherit;
    text-transform: uppercase;
    letter-spacing: 0.025em;
    user-select: none;
}

.status-indicator {
    opacity: 0.9;
}

/* Improved link styling */
.link {
    transition: all 0.15s ease-in-out;
    cursor: pointer;
    stroke-linecap: round;
    stroke-linejoin: round;
}

.link:hover {
    stroke-width: 6px !important;
    opacity: 1 !important;
    filter: drop-shadow(0 4px 12px rgba(0, 0, 0, 0.4)) brightness(1.1) !important;
    transform-origin: center;
    animation: pulse-connection 0.6s ease-in-out infinite alternate;
}

@keyframes pulse-connection {
    0% {
        filter: drop-shadow(0 4px 12px rgba(0, 0, 0, 0.4)) brightness(1.1);
    }
    100% {
        filter: drop-shadow(0 6px 16px rgba(0, 0, 0, 0.5)) brightness(1.2);
    }
}

/* Link animations */
.link path {
    stroke-dasharray: 1000;
    stroke-dashoffset: 1000;
    animation: dash 0.5s ease-in forwards;
}

@keyframes dash {
    to {
        stroke-dashoffset: 0;
    }
}

/* Dependency type specific styling */
.link[data-type="Blocks"] {
    stroke: var(--ctp-red);
}

.link[data-type="Blocks"]:hover {
    stroke: var(--ctp-maroon) !important;
    stroke-width: 6px !important;
}

.link[data-type="epic link"] {
    stroke: var(--ctp-blue);
    stroke-dasharray: 5, 5;
}

.link[data-type="epic link"]:hover {
    stroke: var(--ctp-sapphire) !important;
    stroke-width: 6px !important;
    stroke-dasharray: 8, 8 !important;
}

/* Phase layout specific improvements */
.phase-column {
    transition: all 0.3s ease;
}

.phase-column:hover {
    opacity: 0.15 !important;
    fill: var(--ctp-surface1) !important;
}

/* Enhanced phase spacing visual feedback */
.phase-label {
    transition: all 0.2s ease;
}

.phase-label:hover {
    font-size: 16px !important;
    fill: var(--ctp-blue) !important;
}

/* Status filters */
.status-filters {
    display: flex;
    flex-direction: column;
    gap: var(--spacing-md);
}

.filter-items {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    gap: var(--spacing-sm);
}

.filter-label {
    font-weight: var(--font-medium);
    color: var(--text-primary);
    margin-right: var(--spacing-sm);
    white-space: nowrap;
}

.status-filter-item {
    display: flex;
    align-items: center;
    gap: var(--spacing-xs);
    cursor: pointer;
    padding: var(--spacing-xs) var(--spacing-sm);
    border-radius: var(--radius-sm);
    transition: var(--transition-fast);
    border: 1px solid transparent;
    background: var(--card-bg);
    white-space: nowrap;
}

.status-filter-item:hover {
    background-color: var(--hover-bg);
    border-color: var(--border-color);
}

.status-filter-item input[type="checkbox"] {
    margin: 0;
    cursor: pointer;
}

.status-name {
    font-size: var(--font-sm);
    color: var(--text-secondary);
    white-space: nowrap;
}

/* Assignee row */
.assignee-row {
    margin-top: var(--spacing-xs);
    padding: var(--spacing-xs) var(--spacing-sm);
    background: transparent;
    border-radius: var(--radius-sm);
    border: 1px solid var(--border-light);
    opacity: 0.9;
}

.assignee-filters {
    display: flex;
    align-items: center;
    gap: var(--spacing-sm);
    flex-wrap: wrap;
}

.filter-label {
    font-size: var(--font-xs);
    color: var(--text-muted);
    margin-right: var(--spacing-xs);
}

.assignee-items {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    gap: 2px;
}

.assignee-filter-item {
    display: flex;
    align-items: center;
    gap: 2px;
    cursor: pointer;
    padding: 2px 4px;
    border-radius: var(--radius-xs);
    transition: var(--transition-fast);
    border: 1px solid transparent;
    background: transparent;
    white-space: nowrap;
    font-size: var(--font-xs);
}

.assignee-filter-item:hover {
    background-color: var(--hover-bg);
    border-color: var(--border-light);
}

.assignee-filter-item.active {
    background-color: var(--primary-light);
    border-color: var(--primary-color);
}

.assignee-filter-item input[type="checkbox"] {
    margin: 0;
    cursor: pointer;
    width: 12px;
    height: 12px;
    opacity: 0.7;
}

.assignee-avatar {
    width: 14px;
    height: 14px;
    border-radius: 50%;
    border: 1px solid var(--border-light);
}

.assignee-name {
    font-size: 10px;
    color: var(--text-muted);
    white-space: nowrap;
    max-width: 40px;
    overflow: hidden;
    text-overflow: ellipsis;
}

/* Card assignee text */
.card-assignee {
    font-family: inherit;
    user-select: none;
}

/* Dialog assignee */
.dialog-assignee {
    margin: var(--spacing-sm) 0;
}

.assignee-value {
    font-weight: var(--font-medium);
    color: var(--text-primary);
}

/* Resize handle */
.resize-handle {
    width: 100%;
    height: 8px;
    background:
        linear-gradient(
            45deg,
            transparent 45%,
            var(--border-color) 45%,
            var(--border-color) 55%,
            transparent 55%
        ),
        linear-gradient(
            -45deg,
            transparent 45%,
            var(--border-color) 45%,
            var(--border-color) 55%,
            transparent 55%
        );
    background-size: 8px 8px;
    cursor: ns-resize;
    transition: all 0.2s ease;
    border-radius: var(--radius-sm);
    margin-top: var(--spacing-xs);
}

.resize-handle:hover,
.resize-handle.active {
    background-color: var(--primary-color);
    transform: scaleY(1.5);
}

.resize-handle:hover {
    opacity: 0.8;
}

.resize-handle.active {
    opacity: 1;
    background-color: var(--primary-hover);
}

/* Fullscreen styles */
.dependency-graph:fullscreen {
    background-color: var(--bg-color);
    padding: var(--spacing-md);
}

.dependency-graph:-webkit-full-screen {
    background-color: var(--bg-color);
    padding: var(--spacing-md);
}

.dependency-graph:-moz-full-screen {
    background-color: var(--bg-color);
    padding: var(--spacing-md);
}

.dependency-graph:-ms-fullscreen {
    background-color: var(--bg-color);
    padding: var(--spacing-md);
}

.dependency-graph:fullscreen .graph-container {
    height: calc(100vh - 200px);
    width: 100%;
}

.dependency-graph:-webkit-full-screen .graph-container {
    height: calc(100vh - 200px);
    width: 100%;
}

.dependency-graph:-moz-full-screen .graph-container {
    height: calc(100vh - 200px);
    width: 100%;
}

.dependency-graph:-ms-fullscreen .graph-container {
    height: calc(100vh - 200px);
    width: 100%;
}

.dependency-graph:fullscreen .graph-header {
    background-color: var(--card-bg);
    padding: var(--spacing-md);
    border-radius: var(--radius-lg);
    margin-bottom: var(--spacing-md);
    box-shadow: var(--shadow-md);
}

.dependency-graph:fullscreen .graph-legend {
    background-color: var(--card-bg);
    padding: var(--spacing-md);
    border-radius: var(--radius-lg);
    margin-bottom: var(--spacing-md);
    box-shadow: var(--shadow-md);
}

/* Fullscreen button styling */
.btn:has-text("Fullscreen"),
.btn:has-text("Exit Fullscreen") {
    background-color: var(--ctp-blue);
    color: var(--text-inverse);
}

.btn:has-text("Fullscreen"):hover,
.btn:has-text("Exit Fullscreen"):hover {
    background-color: var(--ctp-sapphire);
}

/* Node highlighting effects */
.node-highlighted .card-background {
    stroke-width: 4px !important;
    filter: drop-shadow(0 4px 12px rgba(0, 0, 0, 0.3)) !important;
}

.node-connected .card-background {
    stroke-width: 4px !important;
    filter: drop-shadow(0 2px 6px rgba(0, 0, 0, 0.2)) !important;
}

.node-card {
    transition: all 0.15s ease-in-out;
}

.node-card:hover .card-background {
    transform: scale(1.02);
}

/* Link highlighting effects */
.link-highlighted {
    stroke-width: 4px !important;
    opacity: 1 !important;
    filter: drop-shadow(0 3px 8px rgba(0, 0, 0, 0.3)) !important;
}
</style>
