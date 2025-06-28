package model

import (
	"encoding/json"
	"strings"
	"time"
)

type SearchResult struct {
	StartAt    int       `json:"startAt"`
	MaxResults int       `json:"maxResults"`
	Total      uint      `json:"total"`
	Issues     []*Ticket `json:"issues"`
}

type Ticket struct {
	Expand string `json:"expand"`
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

type IssueLink struct {
	ID           string   `json:"id"`
	Self         string   `json:"self"`
	Type         LinkType `json:"type"`
	InwardIssue  Ticket   `json:"inwardIssue"`
	OutwardIssue Ticket   `json:"outwardIssue"`
}

type LinkType struct {
	ID      string `json:"id"`
	Self    string `json:"self"`
	Name    string `json:"name"`
	Inward  string `json:"inward"`
	Outward string `json:"outward"`
}

// Fields represents the fields section of a JIRA ticket
type Fields struct {
	StatusCategoryChangeDate *JiraTime          `json:"statuscategorychangedate"`
	Parent                   *ParentIssue       `json:"parent"`
	FixVersions              []any              `json:"fixVersions"`
	StatusCategory           StatusCategory     `json:"statusCategory"`
	Resolution               any                `json:"resolution"`
	LastViewed               *JiraTime          `json:"lastViewed"`
	Priority                 *Priority          `json:"priority"`
	Labels                   []string           `json:"labels"`
	TimeEstimate             any                `json:"timeestimate"`
	AggregateTimeOriginalEst any                `json:"aggregatetimeoriginalestimate"`
	IssueLinks               []IssueLink        `json:"issuelinks"`
	Assignee                 User               `json:"assignee"`
	Status                   Status             `json:"status"`
	Components               []any              `json:"components"`
	AggregateTimeEstimate    any                `json:"aggregatetimeestimate"`
	Creator                  User               `json:"creator"`
	Subtasks                 []any              `json:"subtasks"`
	Reporter                 User               `json:"reporter"`
	AggregateProgress        *Progress          `json:"aggregateprogress"`
	Progress                 *Progress          `json:"progress"`
	Votes                    *Votes             `json:"votes"`
	IssueType                IssueType          `json:"issuetype"`
	TimeSpent                any                `json:"timespent"`
	Project                  Project            `json:"project"`
	AggregateTimeSpent       any                `json:"aggregatetimespent"`
	ResolutionDate           any                `json:"resolutiondate"`
	WorkRatio                int                `json:"workratio"`
	Watches                  *Watches           `json:"watches"`
	Created                  JiraTime           `json:"created"`
	Updated                  JiraTime           `json:"updated"`
	TimeOriginalEstimate     any                `json:"timeoriginalestimate"`
	Description              *AtlassianDocument `json:"description"`
	Summary                  string             `json:"summary"`
	DueDate                  any                `json:"duedate"`
	CustomFields             map[string]any     `json:"-"`
}

// ParentIssue represents a parent issue in JIRA
type ParentIssue struct {
	ID     string       `json:"id"`
	Key    string       `json:"key"`
	Self   string       `json:"self"`
	Fields ParentFields `json:"fields"`
}

// ParentFields represents the fields of a parent issue
type ParentFields struct {
	Summary   string     `json:"summary"`
	Status    *Status    `json:"status"`
	IssueType *IssueType `json:"issuetype"`
}

// StatusCategory represents a JIRA status category
type StatusCategory struct {
	Self      string `json:"self"`
	ID        int    `json:"id"`
	Key       string `json:"key"`
	ColorName string `json:"colorName"`
	Name      string `json:"name"`
}

// Priority represents a JIRA priority
type Priority struct {
	Self    string `json:"self"`
	IconURL string `json:"iconUrl"`
	Name    string `json:"name"`
	ID      string `json:"id"`
}

// Status represents a JIRA status
type Status struct {
	Self           string          `json:"self"`
	Description    string          `json:"description"`
	IconURL        string          `json:"iconUrl"`
	Name           string          `json:"name"`
	ID             string          `json:"id"`
	StatusCategory *StatusCategory `json:"statusCategory"`
}

// User represents a JIRA user
type User struct {
	Self         string            `json:"self"`
	AccountID    string            `json:"accountId"`
	EmailAddress string            `json:"emailAddress"`
	AvatarURLs   map[string]string `json:"avatarUrls"`
	DisplayName  string            `json:"displayName"`
	Active       bool              `json:"active"`
	TimeZone     string            `json:"timeZone"`
	AccountType  string            `json:"accountType"`
}

// Progress represents progress information
type Progress struct {
	Progress int `json:"progress"`
	Total    int `json:"total"`
}

// Votes represents voting information
type Votes struct {
	Self     string `json:"self"`
	Votes    int    `json:"votes"`
	HasVoted bool   `json:"hasVoted"`
}

// IssueType represents a JIRA issue type
type IssueType struct {
	Self           string `json:"self"`
	ID             string `json:"id"`
	Description    string `json:"description"`
	IconURL        string `json:"iconUrl"`
	Name           string `json:"name"`
	Subtask        bool   `json:"subtask"`
	AvatarID       int    `json:"avatarId"`
	HierarchyLevel int    `json:"hierarchyLevel"`
}

// Project represents a JIRA project
type Project struct {
	Self           string            `json:"self"`
	ID             string            `json:"id"`
	Key            string            `json:"key"`
	Name           string            `json:"name"`
	ProjectTypeKey string            `json:"projectTypeKey"`
	Simplified     bool              `json:"simplified"`
	AvatarURLs     map[string]string `json:"avatarUrls"`
}

// Watches represents watch information
type Watches struct {
	Self       string `json:"self"`
	WatchCount int    `json:"watchCount"`
	IsWatching bool   `json:"isWatching"`
}

// AtlassianDocument represents Atlassian Document Format (ADF)
type AtlassianDocument struct {
	Type    string                  `json:"type"`
	Version int                     `json:"version"`
	Content []AtlassianDocumentNode `json:"content"`
}

// AtlassianDocumentNode represents a node in an Atlassian document
type AtlassianDocumentNode struct {
	Type    string                  `json:"type"`
	Content []AtlassianDocumentNode `json:"content,omitempty"`
	Text    string                  `json:"text,omitempty"`
	Marks   []AtlassianDocumentMark `json:"marks,omitempty"`
}

// AtlassianDocumentMark represents a mark in an Atlassian document
type AtlassianDocumentMark struct {
	Type string `json:"type"`
}

// JiraTime wraps time.Time to handle JIRA's date format
type JiraTime struct {
	time.Time
}

// UnmarshalJSON implements custom JSON unmarshalling for JiraTime
func (jt *JiraTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	// Remove quotes from JSON string
	s := strings.Trim(string(data), `"`)

	// Try different time formats that JIRA might use
	layouts := []string{
		"2006-01-02T15:04:05.000-0700",
		"2006-01-02T15:04:05.000Z0700",
		"2006-01-02T15:04:05-0700",
		"2006-01-02T15:04:05Z0700",
		time.RFC3339,
		time.RFC3339Nano,
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			jt.Time = t
			return nil
		}
	}

	return nil // Return nil to handle unparsable dates gracefully
}

// MarshalJSON implements custom JSON marshaling for JiraTime
func (jt JiraTime) MarshalJSON() ([]byte, error) {
	if jt.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + jt.Time.Format("2006-01-02T15:04:05.000-0700") + `"`), nil
}

// CustomFieldOption represents a custom field option
type CustomFieldOption struct {
	Self  string `json:"self"`
	Value string `json:"value"`
	ID    string `json:"id"`
}

// UnmarshalJSON implements custom JSON unmarshalling for Fields
func (f *Fields) UnmarshalJSON(data []byte) error {
	// Create a temporary struct that matches the JSON structure exactly
	type FieldsAlias Fields

	// Unmarshal into a map to capture all fields
	var fieldMap map[string]any
	if err := json.Unmarshal(data, &fieldMap); err != nil {
		return err
	}

	// Unmarshal into the alias to populate known fields
	var fieldsAlias FieldsAlias
	if err := json.Unmarshal(data, &fieldsAlias); err != nil {
		return err
	}

	// Copy the alias to the original struct
	*f = Fields(fieldsAlias)

	// Initialize the custom fields map
	f.CustomFields = make(map[string]any)

	// Extract custom fields (fields starting with "customfield_")
	for key, value := range fieldMap {
		if len(key) > 12 && key[:12] == "customfield_" {
			f.CustomFields[key] = value
		}
	}

	return nil
}

// GetCustomField returns a custom field value by field ID
func (f *Fields) GetCustomField(fieldID string) (any, bool) {
	if f.CustomFields == nil {
		return nil, false
	}
	value, exists := f.CustomFields[fieldID]
	return value, exists
}

// GetCustomFieldAsString returns a custom field as a string
func (f *Fields) GetCustomFieldAsString(fieldID string) string {
	if value, exists := f.GetCustomField(fieldID); exists {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return ""
}

// GetCustomFieldAsFloat returns a custom field as a float64
func (f *Fields) GetCustomFieldAsFloat(fieldID string) float64 {
	if value, exists := f.GetCustomField(fieldID); exists {
		if f, ok := value.(float64); ok {
			return f
		}
	}
	return 0
}

// GetCustomFieldAsOption returns a custom field as a CustomFieldOption
func (f *Fields) GetCustomFieldAsOption(fieldID string) *CustomFieldOption {
	if value, exists := f.GetCustomField(fieldID); exists {
		if optionMap, ok := value.(map[string]any); ok {
			option := &CustomFieldOption{}
			if self, ok := optionMap["self"].(string); ok {
				option.Self = self
			}
			if val, ok := optionMap["value"].(string); ok {
				option.Value = val
			}
			if id, ok := optionMap["id"].(string); ok {
				option.ID = id
			}
			return option
		}
	}
	return nil
}
