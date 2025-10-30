package domain

import "context"

type ProjectRepository interface {
	Create(ctx context.Context, project *Project) error
	GetByID(ctx context.Context, id string) (*Project, error)
	GetAll(ctx context.Context) ([]*Project, error)
	Update(ctx context.Context, project *Project) error
	Delete(ctx context.Context, id string) error
}
