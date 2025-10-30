package handler

import (
	"api/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	service *service.ProjectService
}

func NewProjectHandler(service *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: service}
}

type CreateProjectRequest struct {
	Name           string    `json:"name" binding:"required"`
	Description    string    `json:"description"`
	ExpirationDate time.Time `json:"expiration_date" binding:"required"`
}

type UpdateProjectRequest struct {
	Name           string    `json:"name" binding:"required"`
	Description    string    `json:"description"`
	ExpirationDate time.Time `json:"expiration_date" binding:"required"`
}

type ProjectResponse struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	IsExpired      bool      `json:"is_expired"`
}

// CreateProject godoc
// @Summary Create a new project
// @Description Create a new project with name, description and expiration date
// @Tags projects
// @Accept json
// @Produce json
// @Param project body CreateProjectRequest true "Project data"
// @Success 201 {object} ProjectResponse
// @Failure 400 {object} map[string]string
// @Router /projects [post]
func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var req CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := h.service.CreateProject(c.Request.Context(), req.Name, req.Description, req.ExpirationDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ProjectResponse{
		ID:             project.ID,
		Name:           project.Name,
		Description:    project.Description,
		ExpirationDate: project.ExpirationDate,
		CreatedAt:      project.CreatedAt,
		UpdatedAt:      project.UpdatedAt,
		IsExpired:      project.IsExpired(),
	})
}

// GetProject godoc
// @Summary Get a project by ID
// @Description Get a single project by its ID
// @Tags projects
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} ProjectResponse
// @Failure 404 {object} map[string]string
// @Router /projects/{id} [get]
func (h *ProjectHandler) GetProject(c *gin.Context) {
	id := c.Param("id")

	project, err := h.service.GetProject(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProjectResponse{
		ID:             project.ID,
		Name:           project.Name,
		Description:    project.Description,
		ExpirationDate: project.ExpirationDate,
		CreatedAt:      project.CreatedAt,
		UpdatedAt:      project.UpdatedAt,
		IsExpired:      project.IsExpired(),
	})
}

// GetAllProjects godoc
// @Summary Get all projects
// @Description Get a list of all projects
// @Tags projects
// @Produce json
// @Success 200 {array} ProjectResponse
// @Failure 500 {object} map[string]string
// @Router /projects [get]
func (h *ProjectHandler) GetAllProjects(c *gin.Context) {
	projects, err := h.service.GetAllProjects(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := make([]ProjectResponse, 0, len(projects))
	for _, project := range projects {
		responses = append(responses, ProjectResponse{
			ID:             project.ID,
			Name:           project.Name,
			Description:    project.Description,
			ExpirationDate: project.ExpirationDate,
			CreatedAt:      project.CreatedAt,
			UpdatedAt:      project.UpdatedAt,
			IsExpired:      project.IsExpired(),
		})
	}

	c.JSON(http.StatusOK, responses)
}

// UpdateProject godoc
// @Summary Update a project
// @Description Update an existing project by ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param project body UpdateProjectRequest true "Project data"
// @Success 200 {object} ProjectResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /projects/{id} [put]
func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	id := c.Param("id")

	var req UpdateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := h.service.UpdateProject(c.Request.Context(), id, req.Name, req.Description, req.ExpirationDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ProjectResponse{
		ID:             project.ID,
		Name:           project.Name,
		Description:    project.Description,
		ExpirationDate: project.ExpirationDate,
		CreatedAt:      project.CreatedAt,
		UpdatedAt:      project.UpdatedAt,
		IsExpired:      project.IsExpired(),
	})
}

// DeleteProject godoc
// @Summary Delete a project
// @Description Delete a project by ID
// @Tags projects
// @Param id path string true "Project ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /projects/{id} [delete]
func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteProject(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
