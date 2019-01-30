package usecases

import (
	"github.com/ghostec/Will.IAM/models"
	"github.com/ghostec/Will.IAM/repositories"
)

// Permissions define entrypoints for Permissions actions
type Permissions interface {
	Get(string) (*models.Permission, error)
	Delete(string) error
	Create(*models.Permission) error
	CreateRequest(string, *models.PermissionRequest) error
	GetPermissionRequests(string) ([]models.PermissionRequest, error)
	ForRoles([]models.Role) ([]models.Permission, error)
}

type permissions struct {
	repo *repositories.All
}

func (ps permissions) Get(id string) (*models.Permission, error) {
	return ps.repo.Permissions.Get(id)
}

func (ps permissions) Delete(id string) error {
	return ps.repo.Permissions.Delete(id)
}

func (ps permissions) CreateRequest(
	saID string, r *models.PermissionRequest,
) error {
	r.State = models.PermissionRequestStates.Created
	return ps.repo.Permissions.CreateRequest(saID, r)
}

func (ps permissions) Create(p *models.Permission) error {
	return ps.repo.Permissions.Create(p)
}

func (ps permissions) ForRoles(
	rs []models.Role,
) ([]models.Permission, error) {
	return ps.repo.Permissions.ForRoles(rs)
}

func (ps permissions) GetPermissionRequests(
	saID string,
) ([]models.PermissionRequest, error) {
	return ps.repo.Permissions.GetPermissionRequests(saID)
}

// NewPermissions ctor
func NewPermissions(repo *repositories.All) Permissions {
	return &permissions{repo: repo}
}
