package services

import (
  "api/domain/locations"
  "api/utils/errors"
  "api/providers/locations_provider"
)

type locationsService struct{}

// locationsServiceInterface is implemented by locationsService struct
type locationsServiceInterface interface {
  GetCountry(countryId string) (*locations.Country, *errors.ApiError)
}

var(
  LocationsService locationsServiceInterface    // LocationsService has type locationsServiceInterface
)

func init() {
  LocationsService = &locationsService{}
}

// this function is a method for locationsService
func (s *locationsService) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
  return locations_provider.GetCountry(countryId)
}
