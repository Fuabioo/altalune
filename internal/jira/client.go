package jira

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Fuabioo/altalune/internal/model"
	cliutils "github.com/Fuabioo/altalune/pkg/cliutls"

	"github.com/charmbracelet/log"
	"resty.dev/v3"
)

type Client struct {
	client *resty.Client
}

type Config struct {
	Email      string
	Token      string
	Workspace  string
	SuperDebug bool
}

func NewClient(cfg Config) *Client {

	maskedToken := cliutils.MaskToken(cfg.Token)
	log.Debug("Initializing Jira client",
		"workspace", cfg.Workspace,
		"email", cfg.Email,
		"token", maskedToken,
		"superDebug", cfg.SuperDebug,
	)

	return &Client{
		client: resty.New().
			SetDebug(cfg.SuperDebug).
			SetLogger(log.Default()).
			SetHeader("Content-Type", "application/json").
			SetHeader("Accept", "application/json").
			SetBasicAuth(cfg.Email, cfg.Token).
			SetBaseURL(fmt.Sprintf("https://%s/rest/api/3", cfg.Workspace)),
	}
}

func (c *Client) Ping(ctx context.Context) error {
	resp, err := c.client.R().
		SetContext(ctx).
		Get("/myself")
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return nil
}

type ListEpicRequest struct {
	EpicID   string
	StartAt  uint
	PageSize uint
}

func (c *Client) ListEpicIssues(ctx context.Context, req ListEpicRequest) (*model.SearchResult, error) {

	resp, err := c.client.R().
		SetContext(ctx).
		SetQueryParam("jql", fmt.Sprintf("parent = %s", req.EpicID)).
		SetQueryParam("startAt", fmt.Sprintf("%d", req.StartAt)).
		SetQueryParam("maxResults", fmt.Sprintf("%d", req.PageSize)).
		Get("/search")
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	var result model.SearchResult
	if err := json.Unmarshal(resp.Bytes(), &result); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return &result, nil
}
