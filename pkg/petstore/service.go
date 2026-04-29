// Package petstore This file is generated ONCE as a starting point and will NOT be overwritten.
// Modify it freely to add your business logic.
// To regenerate, delete this file or set generate.handler.output.overwrite: true in config.
package petstore

import (
	"context"

	"github.com/mockzilla/mockzilla/v2/pkg/api"
)

// service implements the ServiceInterface with your business logic.
// Return nil, nil to fall back to the generator for mock responses.
// Return a response to override the generated response.
// Return an error to return an error response.
type service struct {
	params *api.ServiceParams
}

// Ensure service implements ServiceInterface.
var _ ServiceInterface = (*service)(nil)

// newService creates a new service instance.
func newService(params *api.ServiceParams) *service {
	return &service{params: params}
}

// FindPets handles GET /pets
func (s *service) FindPets(ctx context.Context, opts *FindPetsServiceRequestOptions) (*FindPetsResponseData, error) {
	// TODO: Implement your business logic here.
	// Return nil, nil to use the generated mock response.
	return nil, nil
}

// AddPet handles POST /pets
func (s *service) AddPet(ctx context.Context, opts *AddPetServiceRequestOptions) (*AddPetResponseData, error) {
	// TODO: Implement your business logic here.
	// Return nil, nil to use the generated mock response.
	return nil, nil
}

// FindPetByID handles GET /pets/{id}
func (s *service) FindPetByID(ctx context.Context, opts *FindPetByIDServiceRequestOptions) (*FindPetByIDResponseData, error) {
	// TODO: Implement your business logic here.
	// Return nil, nil to use the generated mock response.
	return nil, nil
}

// DeletePet handles DELETE /pets/{id}
func (s *service) DeletePet(ctx context.Context, opts *DeletePetServiceRequestOptions) (*DeletePetResponseData, error) {
	// TODO: Implement your business logic here.
	// Return nil, nil to use the generated mock response.
	return nil, nil
}
