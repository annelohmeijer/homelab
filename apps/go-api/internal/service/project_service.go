package service

import (
	"api/internal/domain"
	"context"
	"time"
)

type ProjectService struct {
	repo domain.ProjectRepository
}

func NewProjectService(repo domain.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) CreateProject(ctx context.Context, name, description string, expirationDate time.Time) (*domain.Project, error) {
	project, err := domain.NewProject(name, description, expirationDate)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) GetProject(ctx context.Context, id string) (*domain.Project, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ProjectService) GetAllProjects(ctx context.Context) ([]*domain.Project, error) {
	return s.repo.GetAll(ctx)
}

func (s *ProjectService) UpdateProject(ctx context.Context, id, name, description string, expirationDate time.Time) (*domain.Project, error) {
	project, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := project.Update(name, description, expirationDate); err != nil {
		return nil, err
	}

	if err := s.repo.Update(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) DeleteProject(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
