package jira

import (
	"github.com/Fuabioo/altalune/internal/model"
)

type Graph struct {
	Nodes []GraphNode `json:"nodes"`
	Edges []GraphEdge `json:"edges"`
}

type GraphNode struct {
	ID             string `json:"id"`             // Ticket Key
	Label          string `json:"label"`          // Summary
	Status         string `json:"status"`         // Status
	Classification string `json:"classification"` // Status classification (new, indeterminate, done)
}

type GraphEdge struct {
	From           string `json:"from"`           // Source ticket key
	To             string `json:"to"`             // Destination ticket key
	Type           string `json:"type"`           // Optional: e.g., "blocks", "relates to"
	Status         string `json:"status"`         // Status
	Classification string `json:"classification"` // Status classification (new, indeterminate, done)
}

func BuildGraph(issues []*model.Ticket, epicKey string) Graph {
	nodeMap := make(map[string]GraphNode)
	var edges []GraphEdge

	// Build nodes from all issues
	for _, issue := range issues {
		nodeMap[issue.Key] = GraphNode{
			ID:             issue.Key,
			Label:          issue.Fields.Summary,
			Status:         issue.Fields.Status.Name,
			Classification: issue.Fields.StatusCategory.Key,
		}
	}

	// Add the epic as a node (if not already in issues)
	if _, exists := nodeMap[epicKey]; !exists {
		nodeMap[epicKey] = GraphNode{
			ID:             epicKey,
			Label:          "Epic",
			Status:         "Epic",
			Classification: "epic", // Special classification for epics
		}
	}

	// Add default edges from epic to each ticket
	for _, issue := range issues {
		if issue.Key != epicKey {
			edges = append(edges, GraphEdge{
				From:   epicKey,
				To:     issue.Key,
				Type:   "epic link",
				Status: "Epic",
			})
		}

		// Include issue links
		for _, link := range issue.Fields.IssueLinks {
			// Outward
			if link.OutwardIssue.Key != "" {
				edges = append(edges, GraphEdge{
					From:           issue.Key,
					To:             link.OutwardIssue.Key,
					Type:           link.Type.Name,
					Status:         link.OutwardIssue.Fields.Status.Name,
					Classification: link.OutwardIssue.Fields.StatusCategory.Key,
				})
			}
			// Inward
			if link.InwardIssue.Key != "" {
				edges = append(edges, GraphEdge{
					From:           link.InwardIssue.Key,
					To:             issue.Key,
					Type:           link.Type.Name,
					Status:         link.InwardIssue.Fields.Status.Name,
					Classification: link.InwardIssue.Fields.StatusCategory.Key,
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
