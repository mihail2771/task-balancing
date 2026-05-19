package app

import (
	"context"
	"log/slog"
	"net/http"
	"task-balancing/internal/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	server *http.Server
	db     *pgxpool.Pool
	log    *slog.Logger
	cfg    *config.Config
}

func New(cfg *config.Config, log *slog.Logger) (*App, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.HTTPServer.Timeout*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	log.Debug("database connected")

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}
	log.Debug("database pinged")

	server := &http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: nil,
	}

	return &App{
		server: server,
		db:     pool,
		log:    log,
		cfg:    cfg,
	}, nil
}

func (a *App) Run() error {
	a.log.Info("http server started", slog.String("addr", a.cfg.HTTPServer.Address))
	return a.server.ListenAndServe()
}

func (a *App) Shutdown(ctx context.Context) error {
	a.log.Info("shutting down server")
	return a.server.Shutdown(ctx)
}

func (a *App) Close() {
	if a.db != nil {
		a.db.Close()
	}
}
