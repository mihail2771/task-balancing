package main

import (
	"fmt"
	"log/slog"
	"os"
	"task-balancing/internal/app"
	"task-balancing/internal/config"
)

func main() {
	//========================================================
	cfg := config.MustLoad()
	//========================================================

	//========================================================
	log := setupLogger(cfg.Env)
	log.Info("logger initialized")
	//========================================================

	application, err := app.New(cfg, log)
	if err != nil {
		log.Error("failed to init app", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer application.Close()

	go func() {
		log.Info("app started")
		if err := application.Run(); err != nil {
			log.Error("failed to start app", slog.String("error", err.Error()))
		}
	}()

	fmt.Println("app finished")
}

func setupLogger(env string) *slog.Logger {
	var level slog.Level

	switch env {
	case "dev":
		level = slog.LevelDebug
	default:
		level = slog.LevelInfo
	}

	return slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}),
	)
}
