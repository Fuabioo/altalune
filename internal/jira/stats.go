package jira

import (
	"github.com/Fuabioo/altalune/internal/model"
)

type Assignee struct {
	AccountID   string `json:"accountId"`
	DisplayName string `json:"displayName"`
	AvatarURL   string `json:"avatarUrl"`
	Active      bool   `json:"active"`
}

type EpicStats struct {
	Total       int     `json:"total"`
	ToDo        int     `json:"toDo"`
	InProgress  int     `json:"inProgress"`
	Done        int     `json:"done"`
	Percentage  float64 `json:"percentage"`
	ProgressPer float64 `json:"progressPer"`
}

type StatusClassified struct {
	Classification string `json:"classification"`
	Count          uint   `json:"count"`
}
type StatusCounts map[string]*StatusClassified
type TypeCounts map[string]int

func CalculateStats(issues []*model.Ticket) EpicStats {
	stats := EpicStats{}

	for _, issue := range issues {
		stats.Total++

		classification := issue.Fields.StatusCategory.Key
		switch classification {
		case "new":
			stats.ToDo++
		case "indeterminate":
			stats.InProgress++
		case "done":
			stats.Done++
		}
	}

	if stats.Total > 0 {
		stats.Percentage = (float64(stats.Done) / float64(stats.Total)) * 100
	}

	totalWithoutCompleted := stats.Total - stats.Done
	stats.ProgressPer = (float64(stats.InProgress) / float64(totalWithoutCompleted)) * 100

	return stats
}

func CalculateStatusCounts(issues []*model.Ticket) StatusCounts {
	counts := make(StatusCounts)

	for _, issue := range issues {
		status := issue.Fields.Status.Name
		value, ok := counts[status]
		if !ok {
			counts[status] = &StatusClassified{
				Classification: issue.Fields.StatusCategory.Key,
				Count:          1,
			}
		} else {
			value.Count++
		}

	}

	return counts
}

func CalculateTypeCounts(issues []*model.Ticket) TypeCounts {
	counts := make(TypeCounts)

	for _, issue := range issues {
		issueType := issue.Fields.IssueType.Name
		counts[issueType]++
	}

	return counts
}

func ExtractAssignees(issues []*model.Ticket) []Assignee {
	assigneeMap := make(map[string]Assignee)

	for _, issue := range issues {
		if issue.Fields.Assignee.AccountID != "" {
			avatarURL := getAvatarURL(issue.Fields.Assignee.AvatarURLs)
			assigneeMap[issue.Fields.Assignee.AccountID] = Assignee{
				AccountID:   issue.Fields.Assignee.AccountID,
				DisplayName: issue.Fields.Assignee.DisplayName,
				AvatarURL:   avatarURL,
				Active:      issue.Fields.Assignee.Active,
			}
		}
	}

	// Convert map to slice
	assignees := make([]Assignee, 0, len(assigneeMap))
	for _, assignee := range assigneeMap {
		assignees = append(assignees, assignee)
	}

	return assignees
}

// getAvatarURL extracts the best available avatar URL from the available URLs
func getAvatarURL(avatarURLs map[string]string) string {
	// Preferred sizes in order of preference
	preferredSizes := []string{"48x48", "32x32", "24x24", "16x16"}

	for _, size := range preferredSizes {
		if url, exists := avatarURLs[size]; exists && url != "" {
			return url
		}
	}

	// If no preferred size is found, return any available URL
	for _, url := range avatarURLs {
		if url != "" {
			return url
		}
	}

	return ""
}
