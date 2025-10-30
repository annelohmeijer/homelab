package container

import (
	"api/internal/handler"
	"api/internal/repository/postgres"
	"api/internal/service"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	DB             *pgxpool.Pool
	ProjectHandler *handler.ProjectHandler
}

func New(ctx context.Context, dbURL string) (*Container, error) {
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	projectRepo := postgres.NewProjectRepository(pool)
	projectService := service.NewProjectService(projectRepo)
	projectHandler := handler.NewProjectHandler(projectService)

	return &Container{
		DB:             pool,
		ProjectHandler: projectHandler,
	}, nil
}

func (c *Container) Close() {
	c.DB.Close()
}
