package jira

import (
	"github.com/Fuabioo/altalune/internal/model"
)

type Graph struct {
	Nodes []GraphNode `json:"nodes"`
	Edges []GraphEdge `json:"edges"`
}

type GraphNode struct {
	ID             string  `json:"id"`             // Ticket Key
	Label          string  `json:"label"`          // Summary
	Status         string  `json:"status"`         // Status
	Classification string  `json:"classification"` // Status classification (new, indeterminate, done)
	StoryPoints    float64 `json:"storyPoints"`    // Story points
	Assignee       string  `json:"assignee"`       // Assignee display name
	AssigneeID     string  `json:"assigneeId"`     // Assignee account ID
}

type GraphEdge struct {
	From           string  `json:"from"`           // Source ticket key
	To             string  `json:"to"`             // Destination ticket key
	Type           string  `json:"type"`           // Optional: e.g., "blocks", "relates to"
	Status         string  `json:"status"`         // Status
	Classification string  `json:"classification"` // Status classification (new, indeterminate, done)
	StoryPoints    float64 `json:"storyPoints"`    // Story points of the target issue
	Assignee       string  `json:"assignee"`       // Assignee display name of the target issue
	AssigneeID     string  `json:"assigneeId"`     // Assignee account ID of the target issue
}

func BuildGraph(issues []*model.Ticket, epicKey string) Graph {
	nodeMap := make(map[string]GraphNode)
	var edges []GraphEdge

	// Build nodes from all issues
	for _, issue := range issues {
		storyPoints := getStoryPoints(issue)
		assigneeName := ""
		assigneeID := ""
		if issue.Fields.Assignee.DisplayName != "" {
			assigneeName = issue.Fields.Assignee.DisplayName
			assigneeID = issue.Fields.Assignee.AccountID
		}

		nodeMap[issue.Key] = GraphNode{
			ID:             issue.Key,
			Label:          issue.Fields.Summary,
			Status:         issue.Fields.Status.Name,
			Classification: issue.Fields.StatusCategory.Key,
			StoryPoints:    storyPoints,
			Assignee:       assigneeName,
			AssigneeID:     assigneeID,
		}
	}

	// Add the epic as a node (if not already in issues)
	if _, exists := nodeMap[epicKey]; !exists {
		nodeMap[epicKey] = GraphNode{
			ID:             epicKey,
			Label:          "Epic",
			Status:         "Epic",
			Classification: "epic", // Special classification for epics
			StoryPoints:    0,
			Assignee:       "",
			AssigneeID:     "",
		}
	}

	// Add default edges from epic to each ticket
	for _, issue := range issues {
		if issue.Key != epicKey {
			storyPoints := getStoryPoints(issue)
			assigneeName := ""
			assigneeID := ""
			if issue.Fields.Assignee.DisplayName != "" {
				assigneeName = issue.Fields.Assignee.DisplayName
				assigneeID = issue.Fields.Assignee.AccountID
			}

			edges = append(edges, GraphEdge{
				From:           epicKey,
				To:             issue.Key,
				Type:           "epic link",
				Status:         "Epic",
				Classification: "epic",
				StoryPoints:    storyPoints,
				Assignee:       assigneeName,
				AssigneeID:     assigneeID,
			})
		}

		// Include issue links
		for _, link := range issue.Fields.IssueLinks {
			// Outward
			if link.OutwardIssue.Key != "" {
				outwardStoryPoints := getStoryPoints(&link.OutwardIssue)
				outwardAssigneeName := ""
				outwardAssigneeID := ""
				if link.OutwardIssue.Fields.Assignee.DisplayName != "" {
					outwardAssigneeName = link.OutwardIssue.Fields.Assignee.DisplayName
					outwardAssigneeID = link.OutwardIssue.Fields.Assignee.AccountID
				}

				edges = append(edges, GraphEdge{
					From:           issue.Key,
					To:             link.OutwardIssue.Key,
					Type:           link.Type.Name,
					Status:         link.OutwardIssue.Fields.Status.Name,
					Classification: link.OutwardIssue.Fields.StatusCategory.Key,
					StoryPoints:    outwardStoryPoints,
					Assignee:       outwardAssigneeName,
					AssigneeID:     outwardAssigneeID,
				})
			}
			// Inward
			if link.InwardIssue.Key != "" {
				inwardStoryPoints := getStoryPoints(&link.InwardIssue)
				inwardAssigneeName := ""
				inwardAssigneeID := ""
				if link.InwardIssue.Fields.Assignee.DisplayName != "" {
					inwardAssigneeName = link.InwardIssue.Fields.Assignee.DisplayName
					inwardAssigneeID = link.InwardIssue.Fields.Assignee.AccountID
				}

				edges = append(edges, GraphEdge{
					From:           link.InwardIssue.Key,
					To:             issue.Key,
					Type:           link.Type.Name,
					Status:         link.InwardIssue.Fields.Status.Name,
					Classification: link.InwardIssue.Fields.StatusCategory.Key,
					StoryPoints:    inwardStoryPoints,
					Assignee:       inwardAssigneeName,
					AssigneeID:     inwardAssigneeID,
				})
			}
		}
	}

	// Build final node list
	nodes := make([]GraphNode, 0, len(nodeMap))
	for _, node := range nodeMap {
		nodes = append(nodes, node)
	}

	return Graph{
		Nodes: nodes,
		Edges: edges,
	}
}

// getStoryPoints extracts story points from common custom fields
// Jira stores story points in custom fields with different IDs depending on the configuration:
// - customfield_10016: Most common in Jira Cloud instances
// - customfield_10002: Common in Jira Server instances
// - customfield_10004: Alternative configuration
// - customfield_10008: Another alternative
// This function tries each field in order and returns the first non-zero value found
func getStoryPoints(issue *model.Ticket) float64 {
	// Common story points custom field IDs
	storyPointFields := []string{
		"customfield_10016", // Common in Jira Cloud
		"customfield_10002", // Common in Jira Server
		"customfield_10004", // Another common one
		"customfield_10008", // Alternative
	}

	for _, fieldID := range storyPointFields {
		if points := issue.Fields.GetCustomFieldAsFloat(fieldID); points > 0 {
			return points
		}
	}

	return 0
}
