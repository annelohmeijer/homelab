package domain

import (
	"errors"
	"time"
)

type Project struct {
	ID             string
	Name           string
	Description    string
	ExpirationDate time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewProject(name, description string, expirationDate time.Time) (*Project, error) {
	if name == "" {
		return nil, errors.New("project name cannot be empty")
	}
	if expirationDate.Before(time.Now()) {
		return nil, errors.New("expiration date cannot be in the past")
	}

	now := time.Now()
	return &Project{
		Name:           name,
		Description:    description,
		ExpirationDate: expirationDate,
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

func (p *Project) Update(name, description string, expirationDate time.Time) error {
	if name == "" {
		return errors.New("project name cannot be empty")
	}
	if expirationDate.Before(time.Now()) {
		return errors.New("expiration date cannot be in the past")
	}

	p.Name = name
	p.Description = description
	p.ExpirationDate = expirationDate
	p.UpdatedAt = time.Now()

	return nil
}

func (p *Project) IsExpired() bool {
	return time.Now().After(p.ExpirationDate)
}
