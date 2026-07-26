package libs

import (
	"backend/internals/models"
)

type UserResponseHttp struct {
	Success bool
	Status  int
	Message string
	Queries *QueryParams
	Results []*models.UserResponse
}
