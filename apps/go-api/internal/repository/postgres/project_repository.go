package postgres

import (
	"api/internal/domain"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepository struct {
	pool *pgxpool.Pool
}

func NewProjectRepository(pool *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{pool: pool}
}

func (r *ProjectRepository) Create(ctx context.Context, project *domain.Project) error {
	query := `
		INSERT INTO projects (id, name, description, expiration_date, created_at, updated_at)
		VALUES (gen_random_uuid(), $1, $2, $3, $4, $5)
		RETURNING id
	`

	err := r.pool.QueryRow(
		ctx,
		query,
		project.Name,
		project.Description,
		project.ExpirationDate,
		project.CreatedAt,
		project.UpdatedAt,
	).Scan(&project.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProjectRepository) GetByID(ctx context.Context, id string) (*domain.Project, error) {
	query := `
		SELECT id, name, description, expiration_date, created_at, updated_at
		FROM projects
		WHERE id = $1
	`

	var project domain.Project
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.ExpirationDate,
		&project.CreatedAt,
		&project.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("project not found")
		}
		return nil, err
	}

	return &project, nil
}

func (r *ProjectRepository) GetAll(ctx context.Context) ([]*domain.Project, error) {
	query := `
		SELECT id, name, description, expiration_date, created_at, updated_at
		FROM projects
		ORDER BY created_at DESC
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*domain.Project
	for rows.Next() {
		var project domain.Project
		err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.Description,
			&project.ExpirationDate,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *ProjectRepository) Update(ctx context.Context, project *domain.Project) error {
	query := `
		UPDATE projects
		SET name = $1, description = $2, expiration_date = $3, updated_at = $4
		WHERE id = $5
	`

	result, err := r.pool.Exec(
		ctx,
		query,
		project.Name,
		project.Description,
		project.ExpirationDate,
		project.UpdatedAt,
		project.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("project not found")
	}

	return nil
}

func (r *ProjectRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM projects WHERE id = $1`

	result, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("project not found")
	}

	return nil
}
