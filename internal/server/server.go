package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/Fuabioo/altalune/internal/jira"
	"github.com/Fuabioo/altalune/internal/model"

	"github.com/charmbracelet/log"
)

type (
	config struct {
		superDebug bool
		server     serverConfig
		jira       jira.Config
	}
	serverConfig struct {
		host   string
		port   uint
		assets fs.FS
	}

	Server struct {
		config config
		server *http.Server
		client *jira.Client
	}
	Option func(*config)
)

func ServerAddress(host string, port uint) Option {
	return func(c *config) {
		c.server.host = host
		c.server.port = port
	}
}

func ServerAssets(assets fs.FS) Option {
	return func(c *config) {
		c.server.assets = assets
	}
}

func ServerJira(workspace string, email string, token string) Option {
	return func(c *config) {
		c.jira.Workspace = workspace
		c.jira.Email = email
		c.jira.Token = token
	}
}

func ServerSuperDebug(superDebug bool) Option {
	return func(c *config) {
		c.superDebug = superDebug
	}
}

func NewServer(options ...Option) (*Server, error) {
	cfg := &config{
		server: serverConfig{},
		jira:   jira.Config{},
	}

	for _, option := range options {
		option(cfg)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.server.host, cfg.server.port),
		Handler: http.DefaultServeMux,
	}

	cfg.jira.SuperDebug = cfg.superDebug

	return &Server{
		config: *cfg,
		server: server,
		client: jira.NewClient(cfg.jira),
	}, nil
}

func (s *Server) Start() error {

	log.Debug("Starting server",
		"host", s.config.server.host,
		"port", s.config.server.port,
	)

	router := http.NewServeMux()

	router.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		err := s.client.Ping(r.Context())
		if err != nil {
			log.Error("Error pinging Jira", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write([]byte(`{}`))
		w.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("/api/epic/{ticket}", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		ticket := r.PathValue("ticket")

		var response struct {
			Stats        jira.EpicStats    `json:"stats"`
			Graph        jira.Graph        `json:"graph"`
			StatusCounts jira.StatusCounts `json:"statusCounts"`
			TypeCounts   jira.TypeCounts   `json:"typeCounts"`
			Issues       []*model.Ticket   `json:"issues"`
			Epic         *model.Ticket     `json:"epic"`
			All          []*model.Ticket   `json:"tickets"`
			Total        int               `json:"total"`
			JiraBaseURL  string            `json:"jiraBaseUrl"`
			Assignees    []jira.Assignee   `json:"assignees"`
		}

		startAt := uint(0)
		pageSize := uint(50)

		for {
			result, err := s.client.ListEpicIssues(ctx, jira.ListEpicRequest{
				EpicID:   ticket,
				StartAt:  startAt,
				PageSize: pageSize,
			})
			if err != nil {
				log.Error("Error listing epic issues", "err", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			log.Info("Epic issues listed",
				"startedAt", result.StartAt,
				"total", result.Total,
				"max", result.MaxResults,
			)
			response.All = append(response.All, result.Issues...)

			if result.Total <= startAt+pageSize {
				break
			}

			startAt += pageSize
		}

		response.Total = len(response.All)
		response.Stats = jira.CalculateStats(response.All)
		response.StatusCounts = jira.CalculateStatusCounts(response.All)
		response.TypeCounts = jira.CalculateTypeCounts(response.All)
		response.Graph = jira.BuildGraph(response.All, ticket)
		response.Issues = response.All
		response.JiraBaseURL = s.config.jira.Workspace
		response.Assignees = jira.ExtractAssignees(response.All)

		// Find epic in the issues (if it exists)
		for _, issue := range response.All {
			if issue.Key == ticket {
				response.Epic = issue
				break
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	frontendFS := http.FileServer(http.FS(s.config.server.assets))

	router.Handle("/", frontendFS)

	s.server.Handler = router

	err := s.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Close() {
	if err := s.server.Close(); err != nil {
		log.Error("Error closing server", "err", err)
	}
}
