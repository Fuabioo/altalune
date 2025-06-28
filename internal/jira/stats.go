package jira

import (
	"github.com/Fuabioo/altalune/internal/model"
)

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
